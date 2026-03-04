package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MiningShowReq 挖矿展示
type MiningShowReq struct {
	g.Meta `path:"/mingOrder/show" tags:"Mining" method:"post" summary:"展示可用矿机与进行中订单"`
}

type MiningProductInfo struct {
	Id       int64   `json:"id"`
	Title    string  `json:"title"`
	MinPrice float64 `json:"minPrice"`
	MaxPrice float64 `json:"maxPrice"`
	MinOdds  float64 `json:"minOdds"`
	MaxOdds  float64 `json:"maxOdds"`
}

type MiningOrderInfo struct {
	OrderNo    string  `json:"orderNo"`
	Amount     float64 `json:"amount"`
	Days       int     `json:"days"`
	Status     int     `json:"status"` // 0 进行中 1 已结束
	CreateTime string  `json:"createTime"`
}

type MiningShowRes struct {
	Products []MiningProductInfo `json:"products"`
	Orders   []MiningOrderInfo   `json:"orders"`
}

// MiningSubmitReq 挖矿购买
type MiningSubmitReq struct {
	g.Meta `path:"/mingOrder/submit" tags:"Mining" method:"post" summary:"申购矿机产品"`
	PlanId int64   `json:"planId" v:"required#矿机ID不能为空"`
	Amount float64 `json:"amount" v:"required|min:1#申购金额无效"`
}

type MiningSubmitRes struct {
	OrderNo string `json:"orderNo"`
}

// MiningRedemptionReq 提前赎回
type MiningRedemptionReq struct {
	g.Meta  `path:"/mingOrder/redemption" tags:"Mining" method:"post" summary:"申请提前赎回矿机本金"`
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
}

type MiningRedemptionRes struct {
	Success bool `json:"success"`
}
