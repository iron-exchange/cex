package defi

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/defi"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetOrderList(ctx context.Context, req *v1.GetAdminDefiOrderListReq) (res *v1.GetAdminDefiOrderListRes, err error) {
	return defi.New().GetOrderList(ctx, req)
}

func (c *Controller) GetActivityList(ctx context.Context, req *v1.GetAdminDefiActivityListReq) (res *v1.GetAdminDefiActivityListRes, err error) {
	return defi.New().GetActivityList(ctx, req)
}

func (c *Controller) GetRateList(ctx context.Context, req *v1.GetAdminDefiRateListReq) (res *v1.GetAdminDefiRateListRes, err error) {
	return defi.New().GetRateList(ctx, req)
}
