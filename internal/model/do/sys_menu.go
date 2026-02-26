// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	MenuId     any         // 菜单ID
	MenuName   any         // 菜单名称
	ParentId   any         // 父菜单ID
	OrderNum   any         // 显示顺序
	Path       any         // 路由地址
	Component  any         // 组件路径
	Query      any         // 路由参数
	IsFrame    any         // 是否为外链（0是 1否）
	IsCache    any         // 是否缓存（0缓存 1不缓存）
	MenuType   any         // 菜单类型（M目录 C菜单 F按钮）
	Visible    any         // 菜单状态（0显示 1隐藏）
	Status     any         // 菜单状态（0正常 1停用）
	Perms      any         // 权限标识
	Icon       any         // 菜单图标
	CreateBy   any         // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     any         // 备注
}
