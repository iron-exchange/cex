// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPostDao is the data access object for the table sys_post.
type SysPostDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysPostColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysPostColumns defines and stores column names for the table sys_post.
type SysPostColumns struct {
	PostId     string // 岗位ID
	PostCode   string // 岗位编码
	PostName   string // 岗位名称
	PostSort   string // 显示顺序
	Status     string // 状态（0正常 1停用）
	CreateBy   string // 创建者
	CreateTime string // 创建时间
	UpdateBy   string // 更新者
	UpdateTime string // 更新时间
	Remark     string // 备注
}

// sysPostColumns holds the columns for the table sys_post.
var sysPostColumns = SysPostColumns{
	PostId:     "post_id",
	PostCode:   "post_code",
	PostName:   "post_name",
	PostSort:   "post_sort",
	Status:     "status",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewSysPostDao creates and returns a new DAO object for table data access.
func NewSysPostDao(handlers ...gdb.ModelHandler) *SysPostDao {
	return &SysPostDao{
		group:    "default",
		table:    "sys_post",
		columns:  sysPostColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPostDao) Columns() SysPostColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPostDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
