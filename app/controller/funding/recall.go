package funding

import (
	"context"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/funding"
)

// UncCallback 优盾代收网关回调
func (c *Controller) UncCallback(ctx context.Context, req *v1.UncCallbackReq) (res *v1.UncCallbackRes, err error) {
	// 直接交由底层验签，成功或失败都以 error 返回
	err = funding.NewRecall().UncCallback(ctx, req)
	return &v1.UncCallbackRes{}, err
}
