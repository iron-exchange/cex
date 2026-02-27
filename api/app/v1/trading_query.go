package v1

import "github.com/gogf/gf/v2/frame/g"

// CurrencyOrderListReq
type CurrencyOrderListReq struct {
	g.Meta `path:"/currency/order/list" tags:"Trading" method:"get" summary:"获取现货持仓历史"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"查询特定交易对"`
	Status int    `json:"status" d:"-1" dc:"-1全部, 0委托中, 1已完成, 2已撤销"`
}

type CurrencyOrderInfo struct {
	OrderNo    string  `json:"orderNo"`
	Symbol     string  `json:"symbol"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	Type       int     `json:"type" dc:"0买 1卖"`
	Status     int     `json:"status"`
	CreateTime string  `json:"createTime"`
}

type CurrencyOrderListRes struct {
	List  []CurrencyOrderInfo `json:"list"`
	Total int                 `json:"total"`
}

// CurrencyOrderCancelReq
type CurrencyOrderCancelReq struct {
	g.Meta  `path:"/currency/order/cancel" tags:"Trading" method:"post" summary:"手工撤单"`
	OrderNo string `json:"orderNo" v:"required#订单编号不能为空"`
}

type CurrencyOrderCancelRes struct {
	Success bool `json:"success"`
}
