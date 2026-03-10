package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 交易所 - 自发币配置 (Exchange Own Coin Config) ---------

type ExchangeOwnCoinInfo struct {
	Id            int64   `json:"id"`
	Coin          string  `json:"coin"`
	Logo          string  `json:"logo"`
	ReferCoin     string  `json:"referCoin"`
	ReferMarket   string  `json:"referMarket"`
	ShowSymbol    string  `json:"showSymbol"`
	Price         float64 `json:"price"`
	Proportion    float64 `json:"proportion"`
	RaisingAmount float64 `json:"raisingAmount"`
	RaisedAmount  float64 `json:"raisedAmount"`
	PurchaseLimit int     `json:"purchaseLimit"`
	TotalAmount   float64 `json:"totalAmount"`
	Status        int     `json:"status"` // 1.未发布  2.筹集中 3 筹集成功 4.筹集失败
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
	Introduce     string  `json:"introduce"`
	Remark        string  `json:"remark"`
	CreateTime    string  `json:"createTime"`
}

type GetExchangeOwnCoinListReq struct {
	g.Meta    `path:"/exchange/ownCoin/list" tags:"AdminExchange" method:"get" summary:"获取交易所自发币列表"`
	Page      int    `json:"page" d:"1"`
	Size      int    `json:"size" d:"20"`
	Coin      string `json:"coin" dc:"币种"`
	ReferCoin string `json:"referCoin" dc:"参考币种"`
}

type GetExchangeOwnCoinListRes struct {
	List  []ExchangeOwnCoinInfo `json:"list"`
	Total int                   `json:"total"`
}

type AddExchangeOwnCoinReq struct {
	g.Meta        `path:"/exchange/ownCoin/add" tags:"AdminExchange" method:"post" summary:"新增自发币"`
	Coin          string  `json:"coin" v:"required#币种不能为空"`
	Logo          string  `json:"logo"`
	ReferCoin     string  `json:"referCoin" v:"required#参考币种不能为空"`
	ReferMarket   string  `json:"referMarket"`
	ShowSymbol    string  `json:"showSymbol"`
	Price         float64 `json:"price"`
	Proportion    float64 `json:"proportion"`
	RaisingAmount float64 `json:"raisingAmount"`
	PurchaseLimit int     `json:"purchaseLimit"`
	TotalAmount   float64 `json:"totalAmount"`
	Status        int     `json:"status"`
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
	Introduce     string  `json:"introduce"`
	Remark        string  `json:"remark"`
}

type AddExchangeOwnCoinRes struct{}

type EditExchangeOwnCoinReq struct {
	g.Meta        `path:"/exchange/ownCoin/edit" tags:"AdminExchange" method:"post" summary:"修改自发币"`
	Id            int64   `json:"id" v:"required#ID不能为空"`
	Coin          string  `json:"coin" v:"required#币种不能为空"`
	Logo          string  `json:"logo"`
	ReferCoin     string  `json:"referCoin" v:"required#参考币种不能为空"`
	ReferMarket   string  `json:"referMarket"`
	ShowSymbol    string  `json:"showSymbol"`
	Price         float64 `json:"price"`
	Proportion    float64 `json:"proportion"`
	RaisingAmount float64 `json:"raisingAmount"`
	PurchaseLimit int     `json:"purchaseLimit"`
	TotalAmount   float64 `json:"totalAmount"`
	Status        int     `json:"status"`
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
	Introduce     string  `json:"introduce"`
	Remark        string  `json:"remark"`
}

type EditExchangeOwnCoinRes struct{}

type DeleteExchangeOwnCoinReq struct {
	g.Meta `path:"/exchange/ownCoin/delete" tags:"AdminExchange" method:"post" summary:"删除自发币"`
	Id     int64 `json:"id" v:"required#ID不能为空"`
}

type DeleteExchangeOwnCoinRes struct{}
