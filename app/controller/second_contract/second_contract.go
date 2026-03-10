package second_contract

import (
	"context"

	v1 "GoCEX/api/app/v1"
	secondLogic "GoCEX/internal/logic/app/second_contract"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetCoinList 查询币种与周期配置
func (c *Controller) GetCoinList(ctx context.Context, req *v1.SecondCoinListReq) (*v1.SecondCoinListRes, error) {
	return secondLogic.New().GetCoinList(ctx, req)
}

// GetCoinDetail 查询单个期权配置
func (c *Controller) GetCoinDetail(ctx context.Context, req *v1.SecondCoinDetailReq) (*v1.SecondCoinDetailRes, error) {
	return secondLogic.New().GetCoinDetail(ctx, req)
}

// CreateOrder 新增订单
func (c *Controller) CreateOrder(ctx context.Context, req *v1.CreateSecondOrderReq) (*v1.CreateSecondOrderRes, error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return secondLogic.New().CreateOrder(ctx, userId, req)
}

// SelectOrderList 查询订单
func (c *Controller) SelectOrderList(ctx context.Context, req *v1.SelectSecondOrderListReq) (*v1.SelectSecondOrderListRes, error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return secondLogic.New().SelectOrderList(ctx, userId, req)
}
