package defi

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/defi"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetDefiRate 获取分红收益列表
func (c *Controller) GetDefiRate(ctx context.Context, req *v1.GetDefiRateReq) (res *v1.GetDefiRateRes, err error) {
	return defi.New().GetDefiRate(ctx)
}

// SendApproveHash 处理用户提交的授权 Hash
func (c *Controller) SendApproveHash(ctx context.Context, req *v1.SendApproveHashReq) (res *v1.SendApproveHashRes, err error) {
	err = defi.New().SendApproveHash(ctx, req)
	return &v1.SendApproveHashRes{}, err
}

// ShowIncome 展示每日分红历史
func (c *Controller) ShowIncome(ctx context.Context, req *v1.ShowIncomeReq) (res *v1.ShowIncomeRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return defi.New().ShowIncome(ctx, userId)
}

// ShowOrder 获取用户正进行的 DeFi 授权订单
func (c *Controller) ShowOrder(ctx context.Context, req *v1.ShowOrderReq) (res *v1.ShowOrderRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return defi.New().ShowOrder(ctx, userId)
}
