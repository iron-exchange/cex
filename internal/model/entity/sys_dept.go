// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId     int64       `json:"dept_id"     orm:"dept_id"     description:"部门id"`
	ParentId   int64       `json:"parent_id"   orm:"parent_id"   description:"父部门id"`
	Ancestors  string      `json:"ancestors"   orm:"ancestors"   description:"祖级列表"`
	DeptName   string      `json:"dept_name"   orm:"dept_name"   description:"部门名称"`
	OrderNum   int         `json:"order_num"   orm:"order_num"   description:"显示顺序"`
	Leader     string      `json:"leader"      orm:"leader"      description:"负责人"`
	Phone      string      `json:"phone"       orm:"phone"       description:"联系电话"`
	Email      string      `json:"email"       orm:"email"       description:"邮箱"`
	Status     string      `json:"status"      orm:"status"      description:"部门状态（0正常 1停用）"`
	DelFlag    string      `json:"del_flag"    orm:"del_flag"    description:"删除标志（0代表存在 2代表删除）"`
	CreateBy   string      `json:"create_by"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"update_by"   orm:"update_by"   description:"更新者"`
	UpdateTime *gtime.Time `json:"update_time" orm:"update_time" description:"更新时间"`
}
