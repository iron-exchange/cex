// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityRecharge is the golang structure for table activity_recharge.
type ActivityRecharge struct {
	Id          int         `json:"id"           orm:"id"           description:"id"`
	OnOff       int         `json:"on_off"       orm:"on_off"       description:"0-关闭 1-开启"`
	RechargePro float64     `json:"recharge_pro" orm:"recharge_pro" description:"充值返点比例"`
	MaxRebate   float64     `json:"max_rebate"   orm:"max_rebate"   description:"充值返点最大值"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:""`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:""`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
}
