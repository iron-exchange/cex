package mining

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminMining struct{}

func New() *sAdminMining {
	return &sAdminMining{}
}

// GetProductList 获取挖矿理财产品
func (s *sAdminMining) GetProductList(ctx context.Context, req *v1.GetAdminMiningProductListReq) (*v1.GetAdminMiningProductListRes, error) {
	m := dao.MingProduct.Ctx(ctx)
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.MingProduct
	err := m.Page(req.Page, req.Size).OrderDesc("sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminMiningProductInfo, 0, len(list))
	for _, p := range list {
		resList = append(resList, v1.AdminMiningProductInfo{
			Id:          p.Id,
			Title:       p.Title,
			Icon:        p.Icon,
			Status:      p.Status,
			Days:        p.Days,
			DefaultOdds: p.DefaultOdds,
			MinOdds:     p.MinOdds,
			MaxOdds:     p.MaxOdds,
			TimeLimit:   p.TimeLimit,
			LimitMin:    p.LimitMin,
			LimitMax:    p.LimitMax,
			Sort:        p.Sort,
			BuyPurchase: p.BuyPurchase,
			Coin:        p.Coin,
			Remark:      p.Remark,
			CreateTime:  p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminMiningProductListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOrderList 获取用户挖矿理财订单
func (s *sAdminMining) GetOrderList(ctx context.Context, req *v1.GetAdminMiningOrderListReq) (*v1.GetAdminMiningOrderListRes, error) {
	m := dao.MingOrder.Ctx(ctx)
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
	var list []entity.MingOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminMiningOrderInfo, 0, len(list))
	for _, o := range list {
		info := v1.AdminMiningOrderInfo{
			Id:           o.Id,
			OrderNo:      o.OrderNo,
			UserId:       o.UserId,
			PlanId:       o.PlanId,
			PlanTitle:    o.PlanTitle,
			Amount:       o.Amount,
			OrderAmount:  o.OrderAmount,
			Days:         o.Days,
			Status:       o.Status,
			AccumulaEarn: o.AccumulaEarn,
			MinOdds:      o.MinOdds,
			MaxOdds:      o.MaxOdds,
			CreateTime:   o.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if o.EndTime != nil {
			info.EndTime = o.EndTime.Format("2006-01-02 15:04:05")
		}
		if o.SettleTime != nil {
			info.SettleTime = o.SettleTime.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}

	return &v1.GetAdminMiningOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
