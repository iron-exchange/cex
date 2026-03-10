package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种列表 (IEO Coin List) ---------
type OwnCoinInfo struct {
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
	Status        int     `json:"status"` // 1.未发布  2.筹集中 3 筹集成功 4.筹集失败
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
}

type GetOwnCoinListReq struct {
	g.Meta `path:"/ieo/coin/list" tags:"AdminIEO" method:"get" summary:"获取IEO新币列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Coin   string `json:"coin" dc:"币种名称过滤"`
}

type GetOwnCoinListRes struct {
	List  []OwnCoinInfo `json:"list"`
	Total int           `json:"total"`
}

// --------- 申购订单 (IEO Subscription Orders) ---------
type OwnCoinSubscribeOrderInfo struct {
	Id          int64   `json:"id"`
	SubscribeId string  `json:"subscribeId"`
	UserId      int64   `json:"userId"`
	OrderId     string  `json:"orderId"`
	OwnCoin     string  `json:"ownCoin"`
	AmountLimit float64 `json:"amountLimit"`
	NumLimit    int     `json:"numLimit"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"` // 1订阅中、2订阅成功、3成功消息推送完成
	CreateTime  string  `json:"createTime"`
}

type GetOwnCoinSubscribeOrderListReq struct {
	g.Meta  `path:"/ieo/order/list" tags:"AdminIEO" method:"get" summary:"获取新币申购订单记录"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OwnCoin string `json:"ownCoin" dc:"申购币种"`
	Status  string `json:"status" dc:"状态"`
}

type GetOwnCoinSubscribeOrderListRes struct {
	List  []OwnCoinSubscribeOrderInfo `json:"list"`
	Total int                         `json:"total"`
}
