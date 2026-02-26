// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MingProductUser is the golang structure for table ming_product_user.
type MingProductUser struct {
	Id          int         `json:"id"           orm:"id"           description:""`
	ProductId   int         `json:"product_id"   orm:"product_id"   description:"产品id"`
	AppUserId   int64       `json:"app_user_id"  orm:"app_user_id"  description:"玩家用户id"`
	PledgeNum   int         `json:"pledge_num"   orm:"pledge_num"   description:"限购次数"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建人"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新人员"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	SearchValue string      `json:"search_value" orm:"search_value" description:"币种"`
	Remark      string      `json:"remark"       orm:"remark"       description:"标签"`
}
