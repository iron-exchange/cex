package wallet_record

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/wallet_record"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetWalletRecordList 归集订单查询
func (c *Controller) GetWalletRecordList(ctx context.Context, req *v1.GetWalletRecordListReq) (res *v1.GetWalletRecordListRes, err error) {
	return wallet_record.New().GetWalletRecordList(ctx, req)
}
