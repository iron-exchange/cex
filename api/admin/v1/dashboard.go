package v1

import "github.com/gogf/gf/v2/frame/g"

type ChartData struct {
	Date     string  `json:"date"`
	Recharge float64 `json:"recharge"`
	Withdraw float64 `json:"withdraw"`
}

type GetDashboardReq struct {
	g.Meta `path:"/dashboard/index" tags:"AdminDashboard" method:"get" summary:"后台首页数据统计"`
}

type GetDashboardRes struct {
	TotalRevenue  float64     `json:"totalRevenue"`  // 平台总收入
	PlayerCount   int64       `json:"playerCount"`   // 玩家数量
	TotalRecharge float64     `json:"totalRecharge"` // 总充值金额
	TotalWithdraw float64     `json:"totalWithdraw"` // 提现金额
	ChartData     []ChartData `json:"chartData"`     // 图表数据
}
