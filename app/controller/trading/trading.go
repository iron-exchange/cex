package trading

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/trading"
	"GoCEX/internal/service/middleware"
)

// Controller v1 版本的交易引擎接口
type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// CurrencyOrderSubmit 现货下注撮合
func (c *Controller) CurrencyOrderSubmit(ctx context.Context, req *v1.CurrencyOrderSubmitReq) (res *v1.CurrencyOrderSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = trading.New().CurrencyOrderSubmit(ctx, userId, req)
	return &v1.CurrencyOrderSubmitRes{}, err
}

// SecondContractSubmit 秒合约下注 (自带杀客包赢包输 Buff 外挂)
func (c *Controller) SecondContractSubmit(ctx context.Context, req *v1.SecondContractSubmitReq) (res *v1.SecondContractSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = trading.New().SecondContractSubmit(ctx, userId, req)
	return &v1.SecondContractSubmitRes{}, err
}

// ContractOrderSubmit 提交永续合约订单
func (c *Controller) ContractOrderSubmit(ctx context.Context, req *v1.ContractOrderSubmitReq) (res *v1.ContractOrderSubmitRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = trading.New().ContractOrderSubmit(ctx, userId, req)
	return &v1.ContractOrderSubmitRes{}, err
}

// CurrencyOrderList 现货持仓历史记录查询
func (c *Controller) CurrencyOrderList(ctx context.Context, req *v1.CurrencyOrderListReq) (res *v1.CurrencyOrderListRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return trading.New().GetCurrencyOrders(ctx, req, userId)
}

// CurrencyOrderCancel 手工撤销指定挂单 (CAS)
func (c *Controller) CurrencyOrderCancel(ctx context.Context, req *v1.CurrencyOrderCancelReq) (res *v1.CurrencyOrderCancelRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return trading.New().CancelCurrencyOrder(ctx, req, userId)
}
