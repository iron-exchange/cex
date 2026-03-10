package loan

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminLoan struct{}

func New() *sAdminLoan {
	return &sAdminLoan{}
}

// GetProductList 获取借贷产品列表
func (s *sAdminLoan) GetProductList(ctx context.Context, req *v1.GetAdminLoadProductListReq) (*v1.GetAdminLoadProductListRes, error) {
	m := dao.LoadProduct.Ctx(ctx)
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.LoadProduct
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminLoadProductInfo, 0, len(list))
	for _, p := range list {
		resList = append(resList, v1.AdminLoadProductInfo{
			Id:         p.Id,
			AmountMin:  p.AmountMin,
			AmountMax:  p.AmountMax,
			CycleType:  p.CycleType,
			RepayType:  p.RepayType,
			Status:     p.Status,
			Odds:       p.Odds,
			RepayOrg:   p.RepayOrg,
			IsFreeze:   p.IsFreeze,
			CreateTime: p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminLoadProductListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOrderList 获取借贷订单列表
func (s *sAdminLoan) GetOrderList(ctx context.Context, req *v1.GetAdminLoadOrderListReq) (*v1.GetAdminLoadOrderListRes, error) {
	m := dao.LoadOrder.Ctx(ctx)
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
	var list []entity.LoadOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminLoadOrderInfo, 0, len(list))
	for _, o := range list {
		info := v1.AdminLoadOrderInfo{
			Id:             o.Id,
			OrderNo:        o.OrderNo,
			UserId:         o.UserId,
			ProId:          o.ProId,
			Amount:         o.Amount,
			Rate:           o.Rate,
			Interest:       o.Interest,
			Status:         o.Status,
			CycleType:      o.CycleType,
			DisburseAmount: o.DisburseAmount,
			CreateTime:     o.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if o.FinalRepayTime != nil {
			info.FinalRepayTime = o.FinalRepayTime.Format("2006-01-02 15:04:05")
		}
		if o.DisburseTime != nil {
			info.DisburseTime = o.DisburseTime.Format("2006-01-02 15:04:05")
		}
		if o.ReturnTime != nil {
			info.ReturnTime = o.ReturnTime.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}

	return &v1.GetAdminLoadOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
