package trading

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	adminV1 "GoCEX/api/admin/v1"
	admin_v1 "GoCEX/api/admin/v1"
	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"
)

type sTrading struct{}

func New() *sTrading {
	return &sTrading{}
}

// CurrencyOrderSubmit 现货挂单 (市价/限价)
func (s *sTrading) CurrencyOrderSubmit(ctx context.Context, userId uint64, in *v1.CurrencyOrderSubmitReq) error {
	pairs := strings.Split(in.Symbol, "/")
	if len(pairs) != 2 {
		return gerror.NewCode(codes.ClientError, "非法的交易对格式")
	}
	targetCoin, baseCoin := pairs[0], pairs[1]

	var freezeSymbol string
	var freezeAmount float64

	// 买入(0): 冻结计价币 (BaseCoin, 例如 USDT)
	// 卖出(1): 冻结目标币 (TargetCoin, 例如 BTC)
	if in.Type == 0 {
		freezeSymbol = baseCoin
		if in.DelegateType == 1 { // 市价买入: DelegateTotal 是法币/USDT额度
			freezeAmount = in.DelegateTotal
		} else { // 限价买入: DelegateTotal 是数量，需乘价格
			freezeAmount = in.DelegateTotal * in.DelegatePrice
		}
	} else if in.Type == 1 {
		freezeSymbol = targetCoin
		freezeAmount = in.DelegateTotal // 卖出时，市价限价冻结的都是持仓数量
	}

	if freezeAmount <= 0 {
		return gerror.NewCode(codes.ClientError, "非法的挂单金额/数量")
	}

	// 1. 调用资产层的安全锁冻结
	freezeReq := &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     freezeSymbol,
		Amount:     freezeAmount,
		RecordType: 30, // 假设 30 为现货挂单冻结
		Remark:     "现货挂单冻结",
	}

	_, err := asset.New().FreezeAmount(ctx, freezeReq)
	if err != nil {
		return err // 余额不足或并发获取锁失败直接打回
	}

	// 2. 生成订单落库
	orderNo := "S" + gtime.Now().Format("YmdHis") // 后续可接 Snowflake
	order := g.Map{
		dao.CurrencyOrder.Columns().UserId:        userId,
		dao.CurrencyOrder.Columns().Type:          in.Type,
		dao.CurrencyOrder.Columns().DelegateType:  in.DelegateType,
		dao.CurrencyOrder.Columns().Status:        0, // 等待成交
		dao.CurrencyOrder.Columns().OrderNo:       orderNo,
		dao.CurrencyOrder.Columns().Symbol:        in.Symbol,
		dao.CurrencyOrder.Columns().Coin:          baseCoin,
		dao.CurrencyOrder.Columns().DelegateTotal: in.DelegateTotal,
		dao.CurrencyOrder.Columns().DelegatePrice: in.DelegatePrice,
		dao.CurrencyOrder.Columns().DelegateValue: freezeAmount,
		dao.CurrencyOrder.Columns().DelegateTime:  gtime.Now(),
		dao.CurrencyOrder.Columns().DealNum:       0,
		dao.CurrencyOrder.Columns().DealPrice:     0,
		dao.CurrencyOrder.Columns().DealValue:     0,
		dao.CurrencyOrder.Columns().Fee:           0,
		dao.CurrencyOrder.Columns().CreateBy:      "",
		dao.CurrencyOrder.Columns().UpdateBy:      "",
	}

	res, err := dao.CurrencyOrder.Ctx(ctx).Data(order).Insert()
	if err != nil {
		// 应当有回滚方案: 如果此步出错，理论上应解冻 (调用 amount = -freezeAmount)
		return err
	}
	orderId, _ := res.LastInsertId()

	// 3. 将现货单打入 Redis 取代原 Java 的跨服异步通信
	_, _ = g.Redis().Do(ctx, "XADD", "CEX:STREAM:SPOT_MATCHING", "*", "orderId", orderId)
	return nil
}

