package collection

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/collection"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetCollectionOrderList 归集订单查询
func (c *Controller) GetCollectionOrderList(ctx context.Context, req *v1.GetCollectionOrderListReq) (res *v1.GetCollectionOrderListRes, err error) {
	return collection.New().GetCollectionOrderList(ctx, req)
}
