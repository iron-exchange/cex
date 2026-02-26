// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzCalendarsDao is the data access object for the table qrtz_calendars.
type QrtzCalendarsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  QrtzCalendarsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// QrtzCalendarsColumns defines and stores column names for the table qrtz_calendars.
type QrtzCalendarsColumns struct {
	SchedName    string // 调度名称
	CalendarName string // 日历名称
	Calendar     string // 存放持久化calendar对象
}

// qrtzCalendarsColumns holds the columns for the table qrtz_calendars.
var qrtzCalendarsColumns = QrtzCalendarsColumns{
	SchedName:    "sched_name",
	CalendarName: "calendar_name",
	Calendar:     "calendar",
}

// NewQrtzCalendarsDao creates and returns a new DAO object for table data access.
func NewQrtzCalendarsDao(handlers ...gdb.ModelHandler) *QrtzCalendarsDao {
	return &QrtzCalendarsDao{
		group:    "default",
		table:    "qrtz_calendars",
		columns:  qrtzCalendarsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzCalendarsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzCalendarsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzCalendarsDao) Columns() QrtzCalendarsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzCalendarsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzCalendarsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzCalendarsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
