package trading

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	adminV1 "GoCEX/api/admin/v1"
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
