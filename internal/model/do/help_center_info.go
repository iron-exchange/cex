// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HelpCenterInfo is the golang structure of table t_help_center_info for DAO operations like Where/Data.
type HelpCenterInfo struct {
	g.Meta       `orm:"table:t_help_center_info, do:true"`
	Id           any         //
	HelpCenterId any         // 帮助中心主键id
	Question     any         // 标题
	Content      any         // 内容
	Language     any         // 语言
	Enable       any         // 1=启用 2=禁用
	DelFlag      any         // 0=正常 1=删除
	CreateTime   *gtime.Time //
	CreateBy     any         //
	UpdateTime   *gtime.Time //
	UpdateBy     any         //
	Remark       any         // 备注
	ShowSymbol   any         //
}
