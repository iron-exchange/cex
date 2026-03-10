package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========= 秒合约币种配置 =========

type SecondCoinListReq struct {
	g.Meta `path:"/coin/list" tags:"SecondContract" method:"post" summary:"查询秒合约币种配置列表"`
}

type SecondPeriodInfo struct {
	Id        int64   `json:"id"`
	SecondId  int64   `json:"secondId"`
	Period    int     `json:"period"`
	Odds      float64 `json:"odds"`
	MaxAmount float64 `json:"maxAmount"`
	MinAmount float64 `json:"minAmount"`
}

type SecondCoinInfo struct {
	Id         int64              `json:"id"`
	Symbol     string             `json:"symbol"`
	Coin       string             `json:"coin"`
	BaseCoin   string             `json:"baseCoin"`
	Logo       string             `json:"logo"`
	PeriodList []SecondPeriodInfo `json:"periodList"`
}

type SecondCoinListRes struct {
	Rows []SecondCoinInfo `json:"rows"`
}

type SecondCoinDetailReq struct {
	g.Meta `path:"/coin/{id}" tags:"SecondContract" method:"post" summary:"获取单个秒合约币种配置"`
	Id     int64 `json:"id" in:"path" v:"required#ID不能为空"`
}

type SecondCoinDetailRes struct {
	Data SecondCoinInfo `json:"data"`
}

// ========= 秒合约订单 =========

type CreateSecondOrderReq struct {
	g.Meta     `path:"/secondContractOrder/createSecondContractOrder" tags:"SecondContract" method:"post" summary:"新增秒合约订单"`
	Symbol     string  `json:"symbol" v:"required#交易对不能为空"`
	BetContent string  `json:"betContent" v:"required#买涨买跌不能为空"` // 0 涨 1 跌
	BetAmount  float64 `json:"betAmount" v:"required|min:1#投资金额无效"`
	PeriodId   int64   `json:"periodId" v:"required#投资周期不能为空"` // 对应 SecondPeriodConfig 的 ID
}

type CreateSecondOrderRes struct {
	OrderNo string `json:"orderNo"`
}

type SelectSecondOrderListReq struct {
	g.Meta `path:"/secondContractOrder/selectOrderList" tags:"SecondContract" method:"post" summary:"查询秒合约订单列表"`
	Status int `json:"status"` // 0参与中 1已开奖 2已撤销 (如果传空或特定值查全部, 需在logic判断)
	Page   int `json:"pageNum" d:"1"`
	Size   int `json:"pageSize" d:"20"`
}

type SecondOrderItem struct {
	Id         int64   `json:"id"`
	OrderNo    string  `json:"orderNo"`
	Symbol     string  `json:"symbol"`
	BetContent string  `json:"betContent"`
	BetAmount  float64 `json:"betAmount"`
	Status     int     `json:"status"`
	OpenPrice  float64 `json:"openPrice"`
	ClosePrice float64 `json:"closePrice"`
	RewardAmt  float64 `json:"rewardAmount"`
	CreateTime string  `json:"createTime"`
	OpenTime   int64   `json:"openTime"`
	CloseTime  int64   `json:"closeTime"`
}

type SelectSecondOrderListRes struct {
	Rows  []SecondOrderItem `json:"rows"`
	Total int               `json:"total"`
}
