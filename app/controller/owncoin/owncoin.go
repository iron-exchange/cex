package owncoin

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/app/owncoin"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
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
	userId := gconv.Int64(middleware.Auth.GetIdentity(ctx))
	return owncoin.New().GetOwnCoinDetail(ctx, userId, req)
}

// SubscribeOwnCoin 新币申购打新下单
func (c *Controller) SubscribeOwnCoin(ctx context.Context, req *v1.SubscribeOwnCoinReq) (*v1.SubscribeOwnCoinRes, error) {
	userId := gconv.Int64(middleware.Auth.GetIdentity(ctx))
	return owncoin.New().SubscribeOwnCoin(ctx, userId, req)
}
