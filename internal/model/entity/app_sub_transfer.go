// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppSubTransfer is the golang structure for table app_sub_transfer.
type AppSubTransfer struct {
	Id         int64       `json:"id"          orm:"id"          description:""`
	UserId     int64       `json:"user_id"     orm:"user_id"     description:"用户id"`
	Email      string      `json:"email"       orm:"email"       description:"子账号邮箱"`
	EmailType  int         `json:"email_type"  orm:"email_type"  description:"1转入，2转出"`
	Type       string      `json:"type"        orm:"type"        description:"划转类型：\"SPOT\",\"USDT_FUTURE\",\"COIN_FUTURE\",\"MARGIN\"(Cross),\"ISOLATED_MARGIN\""`
	Tranid     string      `json:"tranid"      orm:"tranid"      description:"订单号"`
	Amount     float64     `json:"amount"      orm:"amount"      description:"价格"`
	Asset      string      `json:"asset"       orm:"asset"       description:"货币类型"`
	Status     int         `json:"status"      orm:"status"      description:"1正常 2失败"`
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
}
