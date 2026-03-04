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
