package collection

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminCollection struct{}

func New() *sAdminCollection {
	return &sAdminCollection{}
}

func (s *sAdminCollection) GetCollectionOrderList(ctx context.Context, req *v1.GetCollectionOrderListReq) (*v1.GetCollectionOrderListRes, error) {
	m := dao.CollectionOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderId != "" {
		m = m.Where("order_id", req.OrderId)
	}
	if req.Address != "" {
		m = m.Where("address", req.Address)
	}
	if req.Chain != "" {
		m = m.Where("chain", req.Chain)
	}
	if req.Coin != "" {
		m = m.Where("coin", req.Coin)
	}
	if req.Hash != "" {
		m = m.Where("hash", req.Hash)
	}
	if req.Amount > 0 {
		m = m.Where("amount", req.Amount)
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}
	if req.ClientName != "" {
		m = m.WhereLike("client_name", "%"+req.ClientName+"%")
	}
	if req.SearchValue != "" {
		m = m.WhereLike("search_value", "%"+req.SearchValue+"%")
	}

	total, _ := m.Count()
	var list []entity.CollectionOrder
	err := m.Page(req.PageNum, req.PageSize).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resRows := make([]v1.CollectionOrderInfo, 0, len(list))
	for _, c := range list {
		resRows = append(resRows, v1.CollectionOrderInfo{
			Id:          c.Id,
			OrderId:     c.OrderId,
			UserId:      c.UserId,
			Address:     c.Address,
			Chain:       c.Chain,
			Coin:        c.Coin,
			Hash:        c.Hash,
			Amount:      c.Amount,
			Status:      c.Status,
			ClientName:  c.ClientName,
			CreateTime:  c.CreateTime,
			CreateBy:    c.CreateBy,
			UpdateTime:  c.UpdateTime,
			UpdateBy:    c.UpdateBy,
			Remark:      c.Remark,
			SearchValue: c.SearchValue,
		})
	}

	return &v1.GetCollectionOrderListRes{
		Rows:  resRows,
		Total: total,
	}, nil
}
