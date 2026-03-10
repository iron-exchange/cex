package bank

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/bank"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetBankList 银行卡列表查询
func (c *Controller) GetBankList(ctx context.Context, req *v1.GetBankListReq) (res *v1.GetBankListRes, err error) {
	return bank.New().GetBankList(ctx, req)
}