// SecondContractSubmit 秒合约下注 (含外挂 Buff 读取)
func (s *sTrading) SecondContractSubmit(ctx context.Context, userId uint64, in *v1.SecondContractSubmitReq) error {
	pairs := strings.Split(in.Symbol, "/")
	if len(pairs) != 2 {
		return gerror.NewCode(codes.ClientError, "非法的交易对")
	}
	baseCoin := pairs[1] // 秒合约通常用 USDT 作为基础下注货币

	// 1. 冻结下注资金
	freezeReq := &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     baseCoin,
		Amount:     in.BetAmount,
		RecordType: 40, // 假设 40 为秒合约下注扣减
		Remark:     "秒合约下注",
	}

	_, err := asset.New().FreezeAmount(ctx, freezeReq)
	if err != nil {
		return err
	}

	// 2. 取出用户身上的 Buff 标记 (杀猪盘后台人工预设输赢)
	var u entity.AppUser
	_ = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).Scan(&u)

	// User.Buff: 0正常, 1包赢, 2包输
	buffSign := u.Buff

	// 3. 落库
	orderNo := "O" + gtime.Now().Format("YmdHis")
	order := g.Map{
		dao.SecondContractOrder.Columns().UserId:      userId,
		dao.SecondContractOrder.Columns().UserAddress: u.Address,
		dao.SecondContractOrder.Columns().OrderNo:     orderNo,
		dao.SecondContractOrder.Columns().Symbol:      in.Symbol,
		dao.SecondContractOrder.Columns().BetContent:  in.BetContent,
		dao.SecondContractOrder.Columns().BetAmount:   in.BetAmount,
		dao.SecondContractOrder.Columns().Status:      0, // 参与中
		dao.SecondContractOrder.Columns().OpenTime:    gtime.Now().Unix(),
		// 关盘时间 = 当前时间 + 选择的 Period (周期如30秒)
		dao.SecondContractOrder.Columns().CloseTime: gtime.Now().Unix() + in.Period,
		dao.SecondContractOrder.Columns().Sign:      buffSign, // 植入必杀/必赢标记，交由底层 worker 结算使用
		dao.SecondContractOrder.Columns().CreateBy:  "",
		dao.SecondContractOrder.Columns().UpdateBy:  "",
	}

	res, err := dao.SecondContractOrder.Ctx(ctx).Data(order).Insert()
	if err != nil {
		// todo: fallback
		return err
	}

	orderId, _ := res.LastInsertId()
	_, _ = g.Redis().Do(ctx, "XADD", "CEX:STREAM:SECOND_CONTRACT", "*", "orderId", orderId)

	return nil
}

// ContractOrderSubmit 永续合约开仓
func (s *sTrading) ContractOrderSubmit(ctx context.Context, userId uint64, in *v1.ContractOrderSubmitReq) error {
	pairs := strings.Split(in.Symbol, "/")
	if len(pairs) != 2 {
		return gerror.NewCode(codes.ClientError, "非法的交易对")
	}
	baseCoin := pairs[1] // U本位，扣减 USDT

	if in.Leverage <= 0 {
		return gerror.NewCode(codes.ClientError, "非法的杠杆")
	}

	// 保证金 = 数量 * 委托价 / 杠杆倍数
	var marginAmount float64
	if in.DelegateType == 1 {
		// 市价，这里的 Total 当保证金数额扣除计算
		marginAmount = in.DelegateTotal
	} else {
		// 限价
		marginAmount = in.DelegateTotal * in.DelegatePrice / in.Leverage
	}

	if marginAmount <= 0 {
		return gerror.NewCode(codes.ClientError, "开仓保证金过低")
	}

	freezeReq := &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     baseCoin,
		Amount:     marginAmount,
		RecordType: 50, // 假设 50 为合约开仓保证金提取
		Remark:     "合约开仓保证金",
	}

	_, err := asset.New().FreezeAmount(ctx, freezeReq)
	if err != nil {
		return err
	}

	orderNo := "C" + gtime.Now().Format("YmdHis")
	order := g.Map{
		dao.ContractOrder.Columns().UserId:        userId,
		dao.ContractOrder.Columns().OrderNo:       orderNo,
		dao.ContractOrder.Columns().Symbol:        in.Symbol,
		dao.ContractOrder.Columns().Type:          in.Type,
		dao.ContractOrder.Columns().DelegateType:  in.DelegateType,
		dao.ContractOrder.Columns().Leverage:      in.Leverage,
		dao.ContractOrder.Columns().BaseCoin:      baseCoin,
		dao.ContractOrder.Columns().DelegateTotal: in.DelegateTotal,
		dao.ContractOrder.Columns().DelegatePrice: in.DelegatePrice,
		dao.ContractOrder.Columns().DelegateValue: marginAmount * in.Leverage,
		dao.ContractOrder.Columns().Status:        0,
	}

	res, err := dao.ContractOrder.Ctx(ctx).Data(order).Insert()
	if err != nil {
		// todo: fallback
		return err
	}

	orderId, _ := res.LastInsertId()
	_, _ = g.Redis().Do(ctx, "XADD", "CEX:STREAM:PERPETUAL_MATCHING", "*", "orderId", orderId)

	return nil
}

