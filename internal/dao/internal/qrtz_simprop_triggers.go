// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzSimpropTriggersDao is the data access object for the table qrtz_simprop_triggers.
type QrtzSimpropTriggersDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  QrtzSimpropTriggersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// QrtzSimpropTriggersColumns defines and stores column names for the table qrtz_simprop_triggers.
type QrtzSimpropTriggersColumns struct {
	SchedName    string // 调度名称
	TriggerName  string // qrtz_triggers表trigger_name的外键
	TriggerGroup string // qrtz_triggers表trigger_group的外键
	StrProp1     string // String类型的trigger的第一个参数
	StrProp2     string // String类型的trigger的第二个参数
	StrProp3     string // String类型的trigger的第三个参数
	IntProp1     string // int类型的trigger的第一个参数
	IntProp2     string // int类型的trigger的第二个参数
	LongProp1    string // long类型的trigger的第一个参数
	LongProp2    string // long类型的trigger的第二个参数
	DecProp1     string // decimal类型的trigger的第一个参数
	DecProp2     string // decimal类型的trigger的第二个参数
	BoolProp1    string // Boolean类型的trigger的第一个参数
	BoolProp2    string // Boolean类型的trigger的第二个参数
}

// qrtzSimpropTriggersColumns holds the columns for the table qrtz_simprop_triggers.
var qrtzSimpropTriggersColumns = QrtzSimpropTriggersColumns{
	SchedName:    "sched_name",
	TriggerName:  "trigger_name",
	TriggerGroup: "trigger_group",
	StrProp1:     "str_prop_1",
	StrProp2:     "str_prop_2",
	StrProp3:     "str_prop_3",
	IntProp1:     "int_prop_1",
	IntProp2:     "int_prop_2",
	LongProp1:    "long_prop_1",
	LongProp2:    "long_prop_2",
	DecProp1:     "dec_prop_1",
	DecProp2:     "dec_prop_2",
	BoolProp1:    "bool_prop_1",
	BoolProp2:    "bool_prop_2",
}

// NewQrtzSimpropTriggersDao creates and returns a new DAO object for table data access.
func NewQrtzSimpropTriggersDao(handlers ...gdb.ModelHandler) *QrtzSimpropTriggersDao {
	return &QrtzSimpropTriggersDao{
		group:    "default",
		table:    "qrtz_simprop_triggers",
		columns:  qrtzSimpropTriggersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrtzSimpropTriggersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrtzSimpropTriggersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrtzSimpropTriggersDao) Columns() QrtzSimpropTriggersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrtzSimpropTriggersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrtzSimpropTriggersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *QrtzSimpropTriggersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
