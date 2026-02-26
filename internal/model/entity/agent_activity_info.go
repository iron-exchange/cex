// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentActivityInfo is the golang structure for table agent_activity_info.
type AgentActivityInfo struct {
	Id         float64     `json:"id"          orm:"id"          description:"id"`
	Type       int         `json:"type"        orm:"type"        description:"1 充值返利 2挖矿返利"`
	Amount     float64     `json:"amount"      orm:"amount"      description:"返利金额"`
	CoinType   string      `json:"coin_type"   orm:"coin_type"   description:"币种"`
	FromId     int64       `json:"from_id"     orm:"from_id"     description:"返利用户"`
	UserId     int64       `json:"user_id"     orm:"user_id"     description:"用户id"`
	CreateBy   string      `json:"create_by"   orm:"create_by"   description:""`
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"update_by"   orm:"update_by"   description:""`
	UpdateTime *gtime.Time `json:"update_time" orm:"update_time" description:"更新时间"`
	Status     int         `json:"status"      orm:"status"      description:"1  待返  2  已返"`
	LoginName  string      `json:"login_name"  orm:"login_name"  description:""`
	SerialId   string      `json:"serial_id"   orm:"serial_id"   description:""`
}
