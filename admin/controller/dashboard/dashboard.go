package dashboard

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/dashboard"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetDashboard 首页数据统计
func (c *Controller) GetDashboard(ctx context.Context, req *v1.GetDashboardReq) (res *v1.GetDashboardRes, err error) {
	return dashboard.New().GetDashboard(ctx, req)
}
