// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzJobDetailsDao is the data access object for the table qrtz_job_details.
type QrtzJobDetailsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  QrtzJobDetailsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// QrtzJobDetailsColumns defines and stores column names for the table qrtz_job_details.
type QrtzJobDetailsColumns struct {
	SchedName        string // 调度名称
	JobName          string // 任务名称
	JobGroup         string // 任务组名
	Description      string // 相关介绍
	JobClassName     string // 执行任务类名称
	IsDurable        string // 是否持久化
	IsNonconcurrent  string // 是否并发
	IsUpdateData     string // 是否更新数据
	RequestsRecovery string // 是否接受恢复执行
	JobData          string // 存放持久化job对象
}

// qrtzJobDetailsColumns holds the columns for the table qrtz_job_details.
var qrtzJobDetailsColumns = QrtzJobDetailsColumns{
	SchedName:        "sched_name",
	JobName:          "job_name",
	JobGroup:         "job_group",
	Description:      "description",
	JobClassName:     "job_class_name",
	IsDurable:        "is_durable",
	IsNonconcurrent:  "is_nonconcurrent",
	IsUpdateData:     "is_update_data",
	RequestsRecovery: "requests_recovery",
	JobData:          "job_data",
}

// NewQrtzJobDetailsDao creates and returns a new DAO object for table data access.
func NewQrtzJobDetailsDao(handlers ...gdb.ModelHandler) *QrtzJobDetailsDao {
	return &QrtzJobDetailsDao{
		group:    "default",
		table:    "qrtz_job_details",
		columns:  qrtzJobDetailsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzJobDetailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzJobDetailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzJobDetailsDao) Columns() QrtzJobDetailsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzJobDetailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzJobDetailsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzJobDetailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