// CancelOrder 撤销现货委单 (CAS 乐观锁防重退款)
func (s *sTrading) CancelOrder(ctx context.Context, userId uint64, req *v1.CurrencyOrderCancelReq) error {
	// 1. 查询订单
	var order entity.CurrencyOrder
	err := dao.CurrencyOrder.Ctx(ctx).Where(dao.CurrencyOrder.Columns().OrderNo, req.OrderNo).Scan(&order)
	if err != nil {
		return gerror.NewCode(codes.ClientError, "订单不存在")
	}

	if order.UserId != int64(userId) {
		return gerror.NewCode(codes.Forbidden, "非法操作越权")
	}

	if order.Status != 0 {
		return gerror.NewCode(codes.ClientError, "订单已不在可撤销状态")
	}

	// 2. CAS 乐观锁尝试修改状态: UPDATE ... SET status = 3 WHERE id = ? AND status = 0
	// 只有争抢到这条锁 (RowsAffected > 0)，才能执行后续的退款，彻底防范网络高并发导致的重复退款漏洞
	result, err := dao.CurrencyOrder.Ctx(ctx).
		Where(dao.CurrencyOrder.Columns().OrderNo, req.OrderNo).
		Where(dao.CurrencyOrder.Columns().Status, 0).
		Data(g.Map{dao.CurrencyOrder.Columns().Status: 3}).
		Update()

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows <= 0 {
		return gerror.NewCode(codes.ClientError, "撤单失败，订单状态已改变或正在撮合中")
	}

	// 3. 执行资产解冻与返还
	// 退款数量：未成交的部分 (总委托 - 已成交) * 价格 (限价单退计价币、市价买退计价币、卖退基础币)
	// 在此处简化模型，全量退回剩余可用冻结 (实际业务须对应买卖方向和成交比例清算)
	var refundCoin string
	var refundAmount float64

	if order.Type == 0 { // 买入，退还剩下的计价币 (如 USDT)
		refundCoin = order.Coin
		if order.DelegateType == 0 { // 限价
			refundAmount = (order.DelegateTotal - order.DealNum) * order.DelegatePrice
		} else { // 市价
			refundAmount = order.DelegateTotal - order.DealValue
		}
	} else { // 卖出，退还剩下的基础币 (如 BTC)
		refundCoin = order.Symbol
		refundAmount = order.DelegateTotal - order.DealNum
	}

	if refundAmount > 0 {
		// 调取资产聚合服务撤销冻结
		// 这里是解除冻结，等于退回可用余额，但账单类型应标记为 "撤单解冻"
		_, _ = asset.New().AddAmount(ctx, &admin_v1.SubAmountReq{
			UserId:     int64(userId),
			Symbol:     refundCoin,
			Amount:     refundAmount,
			RecordType: 4, // 撤单解冻
			Remark:     "手工撤单解冻 (乐观锁校验通过)",
		})
	}

	return nil
}

// AdjustMargin 调整仓位保证金
func (s *sTrading) AdjustMargin(ctx context.Context, userId uint64, req *v1.AdjustPositionMarginReq) (newAmount float64, err error) {
	var pos entity.ContractPosition
	err = dao.ContractPosition.Ctx(ctx).Where("order_no", req.OrderNo).Where("user_id", userId).Scan(&pos)
	if err != nil || pos.Id == 0 {
		return 0, gerror.NewCode(codes.Failed, "仓位不存在")
	}

	// 0 增加 1 减少
	changeAmount := req.Amount
	if req.Type == 1 {
		changeAmount = -req.Amount
	}

	// 调用资产模块进行实际余额扣减/增加
	// 保证金增加对应可用余额减少 (SubAmount)
	// 保证金减少对应可用余额增加 (AddAmount)
	_, err = asset.New().SubAmount(ctx, &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     "USDT",
		Amount:     -changeAmount, // 如果 changeAmount 为 10, 则这里是 -10 (扣钱); 如果为 -5, 则这里是 5 (加钱)
		RecordType: 51,            // 合约保证金调整
		Remark:     fmt.Sprintf("调整持仓保证金: %s, 变动: %f", req.OrderNo, changeAmount),
	}, func(ctx context.Context, tx gdb.TX) error {
		// 回调: 更新仓位保证金
		newAmount = pos.Amount + changeAmount
		if newAmount < 0 {
			return gerror.NewCode(codes.Failed, "保证金不足以扣除")
		}

		// 简单逻辑: 这里应该重新计算强平价 ClosePrice
		// 强平价计算公式复杂，暂维持原值或按比例微调
		_, err := dao.ContractPosition.Ctx(ctx).TX(tx).
			Data(g.Map{
				"amount":        newAmount,
				"adjust_amount": pos.AdjustAmount + changeAmount,
				"update_time":   gtime.Now(),
			}).
			Where("id", pos.Id).
			Update()
		return err
	})

	return newAmount, err
}

