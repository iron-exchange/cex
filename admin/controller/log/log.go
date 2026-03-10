package log

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/log"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetLoginLogList 登录日志查询
func (c *Controller) GetLoginLogList(ctx context.Context, req *v1.GetLoginLogListReq) (res *v1.GetLoginLogListRes, err error) {
	return log.New().GetLoginLogList(ctx, req)
}
