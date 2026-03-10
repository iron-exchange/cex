package address

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/address"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetAddressAuthList 授权地址列表查询
func (c *Controller) GetAddressAuthList(ctx context.Context, req *v1.GetAddressAuthListReq) (res *v1.GetAddressAuthListRes, err error) {
	return address.New().GetAddressAuthList(ctx, req)
}
