package mining

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/mining"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetProductList 挖矿产品表
func (c *Controller) GetProductList(ctx context.Context, req *v1.GetAdminMiningProductListReq) (res *v1.GetAdminMiningProductListRes, err error) {
	return mining.New().GetProductList(ctx, req)
}

// GetOrderList 质押挖矿订单
func (c *Controller) GetOrderList(ctx context.Context, req *v1.GetAdminMiningOrderListReq) (res *v1.GetAdminMiningOrderListRes, err error) {
	return mining.New().GetOrderList(ctx, req)
}
