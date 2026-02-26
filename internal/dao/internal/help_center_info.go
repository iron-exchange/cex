// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HelpCenterInfoDao is the data access object for the table t_help_center_info.
type HelpCenterInfoDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  HelpCenterInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// HelpCenterInfoColumns defines and stores column names for the table t_help_center_info.
type HelpCenterInfoColumns struct {
	Id           string //
	HelpCenterId string // 帮助中心主键id
	Question     string // 标题
	Content      string // 内容
	Language     string // 语言
	Enable       string // 1=启用 2=禁用
	DelFlag      string // 0=正常 1=删除
	CreateTime   string //
	CreateBy     string //
	UpdateTime   string //
	UpdateBy     string //
	Remark       string // 备注
	ShowSymbol   string //
}

// helpCenterInfoColumns holds the columns for the table t_help_center_info.
var helpCenterInfoColumns = HelpCenterInfoColumns{
	Id:           "id",
	HelpCenterId: "help_center_id",
	Question:     "question",
	Content:      "content",
	Language:     "language",
	Enable:       "enable",
	DelFlag:      "del_flag",
	CreateTime:   "create_time",
	CreateBy:     "create_by",
	UpdateTime:   "update_time",
	UpdateBy:     "update_by",
	Remark:       "remark",
	ShowSymbol:   "show_symbol",
}

// NewHelpCenterInfoDao creates and returns a new DAO object for table data access.
func NewHelpCenterInfoDao(handlers ...gdb.ModelHandler) *HelpCenterInfoDao {
	return &HelpCenterInfoDao{
		group:    "default",
		table:    "t_help_center_info",
		columns:  helpCenterInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HelpCenterInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HelpCenterInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HelpCenterInfoDao) Columns() HelpCenterInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HelpCenterInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HelpCenterInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HelpCenterInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
