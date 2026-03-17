package v1

import "github.com/gogf/gf/v2/frame/g"

type GetDashboardReq struct {
	g.Meta `path:"/dashboard/index" tags:"AdminDashboard" method:"get" summary:"后台首页数据统计"`
}

type DashboardSeries struct {
	Title        int                `json:"title" dc:"统计编号(1-4)"`
	TotalNum     float64            `json:"totalNum" dc:"累计总数"`
	RedLineName  string             `json:"redLineName" dc:"红线名称"`
	BlueLineName string             `json:"blueLineName" dc:"蓝线名称"`
	RedLine      map[string]float64 `json:"redLine" dc:"红线数据(日期:值)"`
	BlueLine     map[string]float64 `json:"blueLine" dc:"蓝线数据(日期:值)"`
}

type GetDashboardRes []DashboardSeries
