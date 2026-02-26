// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta     `orm:"table:sys_dept, do:true"`
	DeptId     any         // 部门id
	ParentId   any         // 父部门id
	Ancestors  any         // 祖级列表
	DeptName   any         // 部门名称
	OrderNum   any         // 显示顺序
	Leader     any         // 负责人
	Phone      any         // 联系电话
	Email      any         // 邮箱
	Status     any         // 部门状态（0正常 1停用）
	DelFlag    any         // 删除标志（0代表存在 2代表删除）
	CreateBy   any         // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         // 更新者
	UpdateTime *gtime.Time // 更新时间
}
