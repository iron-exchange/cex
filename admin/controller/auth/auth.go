package auth

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/auth"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Login(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	return auth.New().Login(ctx, req)
}

func (c *Controller) GetInfo(ctx context.Context, req *v1.AdminGetInfoReq) (res *v1.AdminGetInfoRes, err error) {
	return auth.New().GetInfo(ctx, req)
}

func (c *Controller) GetRouters(ctx context.Context, req *v1.AdminGetRoutersReq) (res *v1.AdminGetRoutersRes, err error) {
	return auth.New().GetRouters(ctx, req)
}
