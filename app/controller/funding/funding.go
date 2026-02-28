package funding

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/funding"
	"GoCEX/internal/service/middleware"
)

// Controller v1 版本的资金接口
type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// RechargeSubmit 申请充币
func (c *Controller) RechargeSubmit(ctx context.Context, req *v1.RechargeSubmitReq) (res *v1.RechargeSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = funding.New().RechargeSubmit(ctx, userId, req)
	return &v1.RechargeSubmitRes{}, err
}

// WithdrawSubmit 申请提现
func (c *Controller) WithdrawSubmit(ctx context.Context, req *v1.WithdrawSubmitReq) (res *v1.WithdrawSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = funding.New().WithdrawSubmit(ctx, userId, req)
	return &v1.WithdrawSubmitRes{}, err
}
