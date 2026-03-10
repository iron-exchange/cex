package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 充值列表 (Recharge List) ---------
type RechargeInfo struct {
	Id         int64   `json:"id"`
	UserId     int64   `json:"userId"`
	LoginName  string  `json:"loginName"`
	OrderNo    string  `json:"orderNo"`
	Coin       string  `json:"coin"`
	Amount     float64 `json:"amount"`
	Address    string  `json:"address"`
	TxHash     string  `json:"txHash"`
	Status     int     `json:"status"` // 1/null充值中，2充值成功，3充值失败
	CreateTime string  `json:"createTime"`
	Remark     string  `json:"remark"`
}

type GetRechargeListReq struct {
	g.Meta  `path:"/funding/recharge/list" tags:"AdminFunding" method:"get" summary:"获取充值列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetRechargeListRes struct {
	List  []RechargeInfo `json:"list"`
	Total int            `json:"total"`
}

// --------- 提现列表 (Withdraw List) ---------
type WithdrawInfo struct {
	Id         int     `json:"id"`
	UserId     int     `json:"userId"`
	LoginName  string  `json:"loginName"`
	OrderNo    string  `json:"orderNo"`
	Coin       string  `json:"coin"`
	Amount     float64 `json:"amount"`
	Address    string  `json:"address"`
	Status     int     `json:"status"` // 1申请，2通过，3拒绝
	CreateTime string  `json:"createTime"`
	Remark     string  `json:"remark"`
}

type GetWithdrawListReq struct {
	g.Meta  `path:"/funding/withdraw/list" tags:"AdminFunding" method:"get" summary:"获取提现列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetWithdrawListRes struct {
	List  []WithdrawInfo `json:"list"`
	Total int            `json:"total"`
}

// --------- 充值通道配置 (Recharge Channel Config) ---------
type RechargeChannelInfo struct {
	Id         int64  `json:"id"`
	Symbol     string `json:"symbol"`
	Enable     string `json:"enable"` // 0关闭 1开启
	Type       string `json:"type"`   // 充值通道类型
	CreateTime string `json:"createTime"`
}

type GetRechargeChannelListReq struct {
	g.Meta `path:"/funding/channel/recharge/list" tags:"AdminFunding" method:"get" summary:"获取充值通道配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"币种"`
}

type GetRechargeChannelListRes struct {
	List  []RechargeChannelInfo `json:"list"`
	Total int                   `json:"total"`
}

// --------- 提现通道配置 (Withdraw Channel Config) ---------
type WithdrawChannelInfo struct {
	Id         int64  `json:"id"`
	Coin       string `json:"coin"`
	Enable     string `json:"enable"`  // 0关闭 1开启
	FeeRate    string `json:"feeRate"` // 手续费
	MinLimit   string `json:"minLimit"`
	MaxLimit   string `json:"maxLimit"`
	CreateTime string `json:"createTime"`
}

type GetWithdrawChannelListReq struct {
	g.Meta `path:"/funding/channel/withdraw/list" tags:"AdminFunding" method:"get" summary:"获取提现通道配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Coin   string `json:"coin" dc:"币种"`
}

type GetWithdrawChannelListRes struct {
	List  []WithdrawChannelInfo `json:"list"`
	Total int                   `json:"total"`
}
