package second_contract

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/second_contract"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetSecondCoinConfigList秒合约币种配置
func (c *Controller) GetSecondCoinConfigList(ctx context.Context, req *v1.GetSecondCoinConfigListReq) (res *v1.GetSecondCoinConfigListRes, err error) {
	return second_contract.New().GetSecondCoinConfigList(ctx, req)
}

// GetSecondContractOrderList 秒合约订单记录
func (c *Controller) GetSecondContractOrderList(ctx context.Context, req *v1.GetSecondContractOrderListReq) (res *v1.GetSecondContractOrderListRes, err error) {
	return second_contract.New().GetSecondContractOrderList(ctx, req)
}
