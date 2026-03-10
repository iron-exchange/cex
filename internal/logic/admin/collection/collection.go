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
	if req.TxHash != "" {
		m = m.WhereLike("hash", "%"+req.TxHash+"%")
	}
	// Status in DB is a string (e.g., "1", "2")
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.CollectionOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.CollectionOrderInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.CollectionOrderInfo{
			Id:         c.Id,
			UserId:     c.UserId,
			Amount:     c.Amount,
			Coin:       c.Coin,
			Address:    c.Address,
			TxHash:     c.Hash,
			Status:     c.Status,
			CreateTime: c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetCollectionOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
