// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzCronTriggersDao is the data access object for the table qrtz_cron_triggers.
type QrtzCronTriggersDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  QrtzCronTriggersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// QrtzCronTriggersColumns defines and stores column names for the table qrtz_cron_triggers.
type QrtzCronTriggersColumns struct {
	SchedName      string // 调度名称
	TriggerName    string // qrtz_triggers表trigger_name的外键
	TriggerGroup   string // qrtz_triggers表trigger_group的外键
	CronExpression string // cron表达式
	TimeZoneId     string // 时区
}

// qrtzCronTriggersColumns holds the columns for the table qrtz_cron_triggers.
var qrtzCronTriggersColumns = QrtzCronTriggersColumns{
	SchedName:      "sched_name",
	TriggerName:    "trigger_name",
	TriggerGroup:   "trigger_group",
	CronExpression: "cron_expression",
	TimeZoneId:     "time_zone_id",
}

// NewQrtzCronTriggersDao creates and returns a new DAO object for table data access.
func NewQrtzCronTriggersDao(handlers ...gdb.ModelHandler) *QrtzCronTriggersDao {
	return &QrtzCronTriggersDao{
		group:    "default",
		table:    "qrtz_cron_triggers",
		columns:  qrtzCronTriggersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzCronTriggersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzCronTriggersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzCronTriggersDao) Columns() QrtzCronTriggersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzCronTriggersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzCronTriggersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzCronTriggersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
