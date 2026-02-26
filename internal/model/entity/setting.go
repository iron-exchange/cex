// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Setting is the golang structure for table setting.
type Setting struct {
	Id           string      `json:"id"            orm:"id"            description:"ID"`
	CreateBy     string      `json:"create_by"     orm:"create_by"     description:"创建者"`
	CreateTime   *gtime.Time `json:"create_time"   orm:"create_time"   description:"创建时间"`
	DeleteFlag   string      `json:"delete_flag"   orm:"delete_flag"   description:"删除标志 true/false 删除/未删除"`
	UpdateBy     string      `json:"update_by"     orm:"update_by"     description:"更新者"`
	UpdateTime   *gtime.Time `json:"update_time"   orm:"update_time"   description:"更新时间"`
	SettingValue string      `json:"setting_value" orm:"setting_value" description:"配置值value"`
}
