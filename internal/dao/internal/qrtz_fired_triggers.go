// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzFiredTriggersDao is the data access object for the table qrtz_fired_triggers.
type QrtzFiredTriggersDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  QrtzFiredTriggersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// QrtzFiredTriggersColumns defines and stores column names for the table qrtz_fired_triggers.
type QrtzFiredTriggersColumns struct {
	SchedName        string // 调度名称
	EntryId          string // 调度器实例id
	TriggerName      string // qrtz_triggers表trigger_name的外键
	TriggerGroup     string // qrtz_triggers表trigger_group的外键
	InstanceName     string // 调度器实例名
	FiredTime        string // 触发的时间
	SchedTime        string // 定时器制定的时间
	Priority         string // 优先级
	State            string // 状态
	JobName          string // 任务名称
	JobGroup         string // 任务组名
	IsNonconcurrent  string // 是否并发
	RequestsRecovery string // 是否接受恢复执行
}

// qrtzFiredTriggersColumns holds the columns for the table qrtz_fired_triggers.
var qrtzFiredTriggersColumns = QrtzFiredTriggersColumns{
	SchedName:        "sched_name",
	EntryId:          "entry_id",
	TriggerName:      "trigger_name",
	TriggerGroup:     "trigger_group",
	InstanceName:     "instance_name",
	FiredTime:        "fired_time",
	SchedTime:        "sched_time",
	Priority:         "priority",
	State:            "state",
	JobName:          "job_name",
	JobGroup:         "job_group",
	IsNonconcurrent:  "is_nonconcurrent",
	RequestsRecovery: "requests_recovery",
}

// NewQrtzFiredTriggersDao creates and returns a new DAO object for table data access.
func NewQrtzFiredTriggersDao(handlers ...gdb.ModelHandler) *QrtzFiredTriggersDao {
	return &QrtzFiredTriggersDao{
		group:    "default",
		table:    "qrtz_fired_triggers",
		columns:  qrtzFiredTriggersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzFiredTriggersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzFiredTriggersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzFiredTriggersDao) Columns() QrtzFiredTriggersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzFiredTriggersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzFiredTriggersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzFiredTriggersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
