// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppWalletRecord is the golang structure for table app_wallet_record.
type AppWalletRecord struct {
	Id             int64       `json:"id"               orm:"id"               description:"卡ID"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"余额"`
	UAmount        float64     `json:"u_amount"         orm:"u_amount"         description:"换算U金额"`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户id"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	BeforeAmount   float64     `json:"before_amount"    orm:"before_amount"    description:"前值"`
	AfterAmount    float64     `json:"after_amount"     orm:"after_amount"     description:"后值"`
	SerialId       string      `json:"serial_id"        orm:"serial_id"        description:""`
	Type           int         `json:"type"             orm:"type"             description:"余额"`
	Symbol         string      `json:"symbol"           orm:"symbol"           description:"币种"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"代理ID"`
	OperateTime    *gtime.Time `json:"operate_time"     orm:"operate_time"     description:"操作时间"`
}
