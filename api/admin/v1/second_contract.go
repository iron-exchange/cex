package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种配置 (Second Coin Config) ---------
type SecondCoinConfigInfo struct {
	Id         int64  `json:"id"`
	Symbol     string `json:"symbol"`
	Market     string `json:"market"`
	Status     int    `json:"status"`   // 2关闭 1启用
	ShowFlag   int    `json:"showFlag"` // 2不展示 1展示
	Coin       string `json:"coin"`
	Sort       int    `json:"sort"`
	Logo       string `json:"logo"`
	BaseCoin   string `json:"baseCoin"`
	ShowSymbol string `json:"showSymbol"`
	Type       int    `json:"type"` // 1 外汇  2虚拟币
	CreateTime string `json:"createTime"`
}

type GetSecondCoinConfigListReq struct {
	g.Meta `path:"/secondContract/coin/list" tags:"AdminSecondContract" method:"get" summary:"获取秒合约币种配置列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"交易对"`
}

type GetSecondCoinConfigListRes struct {
	List  []SecondCoinConfigInfo `json:"list"`
	Total int                    `json:"total"`
}

// --------- 秒合约订单 (Second Contract Orders) ---------
type SecondContractOrderInfo struct {
	Id                 int64   `json:"id"`
	OrderNo            string  `json:"orderNo"`
	Symbol             string  `json:"symbol"`
	Type               string  `json:"type"`
	UserId             int     `json:"userId"`
	UserAddress        string  `json:"userAddress"`
	BetContent         string  `json:"betContent"` // 预测方向:0涨 1跌
	OpenResult         string  `json:"openResult"`
	Status             int     `json:"status"` // 0参与中 1已开奖 2已撤销
	RateFlag           int     `json:"rateFlag"`
	BetAmount          float64 `json:"betAmount"`
	RewardAmount       float64 `json:"rewardAmount"`
	CompensationAmount float64 `json:"compensationAmount"`
	OpenPrice          float64 `json:"openPrice"`
	ClosePrice         float64 `json:"closePrice"`
	CoinSymbol         string  `json:"coinSymbol"`
	BaseSymbol         string  `json:"baseSymbol"`
	Sign               int     `json:"sign"`               // 0正常 1包赢 2包输
	ManualIntervention int     `json:"manualIntervention"` // 0是 1否
	Rate               float64 `json:"rate"`
	CreateTime         string  `json:"createTime"`
}

type GetSecondContractOrderListReq struct {
	g.Meta  `path:"/secondContract/order/list" tags:"AdminSecondContract" method:"get" summary:"获取秒合约订单列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int    `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Symbol  string `json:"symbol" dc:"交易对"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetSecondContractOrderListRes struct {
	List  []SecondContractOrderInfo `json:"list"`
	Total int                       `json:"total"`
}
