package swap

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/swap"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetSymbolManageList(ctx context.Context, req *v1.GetAdminSymbolManageListReq) (res *v1.GetAdminSymbolManageListRes, err error) {
	return swap.New().GetSymbolManageList(ctx, req)
}

func (c *Controller) GetExchangeCoinRecordList(ctx context.Context, req *v1.GetAdminExchangeCoinRecordListReq) (res *v1.GetAdminExchangeCoinRecordListRes, err error) {
	return swap.New().GetExchangeCoinRecordList(ctx, req)
}
