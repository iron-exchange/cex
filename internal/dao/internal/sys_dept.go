// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeptDao is the data access object for the table sys_dept.
type SysDeptDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysDeptColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysDeptColumns defines and stores column names for the table sys_dept.
type SysDeptColumns struct {
	DeptId     string // 部门id
	ParentId   string // 父部门id
	Ancestors  string // 祖级列表
	DeptName   string // 部门名称
	OrderNum   string // 显示顺序
	Leader     string // 负责人
	Phone      string // 联系电话
	Email      string // 邮箱
	Status     string // 部门状态（0正常 1停用）
	DelFlag    string // 删除标志（0代表存在 2代表删除）
	CreateBy   string // 创建者
	CreateTime string // 创建时间
	UpdateBy   string // 更新者
	UpdateTime string // 更新时间
}

// sysDeptColumns holds the columns for the table sys_dept.
var sysDeptColumns = SysDeptColumns{
	DeptId:     "dept_id",
	ParentId:   "parent_id",
	Ancestors:  "ancestors",
	DeptName:   "dept_name",
	OrderNum:   "order_num",
	Leader:     "leader",
	Phone:      "phone",
	Email:      "email",
	Status:     "status",
	DelFlag:    "del_flag",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
}

// NewSysDeptDao creates and returns a new DAO object for table data access.
func NewSysDeptDao(handlers ...gdb.ModelHandler) *SysDeptDao {
	return &SysDeptDao{
		group:    "default",
		table:    "sys_dept",
		columns:  sysDeptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysDeptDao) Columns() SysDeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysDeptDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
