package exchange

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type sAdminExchange struct{}

func New() *sAdminExchange {
	return &sAdminExchange{}
}

// GetOwnCoinList 获取自发币列表
func (s *sAdminExchange) GetOwnCoinList(ctx context.Context, req *v1.GetExchangeOwnCoinListReq) (*v1.GetExchangeOwnCoinListRes, error) {
	m := dao.OwnCoin.Ctx(ctx)
	if req.Coin != "" {
		m = m.WhereLike("coin", "%"+req.Coin+"%")
	}
	if req.ReferCoin != "" {
		m = m.WhereLike("refer_coin", "%"+req.ReferCoin+"%")
	}

	total, _ := m.Count()
	var list []entity.OwnCoin
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.ExchangeOwnCoinInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.ExchangeOwnCoinInfo{
			Id:            c.Id,
			Coin:          c.Coin,
			Logo:          c.Logo,
			ReferCoin:     c.ReferCoin,
			ReferMarket:   c.ReferMarket,
			ShowSymbol:    c.ShowSymbol,
			Price:         c.Price,
			Proportion:    c.Proportion,
			RaisingAmount: c.RaisingAmount,
			RaisedAmount:  c.RaisedAmount,
			PurchaseLimit: c.PurchaseLimit,
			TotalAmount:   c.TotalAmount,
			Status:        c.Status,
			Introduce:     c.Introduce,
			Remark:        c.Remark,
			BeginTime:     c.BeginTime.Format("2006-01-02 15:04:05"),
			EndTime:       c.EndTime.Format("2006-01-02 15:04:05"),
			CreateTime:    c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetExchangeOwnCoinListRes{
		List:  resList,
		Total: total,
	}, nil
}

// AddOwnCoin 新增自发币
func (s *sAdminExchange) AddOwnCoin(ctx context.Context, req *v1.AddExchangeOwnCoinReq) (*v1.AddExchangeOwnCoinRes, error) {
	_, err := dao.OwnCoin.Ctx(ctx).Data(entity.OwnCoin{
		Coin:          req.Coin,
		Logo:          req.Logo,
		ReferCoin:     req.ReferCoin,
		ReferMarket:   req.ReferMarket,
		ShowSymbol:    req.ShowSymbol,
		Price:         req.Price,
		Proportion:    req.Proportion,
		RaisingAmount: req.RaisingAmount,
		PurchaseLimit: req.PurchaseLimit,
		TotalAmount:   req.TotalAmount,
		Status:        req.Status,
		Introduce:     req.Introduce,
		Remark:        req.Remark,
		// Example Time Parsing, add error check if critical in prod
		BeginTime: gtime.New(req.BeginTime),
		EndTime:   gtime.New(req.EndTime),
	}).OmitEmpty().Insert()
	return &v1.AddExchangeOwnCoinRes{}, err
}

// EditOwnCoin 修改自发币
func (s *sAdminExchange) EditOwnCoin(ctx context.Context, req *v1.EditExchangeOwnCoinReq) (*v1.EditExchangeOwnCoinRes, error) {
	_, err := dao.OwnCoin.Ctx(ctx).Where("id", req.Id).Data(entity.OwnCoin{
		Coin:          req.Coin,
		Logo:          req.Logo,
		ReferCoin:     req.ReferCoin,
		ReferMarket:   req.ReferMarket,
		ShowSymbol:    req.ShowSymbol,
		Price:         req.Price,
		Proportion:    req.Proportion,
		RaisingAmount: req.RaisingAmount,
		PurchaseLimit: req.PurchaseLimit,
		TotalAmount:   req.TotalAmount,
		Status:        req.Status,
		Introduce:     req.Introduce,
		Remark:        req.Remark,
		BeginTime:     gtime.New(req.BeginTime),
		EndTime:       gtime.New(req.EndTime),
	}).Update()
	return &v1.EditExchangeOwnCoinRes{}, err
}

// DeleteOwnCoin 删除自发币
func (s *sAdminExchange) DeleteOwnCoin(ctx context.Context, req *v1.DeleteExchangeOwnCoinReq) (*v1.DeleteExchangeOwnCoinRes, error) {
	_, err := dao.OwnCoin.Ctx(ctx).Where("id", req.Id).Delete()
	return &v1.DeleteExchangeOwnCoinRes{}, err
}
