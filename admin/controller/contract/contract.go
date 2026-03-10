package contract

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/contract"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetContractCoinList U本位合约币种配置
func (c *Controller) GetContractCoinList(ctx context.Context, req *v1.GetContractCoinListReq) (res *v1.GetContractCoinListRes, err error) {
	return contract.New().GetContractCoinList(ctx, req)
}

// GetContractOrderList 合约历史委托
func (c *Controller) GetContractOrderList(ctx context.Context, req *v1.GetContractOrderListReq) (res *v1.GetContractOrderListRes, err error) {
	return contract.New().GetContractOrderList(ctx, req)
}

// GetContractPositionList 当前合约持仓
func (c *Controller) GetContractPositionList(ctx context.Context, req *v1.GetContractPositionListReq) (res *v1.GetContractPositionListRes, err error) {
	return contract.New().GetContractPositionList(ctx, req)
}
