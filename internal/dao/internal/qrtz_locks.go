// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzLocksDao is the data access object for the table qrtz_locks.
type QrtzLocksDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  QrtzLocksColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// QrtzLocksColumns defines and stores column names for the table qrtz_locks.
type QrtzLocksColumns struct {
	SchedName string // 调度名称
	LockName  string // 悲观锁名称
}

// qrtzLocksColumns holds the columns for the table qrtz_locks.
var qrtzLocksColumns = QrtzLocksColumns{
	SchedName: "sched_name",
	LockName:  "lock_name",
}

// NewQrtzLocksDao creates and returns a new DAO object for table data access.
func NewQrtzLocksDao(handlers ...gdb.ModelHandler) *QrtzLocksDao {
	return &QrtzLocksDao{
		group:    "default",
		table:    "qrtz_locks",
		columns:  qrtzLocksColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzLocksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzLocksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzLocksDao) Columns() QrtzLocksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzLocksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzLocksDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzLocksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
