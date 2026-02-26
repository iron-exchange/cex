package asset

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/asset"
)

// Controller admin 端的资产接口 (供运营人员调整资金)
type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SubAmount 后台人工上下分 / 强制资金划转
func (c *Controller) SubAmount(ctx context.Context, req *v1.SubAmountReq) (res *v1.SubAmountRes, err error) {
	// 直接调用 asset logic 中的增减逻辑
	res, err = asset.New().SubAmount(ctx, req)
	return
}
