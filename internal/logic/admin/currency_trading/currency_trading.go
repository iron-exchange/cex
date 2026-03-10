package currency_trading

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminCurrencyTrading struct{}

func New() *sAdminCurrencyTrading {
	return &sAdminCurrencyTrading{}
}

// GetCurrencySymbolList 获取现货币币交易对配置
func (s *sAdminCurrencyTrading) GetCurrencySymbolList(ctx context.Context, req *v1.GetCurrencySymbolListReq) (*v1.GetCurrencySymbolListRes, error) {
	m := dao.CurrencySymbol.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}

	total, _ := m.Count()
	var list []entity.CurrencySymbol
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.CurrencySymbolInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.CurrencySymbolInfo{
			Id:            c.Id,
			Symbol:        c.Symbol,
			ShowSymbol:    c.ShowSymbol,
			Coin:          c.Coin,
			BaseCoin:      c.BaseCoin,
			FeeRate:       c.FeeRate,
			CoinPrecision: c.CoinPrecision,
			BasePrecision: c.BasePrecision,
			SellMin:       c.SellMin,
			BuyMax:        c.BuyMax,
			OrderMin:      c.OrderMin,
			OrderMax:      c.OrderMax,
			Enable:        c.Enable,
			IsShow:        c.IsShow,
			IsDeal:        c.IsDeal,
			MarketBuy:     c.MarketBuy,
			MarketSell:    c.MarketSell,
			LimitedBuy:    c.LimitedBuy,
			LimitedSell:   c.LimitedSell,
			Logo:          c.Logo,
			Market:        c.Market,
			CreateTime:    c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetCurrencySymbolListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetCurrencyOrderList 获取币币委托历史订单
func (s *sAdminCurrencyTrading) GetCurrencyOrderList(ctx context.Context, req *v1.GetAdminCurrencyOrderListReq) (*v1.GetAdminCurrencyOrderListRes, error) {
	m := dao.CurrencyOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("order_no", "%"+req.OrderNo+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.CurrencyOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminCurrencyOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.AdminCurrencyOrderInfo{
			Id:            o.Id,
			UserId:        o.UserId,
			Type:          o.Type,
			DelegateType:  o.DelegateType,
			Status:        o.Status,
			OrderNo:       o.OrderNo,
			Symbol:        o.Symbol,
			Coin:          o.Coin,
			DelegateTotal: o.DelegateTotal,
			DelegatePrice: o.DelegatePrice,
			DelegateValue: o.DelegateValue,
			DealNum:       o.DealNum,
			DealPrice:     o.DealPrice,
			DealValue:     o.DealValue,
			Fee:           o.Fee,
			CreateTime:    o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminCurrencyOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
