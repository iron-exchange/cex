package v1

import "github.com/gogf/gf/v2/frame/g"

// OwnCoinAppInfo 用于客户端展示的自发币模型
type OwnCoinAppInfo struct {
	Id            int64   `json:"id"`
	Coin          string  `json:"coin"`
	Logo          string  `json:"logo"`
	ReferCoin     string  `json:"referCoin"`
	ShowSymbol    string  `json:"showSymbol"`
	Price         float64 `json:"price"`
	Proportion    float64 `json:"proportion"`
	RaisingAmount float64 `json:"raisingAmount"`
	RaisedAmount  float64 `json:"raisedAmount"`
	PurchaseLimit int     `json:"purchaseLimit"`
	TotalAmount   float64 `json:"totalAmount"`
	Status        int     `json:"status"` // 1.未发布 2.筹集中 3 筹集成功 4.筹集失败
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
	Introduce     string  `json:"introduce"`
}

// GetOwnCoinListReq 获取自发币列表
type GetOwnCoinListReq struct {
	g.Meta `path:"/api/ownCoin/list" tags:"AppOwnCoin" method:"post" summary:"查询自发币(新币)列表"`
	Status string `json:"status" dc:"按状态过滤，例如2"`
	Page   int    `json:"pageNum" d:"1"`
	Size   int    `json:"pageSize" d:"20"`
}

type GetOwnCoinListRes struct {
	Total int              `json:"total"`
	Rows  []OwnCoinAppInfo `json:"rows"`
}

// GetOwnCoinDetailReq 获取自发币详情
type GetOwnCoinDetailReq struct {
	g.Meta `path:"/api/ownCoin/getDetail/{ownId}" tags:"AppOwnCoin" method:"get" summary:"查询自发币详情"`
	OwnId  int64 `json:"ownId" in:"path" v:"required#自发币ID不能为空"`
}

type GetOwnCoinDetailRes struct {
	Data struct {
		CoinInfo      OwnCoinAppInfo `json:"coinInfo"`
		PurchasedAmt  float64        `json:"purchasedAmt" dc:"当前用户已申购数量"`
		RemainingAmt  float64        `json:"remainingAmt" dc:"全网剩余可申购量"`
		UserRemaining int            `json:"userRemaining" dc:"用户剩余可申购次数/额度（按需求定）"`
	} `json:"data"`
}

// SubscribeOwnCoinReq 认购新币
type SubscribeOwnCoinReq struct {
	g.Meta        `path:"/api/ownCoin/subscribeCoins" tags:"AppOwnCoin" method:"post" summary:"认购新币"`
	OwnId         int64   `json:"ownId" v:"required#自发币ID不能为空"`
	PayCoin       string  `json:"payCoin" v:"required#支付币种不能为空"` // Usually the ReferCoin
	PayAmount     float64 `json:"payAmount" v:"required#支付数量不能为空"`
	ReceiveAmount float64 `json:"receiveAmount"` // Expected receive amount, for validation
}

type SubscribeOwnCoinRes struct{}
