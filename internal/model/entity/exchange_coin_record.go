// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ExchangeCoinRecord is the golang structure for table exchange_coin_record.
type ExchangeCoinRecord struct {
	Id             int64       `json:"id"               orm:"id"               description:""`
	FromCoin       string      `json:"from_coin"        orm:"from_coin"        description:""`
	ToCoin         string      `json:"to_coin"          orm:"to_coin"          description:""`
	UserId         int         `json:"user_id"          orm:"user_id"          description:"用户id"`
	Username       string      `json:"username"         orm:"username"         description:"用户名称"`
	Address        string      `json:"address"          orm:"address"          description:"用户地址"`
	Status         int         `json:"status"           orm:"status"           description:"兑换状态0:已提交;1:成功;2失败"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"金额"`
	ThirdRate      float64     `json:"third_rate"       orm:"third_rate"       description:"三方汇率"`
	SystemRate     float64     `json:"system_rate"      orm:"system_rate"      description:"系统汇率"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:""`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:""`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:""`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
}
