// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta            `orm:"table:sys_role, do:true"`
	RoleId            any         // 角色ID
	RoleName          any         // 角色名称
	RoleKey           any         // 角色权限字符串
	RoleSort          any         // 显示顺序
	DataScope         any         // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly any         // 菜单树选择项是否关联显示
	DeptCheckStrictly any         // 部门树选择项是否关联显示
	Status            any         // 角色状态（0正常 1停用）
	DelFlag           any         // 删除标志（0代表存在 2代表删除）
	CreateBy          any         // 创建者
	CreateTime        *gtime.Time // 创建时间
	UpdateBy          any         // 更新者
	UpdateTime        *gtime.Time // 更新时间
	Remark            any         // 备注
}
