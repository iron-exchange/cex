// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiActivity is the golang structure for table defi_activity.
type DefiActivity struct {
	Id          int64       `json:"id"           orm:"id"           description:"id"`
	TotleAmount float64     `json:"totle_amount" orm:"totle_amount" description:"需要金额"`
	UserId      int64       `json:"user_id"      orm:"user_id"      description:"用户id"`
	EndTime     *gtime.Time `json:"end_time"     orm:"end_time"     description:"结束时间"`
	Amount      float64     `json:"amount"       orm:"amount"       description:"奖励金额"`
	Type        int         `json:"type"         orm:"type"         description:"0-usdt 1-eth"`
	Status      int         `json:"status"       orm:"status"       description:"0未领取 1已读 2已领取"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建者"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
}
