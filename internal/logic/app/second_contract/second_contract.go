package second_contract

import (
	"context"
	"fmt"
	"time"

	"strings"

	v1a "GoCEX/api/admin/v1"
	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAppSecond struct{}

func New() *sAppSecond {
	return &sAppSecond{}
}

// GetCoinList 查询币种与周期配置
func (s *sAppSecond) GetCoinList(ctx context.Context, req *v1.SecondCoinListReq) (*v1.SecondCoinListRes, error) {
	var coins []entity.SecondCoinConfig
	err := dao.SecondCoinConfig.Ctx(ctx).Where("status", 1).Where("show_flag", 1).Order("sort asc").Scan(&coins)
	if err != nil {
		return nil, err
	}

	var coinIds []int64
	for _, c := range coins {
		coinIds = append(coinIds, c.Id)
	}

	var periods []entity.SecondPeriodConfig
	if len(coinIds) > 0 {
		_ = dao.SecondPeriodConfig.Ctx(ctx).Where("status", 1).WhereIn("second_id", coinIds).Order("period asc").Scan(&periods)
	}

	periodMap := make(map[int64][]v1.SecondPeriodInfo)
	for _, p := range periods {
		periodMap[p.SecondId] = append(periodMap[p.SecondId], v1.SecondPeriodInfo{
			Id:        p.Id,
			SecondId:  p.SecondId,
			Period:    p.Period,
			Odds:      p.Odds,
			MaxAmount: p.MaxAmount,
			MinAmount: p.MinAmount,
		})
	}

	res := &v1.SecondCoinListRes{Rows: make([]v1.SecondCoinInfo, 0)}
	for _, c := range coins {
		res.Rows = append(res.Rows, v1.SecondCoinInfo{
			Id:         c.Id,
			Symbol:     c.Symbol,
			Coin:       c.Coin,
			BaseCoin:   c.BaseCoin,
			Logo:       c.Logo,
			PeriodList: periodMap[c.Id],
		})
	}

	return res, nil
}

// GetCoinDetail 查询单个期权配置
func (s *sAppSecond) GetCoinDetail(ctx context.Context, req *v1.SecondCoinDetailReq) (*v1.SecondCoinDetailRes, error) {
	var coin entity.SecondCoinConfig
	err := dao.SecondCoinConfig.Ctx(ctx).Where("id", req.Id).Where("status", 1).Scan(&coin)
	if err != nil || coin.Id == 0 {
		return nil, gerror.NewCode(codes.Failed, "该秒合约配置不存在")
	}

	var periods []entity.SecondPeriodConfig
	_ = dao.SecondPeriodConfig.Ctx(ctx).Where("status", 1).Where("second_id", coin.Id).Order("period asc").Scan(&periods)

	var pList []v1.SecondPeriodInfo
	for _, p := range periods {
		pList = append(pList, v1.SecondPeriodInfo{
			Id:        p.Id,
			SecondId:  p.SecondId,
			Period:    p.Period,
			Odds:      p.Odds,
			MaxAmount: p.MaxAmount,
			MinAmount: p.MinAmount,
		})
	}

	return &v1.SecondCoinDetailRes{
		Data: v1.SecondCoinInfo{
			Id:         coin.Id,
			Symbol:     coin.Symbol,
			Coin:       coin.Coin,
			BaseCoin:   coin.BaseCoin,
			Logo:       coin.Logo,
			PeriodList: pList,
		},
	}, nil
}

// CreateOrder 新增订单
func (s *sAppSecond) CreateOrder(ctx context.Context, userId uint64, req *v1.CreateSecondOrderReq) (*v1.CreateSecondOrderRes, error) {
	if req.BetAmount <= 0 {
		return nil, gerror.NewCode(codes.Failed, "投注金额必须大于0")
	}

	var period entity.SecondPeriodConfig
	err := dao.SecondPeriodConfig.Ctx(ctx).Where("id", req.PeriodId).Where("status", 1).Scan(&period)
	if err != nil || period.Id == 0 {
		return nil, gerror.NewCode(codes.Failed, "无效的周期配置")
	}

	if req.BetAmount < period.MinAmount || (period.MaxAmount > 0 && req.BetAmount > period.MaxAmount) {
		return nil, gerror.NewCode(codes.Failed, fmt.Sprintf("投注金额需在 %v ~ %v 之间", period.MinAmount, period.MaxAmount))
	}

	var coin entity.SecondCoinConfig
	err = dao.SecondCoinConfig.Ctx(ctx).Where("id", period.SecondId).Where("status", 1).Scan(&coin)
	if err != nil || coin.Id == 0 {
		return nil, gerror.NewCode(codes.Failed, "币种配置不存在或已停用")
	}

	reqSymNorm := strings.ReplaceAll(strings.ToLower(req.Symbol), "_", "")
	dbSymNorm := strings.ReplaceAll(strings.ToLower(coin.Symbol), "_", "")

	if reqSymNorm != dbSymNorm {
		return nil, gerror.NewCode(codes.Failed, "交易对配置不匹配: 期望 "+coin.Symbol)
	}

	orderNo := "SC" + gtime.Now().Format("YmdHis") + gconv.String(userId)

	// 获取行情开盘价
	// 简单实现，这里模拟取当前市场最新价，可通过 ticker 获取
	var openPrice float64 = 0.0 // 根据实际行情对接

	err = dao.SecondContractOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 扣减资产 (BaseCoin, 通常是 USDT)
		g.Log().Debugf(ctx, "CreateOrder SubAmount: userId=%d, symbol=%s, amount=%f", userId, strings.ToUpper(coin.BaseCoin), -req.BetAmount)
		_, err := asset.New().SubAmount(ctx, &v1a.SubAmountReq{
			UserId:     gconv.Int64(userId),
			Symbol:     strings.ToUpper(coin.BaseCoin),
			Amount:     -req.BetAmount,
			RecordType: 22, // 22 是秒合约下单类型
			Remark:     "秒合约下单",
		}, func(ctx context.Context, subTx gdb.TX) error {
			// 2. 写入订单（用 g.Map 而非 struct 避免 OmitEmpty 吃掉 status=0 等零値字段）
			nowTS := time.Now().Unix()
			_, errAct := dao.SecondContractOrder.Ctx(ctx).TX(subTx).Data(g.Map{
				"order_no":    orderNo,
				"symbol":      req.Symbol,
				"type":        "1",
				"user_id":     int(userId),
				"bet_content": req.BetContent,
				"status":      0, // 0 = 参与中
				"bet_amount":  req.BetAmount,
				"create_time": gtime.Now(),
				"open_price":  openPrice,
				"open_time":   nowTS,
				"close_time":  nowTS + int64(period.Period),
				"coin_symbol": coin.Coin,
				"base_symbol": coin.BaseCoin,
				"rate":        period.Odds,
				"rate_flag":   period.Flag,
				"is_handling": 0, // 0 = 未处理
			}).Insert()
			return errAct
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return &v1.CreateSecondOrderRes{OrderNo: orderNo}, nil
}

// SelectOrderList 查询订单
func (s *sAppSecond) SelectOrderList(ctx context.Context, userId uint64, req *v1.SelectSecondOrderListReq) (*v1.SelectSecondOrderListRes, error) {
	m := dao.SecondContractOrder.Ctx(ctx).Where("user_id", userId)
	if req.Status != -1 {
		// 传了具体的 status 才过滤
		// frontend 可能传 0 或 1
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.SecondContractOrder
	err := m.Page(req.Page, req.Size).Order("create_time desc").Scan(&list)
	if err != nil {
		return nil, err
	}

	res := &v1.SelectSecondOrderListRes{
		Rows:  make([]v1.SecondOrderItem, 0),
		Total: total,
	}

	for _, o := range list {
		res.Rows = append(res.Rows, v1.SecondOrderItem{
			Id:         o.Id,
			OrderNo:    o.OrderNo,
			Symbol:     o.Symbol,
			BetContent: o.BetContent,
			BetAmount:  o.BetAmount,
			Status:     o.Status,
			OpenPrice:  o.OpenPrice,
			ClosePrice: o.ClosePrice,
			RewardAmt:  o.RewardAmount,
			CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
			OpenTime:   o.OpenTime,
			CloseTime:  o.CloseTime,
		})
	}

	return res, nil
}
