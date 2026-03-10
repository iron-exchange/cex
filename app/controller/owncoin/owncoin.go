package owncoin

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/app/owncoin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetOwnCoinList 获取自发币(新币)列表
func (c *Controller) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (*v1.GetOwnCoinListRes, error) {
	return owncoin.New().GetOwnCoinList(ctx, req)
}

// GetOwnCoinDetail 获取指定自发币全网与个人明细
func (c *Controller) GetOwnCoinDetail(ctx context.Context, req *v1.GetOwnCoinDetailReq) (*v1.GetOwnCoinDetailRes, error) {
	return owncoin.New().GetOwnCoinDetail(ctx, req)
}

// SubscribeOwnCoin 新币申购打新下单
func (c *Controller) SubscribeOwnCoin(ctx context.Context, req *v1.SubscribeOwnCoinReq) (*v1.SubscribeOwnCoinRes, error) {
	return owncoin.New().SubscribeOwnCoin(ctx, req)
}
