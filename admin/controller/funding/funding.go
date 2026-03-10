package funding

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/funding"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetRechargeList 充值记录查询
func (c *Controller) GetRechargeList(ctx context.Context, req *v1.GetRechargeListReq) (res *v1.GetRechargeListRes, err error) {
	return funding.New().GetRechargeList(ctx, req)
}

// GetWithdrawList 提现记录查询
func (c *Controller) GetWithdrawList(ctx context.Context, req *v1.GetWithdrawListReq) (res *v1.GetWithdrawListRes, err error) {
	return funding.New().GetWithdrawList(ctx, req)
}

// GetRechargeChannelList 充值通道配置
func (c *Controller) GetRechargeChannelList(ctx context.Context, req *v1.GetRechargeChannelListReq) (res *v1.GetRechargeChannelListRes, err error) {
	return funding.New().GetRechargeChannelList(ctx, req)
}

// GetWithdrawChannelList 提现通道配置
func (c *Controller) GetWithdrawChannelList(ctx context.Context, req *v1.GetWithdrawChannelListReq) (res *v1.GetWithdrawChannelListRes, err error) {
	return funding.New().GetWithdrawChannelList(ctx, req)
}
