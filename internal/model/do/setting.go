// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Setting is the golang structure of table t_setting for DAO operations like Where/Data.
type Setting struct {
	g.Meta       `orm:"table:t_setting, do:true"`
	Id           any         // ID
	CreateBy     any         // 创建者
	CreateTime   *gtime.Time // 创建时间
	DeleteFlag   any         // 删除标志 true/false 删除/未删除
	UpdateBy     any         // 更新者
	UpdateTime   *gtime.Time // 更新时间
	SettingValue any         // 配置值value
}
