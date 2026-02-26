// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiOrder is the golang structure for table defi_order.
type DefiOrder struct {
	Id             int64       `json:"id"               orm:"id"               description:"id"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"收益金额"`
	TotleAmount    float64     `json:"totle_amount"     orm:"totle_amount"     description:"钱包金额"`
	Rate           float64     `json:"rate"             orm:"rate"             description:"收益率"`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:"创建者"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:"更新者"`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark         string      `json:"remark"           orm:"remark"           description:"备注"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户id"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"代理ids"`
}
