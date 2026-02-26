// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HelpCenter is the golang structure of table t_help_center for DAO operations like Where/Data.
type HelpCenter struct {
	g.Meta     `orm:"table:t_help_center, do:true"`
	Id         any         //
	Title      any         // 标题
	Language   any         // 语言
	Enable     any         // 1=启用 2=禁用
	DelFlag    any         // 0=正常 1=删除
	CreateTime *gtime.Time //
	CreateBy   any         //
	UpdateTime *gtime.Time //
	UpdateBy   any         //
	Remark     any         // 备注
	ShowSymbol any         //
}
