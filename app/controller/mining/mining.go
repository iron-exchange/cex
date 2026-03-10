package mining

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/mining"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// PersonalIncome 展示个人理财总收益详情
func (c *Controller) PersonalIncome(ctx context.Context, req *v1.PersonalIncomeReq) (res *v1.PersonalIncomeRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return mining.New().GetFinancialIncome(ctx, userId)
}

// FinancialSubmit 购买理财产品
func (c *Controller) FinancialSubmit(ctx context.Context, req *v1.FinancialSubmitReq) (res *v1.FinancialSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = mining.New().BuyFinancial(ctx, userId, req)
	return &v1.FinancialSubmitRes{}, err
}

// MiningShow 展示矿机运行与可用产品
func (c *Controller) MiningShow(ctx context.Context, req *v1.MiningShowReq) (res *v1.MiningShowRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return mining.New().ShowMining(ctx, userId)
}

// MiningSubmit 申购矿机
func (c *Controller) MiningSubmit(ctx context.Context, req *v1.MiningSubmitReq) (res *v1.MiningSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = mining.New().BuyMining(ctx, userId, req)
	return &v1.MiningSubmitRes{}, err
}

// MiningRedemption 提前赎回矿机
func (c *Controller) MiningRedemption(ctx context.Context, req *v1.MiningRedemptionReq) (res *v1.MiningRedemptionRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = mining.New().Redemption(ctx, userId, req.OrderNo)
	return &v1.MiningRedemptionRes{Success: err == nil}, err
}

// MingProductList 单独查询可用矿机列表
func (c *Controller) MingProductList(ctx context.Context, req *v1.MingProductListReq) (res *v1.MingProductListRes, err error) {
	return mining.New().GetProductList(ctx, req)
}

// MingOrderList 单独查询我的订单
func (c *Controller) MingOrderList(ctx context.Context, req *v1.MingOrderListReq) (res *v1.MingOrderListRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return mining.New().GetOrderList(ctx, userId, req)
}

// MingOrderDetail 查询单笔订单详情
func (c *Controller) MingOrderDetail(ctx context.Context, req *v1.MingOrderDetailReq) (res *v1.MingOrderDetailRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return mining.New().GetOrderDetail(ctx, userId, req)
}

// MingOrderRedemptionNew 特殊标识提前赎回 (Wallet等客户端不同标识)
func (c *Controller) MingOrderRedemptionNew(ctx context.Context, req *v1.MingOrderRedemptionNewReq) (res *v1.MingOrderRedemptionNewRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = mining.New().Redemption(ctx, userId, req.OrderNo)
	return &v1.MingOrderRedemptionNewRes{Success: err == nil}, err
}
