package report

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/report"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetDailyData 每日数据统计
func (c *Controller) GetDailyData(ctx context.Context, req *v1.GetDailyDataReq) (res *v1.GetDailyDataRes, err error) {
	return report.New().GetDailyData(ctx, req)
}

// GetAgentData 代理数据统计
func (c *Controller) GetAgentData(ctx context.Context, req *v1.GetAgentDataReq) (res *v1.GetAgentDataRes, err error) {
	return report.New().GetAgentData(ctx, req)
}

// GetPlayerData 单玩家历史统计
func (c *Controller) GetPlayerData(ctx context.Context, req *v1.GetPlayerDataReq) (res *v1.GetPlayerDataRes, err error) {
	return report.New().GetPlayerData(ctx, req)
}
