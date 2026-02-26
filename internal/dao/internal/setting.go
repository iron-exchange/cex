// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingDao is the data access object for the table t_setting.
type SettingDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SettingColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SettingColumns defines and stores column names for the table t_setting.
type SettingColumns struct {
	Id           string // ID
	CreateBy     string // 创建者
	CreateTime   string // 创建时间
	DeleteFlag   string // 删除标志 true/false 删除/未删除
	UpdateBy     string // 更新者
	UpdateTime   string // 更新时间
	SettingValue string // 配置值value
}

// settingColumns holds the columns for the table t_setting.
var settingColumns = SettingColumns{
	Id:           "id",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	DeleteFlag:   "delete_flag",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
	SettingValue: "setting_value",
}

// NewSettingDao creates and returns a new DAO object for table data access.
func NewSettingDao(handlers ...gdb.ModelHandler) *SettingDao {
	return &SettingDao{
		group:    "default",
		table:    "t_setting",
		columns:  settingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingDao) Columns() SettingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
