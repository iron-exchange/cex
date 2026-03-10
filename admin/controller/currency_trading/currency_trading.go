package currency_trading

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/currency_trading"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetCurrencySymbolList 币种配置
func (c *Controller) GetCurrencySymbolList(ctx context.Context, req *v1.GetCurrencySymbolListReq) (res *v1.GetCurrencySymbolListRes, err error) {
	return currency_trading.New().GetCurrencySymbolList(ctx, req)
}

// GetCurrencyOrderList 现货订单委托
func (c *Controller) GetCurrencyOrderList(ctx context.Context, req *v1.GetAdminCurrencyOrderListReq) (res *v1.GetAdminCurrencyOrderListRes, err error) {
	return currency_trading.New().GetCurrencyOrderList(ctx, req)
}
