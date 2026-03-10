package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 挖矿产品 (Mining Products) ---------
type AdminMiningProductInfo struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Icon        string  `json:"icon"`
	Status      int     `json:"status"` // 1开 0关
	Days        string  `json:"days"`
	DefaultOdds float64 `json:"defaultOdds"`
	MinOdds     float64 `json:"minOdds"`
	MaxOdds     float64 `json:"maxOdds"`
	TimeLimit   int     `json:"timeLimit"`
	LimitMin    float64 `json:"limitMin"`
	LimitMax    float64 `json:"limitMax"`
	Sort        int     `json:"sort"`
	BuyPurchase int     `json:"buyPurchase"`
	Coin        string  `json:"coin"`
	Remark      string  `json:"remark"`
	CreateTime  string  `json:"createTime"`
}

type GetAdminMiningProductListReq struct {
	g.Meta `path:"/mining/product/list" tags:"AdminMining" method:"get" summary:"获取挖矿产品列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Title  string `json:"title" dc:"产品名称"`
	Status *int   `json:"status" dc:"状态"`
}

type GetAdminMiningProductListRes struct {
	List  []AdminMiningProductInfo `json:"list"`
	Total int                      `json:"total"`
}

// --------- 订单列表 (Mining Orders) ---------
type AdminMiningOrderInfo struct {
	Id           int     `json:"id"`
	OrderNo      string  `json:"orderNo"`
	UserId       int64   `json:"userId"`
	PlanId       int64   `json:"planId"`
	PlanTitle    string  `json:"planTitle"`
	Amount       float64 `json:"amount"` // 分
	OrderAmount  float64 `json:"orderAmount"`
	Days         int     `json:"days"`
	Status       int     `json:"status"` // 0 收益  1 结算
	AccumulaEarn float64 `json:"accumulaEarn"`
	MinOdds      float64 `json:"minOdds"`
	MaxOdds      float64 `json:"maxOdds"`
	CreateTime   string  `json:"createTime"`
	EndTime      string  `json:"endTime"`
	SettleTime   string  `json:"settleTime"`
}

type GetAdminMiningOrderListReq struct {
	g.Meta  `path:"/mining/order/list" tags:"AdminMining" method:"get" summary:"获取用户质押挖矿订单"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetAdminMiningOrderListRes struct {
	List  []AdminMiningOrderInfo `json:"list"`
	Total int                    `json:"total"`
}
