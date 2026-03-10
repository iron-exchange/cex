package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种配置 (Currency Symbol Config) ---------
type CurrencySymbolInfo struct {
	Id            int64   `json:"id"`
	Symbol        string  `json:"symbol"`
	ShowSymbol    string  `json:"showSymbol"`
	Coin          string  `json:"coin"`
	BaseCoin      string  `json:"baseCoin"`
	FeeRate       float64 `json:"feeRate"`
	CoinPrecision int     `json:"coinPrecision"`
	BasePrecision int     `json:"basePrecision"`
	SellMin       float64 `json:"sellMin"`
	BuyMax        float64 `json:"buyMax"`
	OrderMin      float64 `json:"orderMin"`
	OrderMax      float64 `json:"orderMax"`
	Enable        string  `json:"enable"`      // 1=启用 2=禁用
	IsShow        string  `json:"isShow"`      // 1=显示 2=隐藏
	IsDeal        string  `json:"isDeal"`      // 1=能交易 2=否
	MarketBuy     string  `json:"marketBuy"`   // 1=可以 2=不可
	MarketSell    string  `json:"marketSell"`  // 1=可以 2=不可
	LimitedBuy    string  `json:"limitedBuy"`  // 1=可以 2=不可
	LimitedSell   string  `json:"limitedSell"` // 1=可以 2=不可
	Logo          string  `json:"logo"`
	Market        string  `json:"market"`
	CreateTime    string  `json:"createTime"`
}

type GetCurrencySymbolListReq struct {
	g.Meta `path:"/currency/symbol/list" tags:"AdminCurrencyTrading" method:"get" summary:"获取币币交易对配置列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"交易对名称"`
}

type GetCurrencySymbolListRes struct {
	List  []CurrencySymbolInfo `json:"list"`
	Total int                  `json:"total"`
}

// --------- 币币交易订单 (Currency Orders) ---------
type AdminCurrencyOrderInfo struct {
	Id            int64   `json:"id"`
	UserId        int64   `json:"userId"`
	Type          int     `json:"type"`         // 0 买入 1卖出
	DelegateType  int     `json:"delegateType"` // 0 限价 1 市价
	Status        int     `json:"status"`       // 0 等待成交 1 完全成交 3 已撤销
	OrderNo       string  `json:"orderNo"`
	Symbol        string  `json:"symbol"`
	Coin          string  `json:"coin"`
	DelegateTotal float64 `json:"delegateTotal"`
	DelegatePrice float64 `json:"delegatePrice"`
	DelegateValue float64 `json:"delegateValue"`
	DealNum       float64 `json:"dealNum"`
	DealPrice     float64 `json:"dealPrice"`
	DealValue     float64 `json:"dealValue"`
	Fee           float64 `json:"fee"`
	CreateTime    string  `json:"createTime"`
}

type GetAdminCurrencyOrderListReq struct {
	g.Meta  `path:"/currency/order/list" tags:"AdminCurrencyTrading" method:"get" summary:"获取币币现货交易订单"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetAdminCurrencyOrderListRes struct {
	List  []AdminCurrencyOrderInfo `json:"list"`
	Total int                      `json:"total"`
}
