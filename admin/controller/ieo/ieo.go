package ieo

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/ieo"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetOwnCoinList 币种列表
func (c *Controller) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (res *v1.GetOwnCoinListRes, err error) {
	return ieo.New().GetOwnCoinList(ctx, req)
}

// GetOwnCoinSubscribeOrderList 申购订单查询
func (c *Controller) GetOwnCoinSubscribeOrderList(ctx context.Context, req *v1.GetOwnCoinSubscribeOrderListReq) (res *v1.GetOwnCoinSubscribeOrderListRes, err error) {
	return ieo.New().GetOwnCoinSubscribeOrderList(ctx, req)
}
