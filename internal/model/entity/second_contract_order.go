// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondContractOrder is the golang structure for table second_contract_order.
type SecondContractOrder struct {
	Id                 int64       `json:"id"                  orm:"id"                  description:""`
	OrderNo            string      `json:"order_no"            orm:"order_no"            description:"订单号"`
	Symbol             string      `json:"symbol"              orm:"symbol"              description:"交易对"`
	Type               string      `json:"type"                orm:"type"                description:"类型"`
	UserId             int         `json:"user_id"             orm:"user_id"             description:"用户id"`
	UserAddress        string      `json:"user_address"        orm:"user_address"        description:"用户地址"`
	BetContent         string      `json:"bet_content"         orm:"bet_content"         description:"预测方向:0 涨 1跌"`
	OpenResult         string      `json:"open_result"         orm:"open_result"         description:"开奖结果"`
	Status             int         `json:"status"              orm:"status"              description:"订单状态 0参与中 1已开奖 2已撤销"`
	RateFlag           int         `json:"rate_flag"           orm:"rate_flag"           description:"是否全输"`
	BetAmount          float64     `json:"bet_amount"          orm:"bet_amount"          description:"投注金额"`
	RewardAmount       float64     `json:"reward_amount"       orm:"reward_amount"       description:"获奖金额"`
	CompensationAmount float64     `json:"compensation_amount" orm:"compensation_amount" description:"赔偿金额"`
	CreateTime         *gtime.Time `json:"create_time"         orm:"create_time"         description:"创建时间"`
	OpenPrice          float64     `json:"open_price"          orm:"open_price"          description:"开盘价格"`
	ClosePrice         float64     `json:"close_price"         orm:"close_price"         description:"关盘价格"`
	OpenTime           int64       `json:"open_time"           orm:"open_time"           description:"开盘时间"`
	CloseTime          int64       `json:"close_time"          orm:"close_time"          description:"关盘时间"`
	CoinSymbol         string      `json:"coin_symbol"         orm:"coin_symbol"         description:"交易币符号"`
	BaseSymbol         string      `json:"base_symbol"         orm:"base_symbol"         description:"结算币符号"`
	Sign               int         `json:"sign"                orm:"sign"                description:"订单标记 0正常  1包赢  2包输"`
	ManualIntervention int         `json:"manual_intervention" orm:"manual_intervention" description:"是否人工干预 0是 1否"`
	Rate               float64     `json:"rate"                orm:"rate"                description:""`
	CreateBy           string      `json:"create_by"           orm:"create_by"           description:""`
	UpdateTime         *gtime.Time `json:"update_time"         orm:"update_time"         description:""`
	UpdateBy           string      `json:"update_by"           orm:"update_by"           description:""`
	SearchValue        string      `json:"search_value"        orm:"search_value"        description:""`
	Remark             string      `json:"remark"              orm:"remark"              description:""`
	AdminParentIds     string      `json:"admin_parent_ids"    orm:"admin_parent_ids"    description:"后台代理ID"`
	IsHandling         int         `json:"is_handling"         orm:"is_handling"         description:"行锁"`
}