// ContractLossSett 设置止盈止损
func (s *sTrading) ContractLossSett(ctx context.Context, userId uint64, req *v1.ContractLossSettReq) error {
	var pos entity.ContractPosition
	err := dao.ContractPosition.Ctx(ctx).Where("order_no", req.OrderNo).Where("user_id", userId).Scan(&pos)
	if err != nil || pos.Id == 0 {
		return gerror.NewCode(codes.Failed, "仓位不存在")
	}

	// 插入或更新止盈止损表
	// 为了简化，我们直接更新 contract_position 表中的 earn_rate / loss_rate 字段 (如果表里有的话)
	// 同时也插入一条记录到 contract_loss
	_, err = dao.ContractPosition.Ctx(ctx).
		Data(g.Map{
			"earn_rate":   req.EarnRate,
			"loss_rate":   req.LossRate,
			"update_time": gtime.Now(),
		}).
		Where("id", pos.Id).
		Update()

	if err != nil {
		return err
	}

	_, err = dao.ContractLoss.Ctx(ctx).Data(g.Map{
		"position_id": gconv.String(pos.Id),
		"user_id":     userId,
		"earn_price":  req.EarnRate,
		"lose_price":  req.LossRate,
		"status":      0,
		"create_time": gtime.Now(),
		"symbol":      pos.Symbol,
		"leverage":    pos.Leverage,
	}).Insert()

	return err
}

// ClosePosition 市价平仓
func (s *sTrading) ClosePosition(ctx context.Context, userId uint64, req *v1.ClosePositionReq) (profit float64, err error) {
	var pos entity.ContractPosition
	err = dao.ContractPosition.Ctx(ctx).Where("order_no", req.OrderNo).Where("user_id", userId).Scan(&pos)
	if err != nil || pos.Id == 0 {
		return 0, gerror.NewCode(codes.Failed, "仓位不存在")
	}

	// 1. 获取当前市价
	symbol := pos.Symbol // 如 BTC/USDT
	priceV, _ := g.Redis().Get(ctx, "CURRENCY_PRICE:"+symbol)
	currentPrice := priceV.Float64()
	if currentPrice <= 0 {
		return 0, gerror.NewCode(codes.Failed, "获取实时行情失败，无法平仓")
	}

	// 2. 计算盈亏
	// 多单盈亏 = (现价 - 开仓价) * 数量
	// 空单盈亏 = (开仓价 - 现价) * 数量
	if pos.Type == 0 {
		profit = (currentPrice - pos.OpenPrice) * pos.OpenNum
	} else {
		profit = (pos.OpenPrice - currentPrice) * pos.OpenNum
	}

	// 3. 退还保证金 + 盈亏 (AddAmount)
	totalReturn := pos.Amount + profit
	if totalReturn < 0 {
		totalReturn = 0 // 最多扣完保证金
	}

	_, err = asset.New().AddAmount(ctx, &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     "USDT",
		Amount:     totalReturn,
		RecordType: 52, // 合约平仓结算
		Remark:     fmt.Sprintf("合约平仓结算: %s, 盈亏: %f", req.OrderNo, profit),
	}, func(ctx context.Context, tx gdb.TX) error {
		// 修改仓位状态为已关闭 (状态 2)
		_, err := dao.ContractPosition.Ctx(ctx).TX(tx).
			Data(g.Map{
				"status":      2,
				"deal_price":  currentPrice,
				"earn":        profit,
				"update_time": gtime.Now(),
			}).
			Where("id", pos.Id).
			Update()
		return err
	})

	return profit, err
}
