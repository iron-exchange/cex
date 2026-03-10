package ieo

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminIeo struct{}

func New() *sAdminIeo {
	return &sAdminIeo{}
}

// GetOwnCoinList 查询 IEO 新币列表
func (s *sAdminIeo) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (*v1.GetOwnCoinListRes, error) {
	m := dao.OwnCoin.Ctx(ctx)
	if req.Coin != "" {
		m = m.WhereLike("coin", "%"+req.Coin+"%")
	}

	total, _ := m.Count()
	var list []entity.OwnCoin
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.OwnCoinInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.OwnCoinInfo{
			Id:            c.Id,
			Coin:          c.Coin,
			Logo:          c.Logo,
			ReferCoin:     c.ReferCoin,
			ShowSymbol:    c.ShowSymbol,
			Price:         c.Price,
			Proportion:    c.Proportion,
			RaisingAmount: c.RaisingAmount,
			RaisedAmount:  c.RaisedAmount,
			PurchaseLimit: c.PurchaseLimit,
			TotalAmount:   c.TotalAmount,
			Status:        c.Status,
			BeginTime:     c.BeginTime.Format("2006-01-02 15:04:05"),
			EndTime:       c.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetOwnCoinListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOwnCoinSubscribeOrderList 查询申购打新订单列表
func (s *sAdminIeo) GetOwnCoinSubscribeOrderList(ctx context.Context, req *v1.GetOwnCoinSubscribeOrderListReq) (*v1.GetOwnCoinSubscribeOrderListRes, error) {
	m := dao.OwnCoinSubscribeOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OwnCoin != "" {
		m = m.WhereLike("own_coin", "%"+req.OwnCoin+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.OwnCoinSubscribeOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.OwnCoinSubscribeOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.OwnCoinSubscribeOrderInfo{
			Id:          o.Id,
			SubscribeId: o.SubscribeId,
			UserId:      o.UserId,
			OrderId:     o.OrderId,
			OwnCoin:     o.OwnCoin,
			AmountLimit: o.AmountLimit,
			NumLimit:    o.NumLimit,
			Price:       o.Price,
			Status:      o.Status,
			CreateTime:  o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetOwnCoinSubscribeOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
