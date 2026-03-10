package financial

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/financial"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetProductList(ctx context.Context, req *v1.GetAdminFinancialProductListReq) (res *v1.GetAdminFinancialProductListRes, err error) {
	return financial.New().GetProductList(ctx, req)
}

func (c *Controller) GetOrderList(ctx context.Context, req *v1.GetAdminFinancialOrderListReq) (res *v1.GetAdminFinancialOrderListRes, err error) {
	return financial.New().GetOrderList(ctx, req)
}
