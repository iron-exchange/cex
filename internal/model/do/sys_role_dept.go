// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDept is the golang structure of table sys_role_dept for DAO operations like Where/Data.
type SysRoleDept struct {
	g.Meta `orm:"table:sys_role_dept, do:true"`
	RoleId any // 角色ID
	DeptId any // 部门ID
}
