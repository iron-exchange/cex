// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiRate is the golang structure for table defi_rate.
type DefiRate struct {
	Id          int64       `json:"id"           orm:"id"           description:"id"`
	MinAmount   float64     `json:"min_amount"   orm:"min_amount"   description:"最小金额"`
	MaxAmount   float64     `json:"max_amount"   orm:"max_amount"   description:"最大金额"`
	Rate        float64     `json:"rate"         orm:"rate"         description:"利率"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建者"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
}
