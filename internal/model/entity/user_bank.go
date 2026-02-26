// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBank is the golang structure for table user_bank.
type UserBank struct {
	Id             int64       `json:"id"               orm:"id"               description:""`
	UserName       string      `json:"user_name"        orm:"user_name"        description:"姓名"`
	CardNumber     string      `json:"card_number"      orm:"card_number"      description:"银行卡号"`
	BankName       string      `json:"bank_name"        orm:"bank_name"        description:"开户银行名称"`
	BankAddress    string      `json:"bank_address"     orm:"bank_address"     description:"开户省市"`
	BankBranch     string      `json:"bank_branch"      orm:"bank_branch"      description:"开户网点"`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户名称"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:""`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	BankCode       string      `json:"bank_code"        orm:"bank_code"        description:"银行编码"`
	UserAddress    string      `json:"user_address"     orm:"user_address"     description:"用户地址"`
}
