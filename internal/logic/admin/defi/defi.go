package defi

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminDefi struct{}

func New() *sAdminDefi {
	return &sAdminDefi{}
}

// GetOrderList 获取DEFI凭证持仓列表
func (s *sAdminDefi) GetOrderList(ctx context.Context, req *v1.GetAdminDefiOrderListReq) (*v1.GetAdminDefiOrderListRes, error) {
	m := dao.DefiOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}

	total, _ := m.Count()
	var list []entity.DefiOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminDefiOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.AdminDefiOrderInfo{
			Id:          o.Id,
			UserId:      o.UserId,
			Amount:      o.Amount,
			TotleAmount: o.TotleAmount,
			Rate:        o.Rate,
			CreateTime:  o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminDefiOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetActivityList 获取空投奖励记录
func (s *sAdminDefi) GetActivityList(ctx context.Context, req *v1.GetAdminDefiActivityListReq) (*v1.GetAdminDefiActivityListRes, error) {
	m := dao.DefiActivity.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.DefiActivity
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminDefiActivityInfo, 0, len(list))
	for _, a := range list {
		info := v1.AdminDefiActivityInfo{
			Id:          a.Id,
			UserId:      a.UserId,
			TotleAmount: a.TotleAmount,
			Amount:      a.Amount,
			Type:        a.Type,
			Status:      a.Status,
			CreateTime:  a.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if a.EndTime != nil {
			info.EndTime = a.EndTime.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}

	return &v1.GetAdminDefiActivityListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetRateList 获取挖矿档次配置表
func (s *sAdminDefi) GetRateList(ctx context.Context, req *v1.GetAdminDefiRateListReq) (*v1.GetAdminDefiRateListRes, error) {
	m := dao.DefiRate.Ctx(ctx)

	total, _ := m.Count()
	var list []entity.DefiRate
	err := m.Page(req.Page, req.Size).OrderAsc("min_amount").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminDefiRateInfo, 0, len(list))
	for _, r := range list {
		resList = append(resList, v1.AdminDefiRateInfo{
			Id:         r.Id,
			Symbol:     r.Symbol,
			RewardCoin: r.RewardCoin,
			MinAmount:  r.MinAmount,
			MaxAmount:  r.MaxAmount,
			Rate:       r.Rate,
			CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminDefiRateListRes{
		List:  resList,
		Total: total,
	}, nil
}
