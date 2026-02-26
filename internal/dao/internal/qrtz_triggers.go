// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzTriggersDao is the data access object for the table qrtz_triggers.
type QrtzTriggersDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  QrtzTriggersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// QrtzTriggersColumns defines and stores column names for the table qrtz_triggers.
type QrtzTriggersColumns struct {
	SchedName    string // 调度名称
	TriggerName  string // 触发器的名字
	TriggerGroup string // 触发器所属组的名字
	JobName      string // qrtz_job_details表job_name的外键
	JobGroup     string // qrtz_job_details表job_group的外键
	Description  string // 相关介绍
	NextFireTime string // 上一次触发时间（毫秒）
	PrevFireTime string // 下一次触发时间（默认为-1表示不触发）
	Priority     string // 优先级
	TriggerState string // 触发器状态
	TriggerType  string // 触发器的类型
	StartTime    string // 开始时间
	EndTime      string // 结束时间
	CalendarName string // 日程表名称
	MisfireInstr string // 补偿执行的策略
	JobData      string // 存放持久化job对象
}

// qrtzTriggersColumns holds the columns for the table qrtz_triggers.
var qrtzTriggersColumns = QrtzTriggersColumns{
	SchedName:    "sched_name",
	TriggerName:  "trigger_name",
	TriggerGroup: "trigger_group",
	JobName:      "job_name",
	JobGroup:     "job_group",
	Description:  "description",
	NextFireTime: "next_fire_time",
	PrevFireTime: "prev_fire_time",
	Priority:     "priority",
	TriggerState: "trigger_state",
	TriggerType:  "trigger_type",
	StartTime:    "start_time",
	EndTime:      "end_time",
	CalendarName: "calendar_name",
	MisfireInstr: "misfire_instr",
	JobData:      "job_data",
}

// NewQrtzTriggersDao creates and returns a new DAO object for table data access.
func NewQrtzTriggersDao(handlers ...gdb.ModelHandler) *QrtzTriggersDao {
	return &QrtzTriggersDao{
		group:    "default",
		table:    "qrtz_triggers",
		columns:  qrtzTriggersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzTriggersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzTriggersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzTriggersDao) Columns() QrtzTriggersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzTriggersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzTriggersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzTriggersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
