// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzSimpleTriggersDao is the data access object for the table qrtz_simple_triggers.
type QrtzSimpleTriggersDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  QrtzSimpleTriggersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// QrtzSimpleTriggersColumns defines and stores column names for the table qrtz_simple_triggers.
type QrtzSimpleTriggersColumns struct {
	SchedName      string // 调度名称
	TriggerName    string // qrtz_triggers表trigger_name的外键
	TriggerGroup   string // qrtz_triggers表trigger_group的外键
	RepeatCount    string // 重复的次数统计
	RepeatInterval string // 重复的间隔时间
	TimesTriggered string // 已经触发的次数
}

// qrtzSimpleTriggersColumns holds the columns for the table qrtz_simple_triggers.
var qrtzSimpleTriggersColumns = QrtzSimpleTriggersColumns{
	SchedName:      "sched_name",
	TriggerName:    "trigger_name",
	TriggerGroup:   "trigger_group",
	RepeatCount:    "repeat_count",
	RepeatInterval: "repeat_interval",
	TimesTriggered: "times_triggered",
}

// NewQrtzSimpleTriggersDao creates and returns a new DAO object for table data access.
func NewQrtzSimpleTriggersDao(handlers ...gdb.ModelHandler) *QrtzSimpleTriggersDao {
	return &QrtzSimpleTriggersDao{
		group:    "default",
		table:    "qrtz_simple_triggers",
		columns:  qrtzSimpleTriggersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzSimpleTriggersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzSimpleTriggersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzSimpleTriggersDao) Columns() QrtzSimpleTriggersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzSimpleTriggersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzSimpleTriggersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzSimpleTriggersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
