package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 理财产品 (Mine Financial) ---------
type AdminFinancialProductInfo struct {
	Id                int     `json:"id"`
	Title             string  `json:"title"`
	Icon              string  `json:"icon"`
	Status            int     `json:"status"` // 1开 0关
	Days              string  `json:"days"`
	MinOdds           float64 `json:"minOdds"`
	MaxOdds           float64 `json:"maxOdds"`
	TimeLimit         int     `json:"timeLimit"`
	LimitMin          float64 `json:"limitMin"`
	LimitMax          float64 `json:"limitMax"`
	Sort              int     `json:"sort"`
	BuyPurchase       int     `json:"buyPurchase"`
	AvgRate           float64 `json:"avgRate"`
	Coin              string  `json:"coin"`
	Classify          string  `json:"classify"` // 0 普通 1 vip 2 增值
	Level             int     `json:"level"`
	TotalInvestAmount float64 `json:"totalInvestAmount"`
	Process           float64 `json:"process"`
	RemainAmount      float64 `json:"remainAmount"`
	CreateTime        string  `json:"createTime"`
}

type GetAdminFinancialProductListReq struct {
	g.Meta `path:"/financial/product/list" tags:"AdminFinancial" method:"get" summary:"获取理财产品配置列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Title  string `json:"title" dc:"产品名称"`
	Status *int   `json:"status" dc:"状态"`
}

type GetAdminFinancialProductListRes struct {
	List  []AdminFinancialProductInfo `json:"list"`
	Total int                         `json:"total"`
}

// --------- 理财订单 (Mine Order) ---------
type AdminFinancialOrderInfo struct {
	Id           int     `json:"id"`
	OrderNo      string  `json:"orderNo"`
	UserId       int64   `json:"userId"`
	PlanId       int64   `json:"planId"`
	PlanTitle    string  `json:"planTitle"`
	Amount       float64 `json:"amount"` // 分
	OrderAmount  float64 `json:"orderAmount"`
	Days         int     `json:"days"`
	Status       int     `json:"status"` // 0 收益 1 结算
	AccumulaEarn float64 `json:"accumulaEarn"`
	MinOdds      float64 `json:"minOdds"`
	MaxOdds      float64 `json:"maxOdds"`
	Coin         string  `json:"coin"`
	AvgRate      float64 `json:"avgRate"`
	CreateTime   string  `json:"createTime"`
	EndTime      string  `json:"endTime"`
	SettleTime   string  `json:"settleTime"`
}

type GetAdminFinancialOrderListReq struct {
	g.Meta  `path:"/financial/order/list" tags:"AdminFinancial" method:"get" summary:"获取用户理财订单记录"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetAdminFinancialOrderListRes struct {
	List  []AdminFinancialOrderInfo `json:"list"`
	Total int                       `json:"total"`
}
