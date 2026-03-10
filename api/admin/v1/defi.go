package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- DEFI订单 (DEFI Orders) ---------
type AdminDefiOrderInfo struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"userId"`
	Amount      float64 `json:"amount"`      // 收益金额
	TotleAmount float64 `json:"totleAmount"` // 钱包金额
	Rate        float64 `json:"rate"`        // 收益率
	CreateTime  string  `json:"createTime"`
}

type GetAdminDefiOrderListReq struct {
	g.Meta `path:"/defi/order/list" tags:"AdminDefi" method:"get" summary:"获取DEFI质押订单"`
	Page   int   `json:"page" d:"1"`
	Size   int   `json:"size" d:"20"`
	UserId int64 `json:"userId" dc:"用户ID"`
}

type GetAdminDefiOrderListRes struct {
	List  []AdminDefiOrderInfo `json:"list"`
	Total int                  `json:"total"`
}

// --------- 空投活动 (DEFI Airdrop Activity) ---------
type AdminDefiActivityInfo struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"userId"`
	TotleAmount float64 `json:"totleAmount"` // 达标所需金额
	Amount      float64 `json:"amount"`      // 奖励金额
	Type        int     `json:"type"`        // 0-usdt 1-eth
	Status      int     `json:"status"`      // 0未领取 1已读 2已领取
	CreateTime  string  `json:"createTime"`
	EndTime     string  `json:"endTime"`
}

type GetAdminDefiActivityListReq struct {
	g.Meta `path:"/defi/activity/list" tags:"AdminDefi" method:"get" summary:"获取空投活动发奖记录"`
	Page   int   `json:"page" d:"1"`
	Size   int   `json:"size" d:"20"`
	UserId int64 `json:"userId" dc:"用户ID"`
	Status *int  `json:"status" dc:"状态"`
}

type GetAdminDefiActivityListRes struct {
	List  []AdminDefiActivityInfo `json:"list"`
	Total int                     `json:"total"`
}

// --------- 挖矿利率配置 (DEFI Rate Config) ---------
type AdminDefiRateInfo struct {
	Id         int64   `json:"id"`
	Symbol     string  `json:"symbol"`     // 质押币种
	RewardCoin string  `json:"rewardCoin"` // 奖励币种
	MinAmount  float64 `json:"minAmount"`
	MaxAmount  float64 `json:"maxAmount"`
	Rate       float64 `json:"rate"`
	CreateTime string  `json:"createTime"`
}

type GetAdminDefiRateListReq struct {
	g.Meta `path:"/defi/rate/list" tags:"AdminDefi" method:"get" summary:"获取挖矿利率配置"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"20"`
}

type GetAdminDefiRateListRes struct {
	List  []AdminDefiRateInfo `json:"list"`
	Total int                 `json:"total"`
}
