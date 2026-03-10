package exchange

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/exchange"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetOwnCoinList(ctx context.Context, req *v1.GetExchangeOwnCoinListReq) (res *v1.GetExchangeOwnCoinListRes, err error) {
	return exchange.New().GetOwnCoinList(ctx, req)
}

func (c *Controller) AddOwnCoin(ctx context.Context, req *v1.AddExchangeOwnCoinReq) (res *v1.AddExchangeOwnCoinRes, err error) {
	return exchange.New().AddOwnCoin(ctx, req)
}

func (c *Controller) EditOwnCoin(ctx context.Context, req *v1.EditExchangeOwnCoinReq) (res *v1.EditExchangeOwnCoinRes, err error) {
	return exchange.New().EditOwnCoin(ctx, req)
}

func (c *Controller) DeleteOwnCoin(ctx context.Context, req *v1.DeleteExchangeOwnCoinReq) (res *v1.DeleteExchangeOwnCoinRes, err error) {
	return exchange.New().DeleteOwnCoin(ctx, req)
}
