package loan

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/loan"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetProductList(ctx context.Context, req *v1.GetAdminLoadProductListReq) (res *v1.GetAdminLoadProductListRes, err error) {
	return loan.New().GetProductList(ctx, req)
}

func (c *Controller) GetOrderList(ctx context.Context, req *v1.GetAdminLoadOrderListReq) (res *v1.GetAdminLoadOrderListRes, err error) {
	return loan.New().GetOrderList(ctx, req)
}
