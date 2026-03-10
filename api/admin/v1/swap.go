package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种管理 (Swap Coin Config) ---------
type AdminSymbolManageInfo struct {
	Id           int64   `json:"id"`
	Symbol       string  `json:"symbol"`
	MinChargeNum float64 `json:"minChargeNum"`
	MaxChargeNum float64 `json:"maxChargeNum"`
	Commission   float64 `json:"commission"`
	Sort         int     `json:"sort"`
	Enable       string  `json:"enable"` // 1 启用 2 禁用
	Logo         string  `json:"logo"`
	Market       string  `json:"market"`
	CreateTime   string  `json:"createTime"`
}

type GetAdminSymbolManageListReq struct {
	g.Meta `path:"/swap/symbol/list" tags:"AdminSwap" method:"get" summary:"获取闪兑币种配置列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"币种名称"`
}

type GetAdminSymbolManageListRes struct {
	List  []AdminSymbolManageInfo `json:"list"`
	Total int                     `json:"total"`
}

// --------- 兑换订单 (Swap Orders) ---------
type AdminExchangeCoinRecordInfo struct {
	Id         int64   `json:"id"`
	UserId     int     `json:"userId"`
	Username   string  `json:"username"`
	FromCoin   string  `json:"fromCoin"`
	ToCoin     string  `json:"toCoin"`
	Amount     float64 `json:"amount"`
	ThirdRate  float64 `json:"thirdRate"`
	SystemRate float64 `json:"systemRate"`
	Status     int     `json:"status"` // 0:已提交;1:成功;2失败
	CreateTime string  `json:"createTime"`
}

type GetAdminExchangeCoinRecordListReq struct {
	g.Meta `path:"/swap/order/list" tags:"AdminSwap" method:"get" summary:"获取用户的闪兑订单记录"`
	Page   int  `json:"page" d:"1"`
	Size   int  `json:"size" d:"20"`
	UserId int  `json:"userId" dc:"用户ID"`
	Status *int `json:"status" dc:"状态"`
}

type GetAdminExchangeCoinRecordListRes struct {
	List  []AdminExchangeCoinRecordInfo `json:"list"`
	Total int                           `json:"total"`
}
