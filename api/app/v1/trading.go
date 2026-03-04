package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CurrencyOrderSubmitReq 现货挂单
type CurrencyOrderSubmitReq struct {
	g.Meta        `path:"/trading/currency_submit" tags:"Trading" method:"post" summary:"现货挂单"`
	Symbol        string  `json:"symbol" v:"required#请输入交易对"`
	Type          int     `json:"type" v:"required|in:0,1#交易方向错误"`         // 0 买入 1 卖出
	DelegateType  int     `json:"delegateType" v:"required|in:0,1#委托类型错误"` // 0 限价 1 市价
	DelegateTotal float64 `json:"delegateTotal" v:"required#请输入委托数量/总额"`
	DelegatePrice float64 `json:"delegatePrice"`
}
type CurrencyOrderSubmitRes struct{}

// SecondContractSubmitReq 秒合约挂单
type SecondContractSubmitReq struct {
	g.Meta     `path:"/trading/second_contract_submit" tags:"Trading" method:"post" summary:"秒合约挂单"`
	Symbol     string  `json:"symbol" v:"required#请输入交易对"`
	BetAmount  float64 `json:"betAmount" v:"required#请输入下注金额"`
	BetContent string  `json:"betContent" v:"required#请输入方向"`
	Period     int64   `json:"period" v:"required#请输入周期"`
}
type SecondContractSubmitRes struct{}

// ContractOrderSubmitReq 永续合约开仓
type ContractOrderSubmitReq struct {
	g.Meta        `path:"/trading/contract_submit" tags:"Trading" method:"post" summary:"合约开仓"`
	Symbol        string  `json:"symbol" v:"required#请输入交易对"`
	Type          int     `json:"type" v:"required|in:0,1#做多做空方向错误"`
	DelegateType  int     `json:"delegateType" v:"required|in:0,1#委托类型错误"`
	Leverage      float64 `json:"leverage" v:"required#请输入杠杆倍数"`
	DelegateTotal float64 `json:"delegateTotal" v:"required#请输入数量"`
	DelegatePrice float64 `json:"delegatePrice"`
}
type ContractOrderSubmitRes struct{}

// CurrencyOrderCancelReq 撤单
type CurrencyOrderCancelReq struct {
	g.Meta  `path:"/currency/order/cancel" tags:"Trading" method:"post" summary:"手工撤单"`
	OrderNo string `json:"orderNo" v:"required#订单编号不能为空"`
}
type CurrencyOrderCancelRes struct {
	Success bool `json:"success"`
}

// CurrencyOrderListReq 现货持仓历史查询
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
