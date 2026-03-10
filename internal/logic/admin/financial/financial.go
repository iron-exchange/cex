package financial

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminFinancial struct{}

func New() *sAdminFinancial {
	return &sAdminFinancial{}
}

// GetProductList 获取理财产品列表
func (s *sAdminFinancial) GetProductList(ctx context.Context, req *v1.GetAdminFinancialProductListReq) (*v1.GetAdminFinancialProductListRes, error) {
	m := dao.MineFinancial.Ctx(ctx)
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.MineFinancial
	err := m.Page(req.Page, req.Size).OrderDesc("sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminFinancialProductInfo, 0, len(list))
	for _, p := range list {
		resList = append(resList, v1.AdminFinancialProductInfo{
			Id:                p.Id,
			Title:             p.Title,
			Icon:              p.Icon,
			Status:            p.Status,
			Days:              p.Days,
			MinOdds:           p.MinOdds,
			MaxOdds:           p.MaxOdds,
			TimeLimit:         p.TimeLimit,
			LimitMin:          p.LimitMin,
			LimitMax:          p.LimitMax,
			Sort:              p.Sort,
			BuyPurchase:       p.BuyPurchase,
			AvgRate:           p.AvgRate,
			Coin:              p.Coin,
			Classify:          p.Classify,
			Level:             p.Level,
			TotalInvestAmount: p.TotalInvestAmount,
			Process:           p.Process,
			RemainAmount:      p.RemainAmount,
			CreateTime:        p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminFinancialProductListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOrderList 获取用户理财订单
func (s *sAdminFinancial) GetOrderList(ctx context.Context, req *v1.GetAdminFinancialOrderListReq) (*v1.GetAdminFinancialOrderListRes, error) {
	m := dao.MineOrder.Ctx(ctx)
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
	var list []entity.MineOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminFinancialOrderInfo, 0, len(list))
	for _, o := range list {
		info := v1.AdminFinancialOrderInfo{
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
			Coin:         o.Coin,
			AvgRate:      o.AvgRate,
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

	return &v1.GetAdminFinancialOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
