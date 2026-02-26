// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HelpCenterInfo is the golang structure for table help_center_info.
type HelpCenterInfo struct {
	Id           float64     `json:"id"             orm:"id"             description:""`
	HelpCenterId int64       `json:"help_center_id" orm:"help_center_id" description:"帮助中心主键id"`
	Question     string      `json:"question"       orm:"question"       description:"标题"`
	Content      string      `json:"content"        orm:"content"        description:"内容"`
	Language     string      `json:"language"       orm:"language"       description:"语言"`
	Enable       string      `json:"enable"         orm:"enable"         description:"1=启用 2=禁用"`
	DelFlag      string      `json:"del_flag"       orm:"del_flag"       description:"0=正常 1=删除"`
	CreateTime   *gtime.Time `json:"create_time"    orm:"create_time"    description:""`
	CreateBy     string      `json:"create_by"      orm:"create_by"      description:""`
	UpdateTime   *gtime.Time `json:"update_time"    orm:"update_time"    description:""`
	UpdateBy     string      `json:"update_by"      orm:"update_by"      description:""`
	Remark       string      `json:"remark"         orm:"remark"         description:"备注"`
	ShowSymbol   string      `json:"show_symbol"    orm:"show_symbol"    description:""`
}
