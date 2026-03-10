package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 每日数据 (Daily Data) ---------
type DailyDataInfo struct {
	Date          string  `json:"date"`
	NewUsers      int     `json:"newUsers"`
	TotalRecharge float64 `json:"totalRecharge"`
	TotalWithdraw float64 `json:"totalWithdraw"`
	CompanyProfit float64 `json:"companyProfit"`
}

type GetDailyDataReq struct {
	g.Meta    `path:"/report/daily" tags:"AdminReport" method:"get" summary:"每日数据统计"`
	Page      int    `json:"page" d:"1"`
	Size      int    `json:"size" d:"20"`
	StartDate string `json:"startDate" dc:"开始日期"`
	EndDate   string `json:"endDate" dc:"结束日期"`
}

type GetDailyDataRes struct {
	List  []DailyDataInfo `json:"list"`
	Total int             `json:"total"`
}

// --------- 代理数据 (Agent Data) ---------
type AgentDataInfo struct {
	AgentId       int64   `json:"agentId"`
	AgentName     string  `json:"agentName"`
	SubUsersCount int     `json:"subUsersCount"`
	TotalRecharge float64 `json:"totalRecharge"`
	TotalWithdraw float64 `json:"totalWithdraw"`
	AgentProfit   float64 `json:"agentProfit"`
}

type GetAgentDataReq struct {
	g.Meta  `path:"/report/agent" tags:"AdminReport" method:"get" summary:"代理数据统计"`
	Page    int   `json:"page" d:"1"`
	Size    int   `json:"size" d:"20"`
	AgentId int64 `json:"agentId" dc:"查询特定代理"`
}

type GetAgentDataRes struct {
	List  []AgentDataInfo `json:"list"`
	Total int             `json:"total"`
}

// --------- 玩家数据 (Player Data) ---------
type PlayerDataInfo struct {
	UserId         int64   `json:"userId"`
	LoginName      string  `json:"loginName"`
	TotalRecharge  float64 `json:"totalRecharge"`
	TotalWithdraw  float64 `json:"totalWithdraw"`
	NetProfit      float64 `json:"netProfit"`
	CurrentBalance float64 `json:"currentBalance"`
}

type GetPlayerDataReq struct {
	g.Meta `path:"/report/player" tags:"AdminReport" method:"get" summary:"单玩家盈亏统计数据"`
	Page   int   `json:"page" d:"1"`
	Size   int   `json:"size" d:"20"`
	UserId int64 `json:"userId" dc:"特定玩家ID查询"`
}

type GetPlayerDataRes struct {
	List  []PlayerDataInfo `json:"list"`
	Total int              `json:"total"`
}
