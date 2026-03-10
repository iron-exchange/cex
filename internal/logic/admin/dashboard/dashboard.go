package dashboard

import (
	"context"

	v1 "GoCEX/api/admin/v1"
)

type sAdminDashboard struct{}

func New() *sAdminDashboard {
	return &sAdminDashboard{}
}

func (s *sAdminDashboard) GetDashboard(ctx context.Context, req *v1.GetDashboardReq) (*v1.GetDashboardRes, error) {
	// TODO: 完整的首页数据由复杂的 SQL 聚合查询生成，此处先构建占位返回结构方便前端联调。
	// 玩家数量: SELECT count(1) FROM app_user
	// 充值金额: SELECT sum(amount) FROM app_recharge WHERE status = 2
	// 提现金额: SELECT sum(amount) FROM withdraw WHERE status = 2
	// 平台总收入: 可以是 (充值 - 提现) 或者是计算客损 (包输赢导致的利润)
	// 折线图数据: GROUP BY DATE(time) 聚合

	return &v1.GetDashboardRes{
		TotalRevenue:  35811038.0,
		PlayerCount:   630,
		TotalRecharge: 40442481.0,
		TotalWithdraw: 4631443.0,
		ChartData: []v1.ChartData{
			{Date: "2026-02-27", Recharge: 200, Withdraw: 100},
			{Date: "2026-02-28", Recharge: 1560, Withdraw: 300},
			{Date: "2026-03-01", Recharge: 2100, Withdraw: 321},
			{Date: "2026-03-02", Recharge: 2200, Withdraw: 310},
			{Date: "2026-03-03", Recharge: 15000, Withdraw: 321},
		},
	}, nil
}
