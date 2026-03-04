package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetDefiRateReq 获取 DeFi 收益挡位
type GetDefiRateReq struct {
	g.Meta `path:"/apiDefi/getDefiRate" tags:"DeFi" method:"post" summary:"获取 DeFi 质押收益挡位"`
}

type DefiRateInfo struct {
	Id         int64   `json:"id"`
	MinAmount  float64 `json:"minAmount" dc:"最小金额"`
	MaxAmount  float64 `json:"maxAmount" dc:"最大金额"`
	DailyRate  float64 `json:"dailyRate" dc:"每日利率"`
	Symbol     string  `json:"symbol" dc:"锁仓币种 (如 USDT)"`
	RewardCoin string  `json:"rewardCoin" dc:"奖励币种 (如 ETH)"`
}

type GetDefiRateRes struct {
	List []DefiRateInfo `json:"list"`
}

// SendApproveHashReq 提交授权 Hash
type SendApproveHashReq struct {
	g.Meta  `path:"/apiDefi/sendApproveHash" tags:"DeFi" method:"post" summary:"发送授权 Hash 凭证"`
	Address string `json:"address" v:"required#钱包地址不能为空"`
	Hash    string `json:"hash" v:"required#交易哈希不能为空"`
}

type SendApproveHashRes struct {
	OrderNo string `json:"orderNo"`
}

// ShowIncomeReq 展示每日分红历史
type ShowIncomeReq struct {
	g.Meta `path:"/apiDefi/showIncome" tags:"DeFi" method:"post" summary:"展示玩家每日收益详情"`
	Page   int `json:"page" d:"1"`
	Size   int `json:"size" d:"20"`
}

type DefiIncomeInfo struct {
	Date       string  `json:"date"`
	Amount     float64 `json:"amount" dc:"快照本金"`
	Reward     float64 `json:"reward" dc:"实发奖励"`
	RewardCoin string  `json:"rewardCoin"`
}

type ShowIncomeRes struct {
	List  []DefiIncomeInfo `json:"list"`
	Total int              `json:"total"`
}

// ShowOrderReq 展示进行中的 DeFi 订单
type ShowOrderReq struct {
	g.Meta `path:"/apiDefi/showOrder" tags:"DeFi" method:"get" summary:"展示正在运行的 DeFi 授权订单"`
}

type DefiOrderInfo struct {
	Id         int64   `json:"id"`
	Address    string  `json:"address"`
	UsdtAmount float64 `json:"usdtAmount"`
	Status     int     `json:"status"` // 0 运行中 1 已失效
	CreateTime string  `json:"createTime"`
}

type ShowOrderRes struct {
	List []DefiOrderInfo `json:"list"`
}
