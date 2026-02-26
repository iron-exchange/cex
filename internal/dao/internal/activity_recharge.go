// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRechargeDao is the data access object for the table t_activity_recharge.
type ActivityRechargeDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ActivityRechargeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ActivityRechargeColumns defines and stores column names for the table t_activity_recharge.
type ActivityRechargeColumns struct {
	Id          string // id
	OnOff       string // 0-关闭 1-开启
	RechargePro string // 充值返点比例
	MaxRebate   string // 充值返点最大值
	CreateBy    string //
	CreateTime  string // 创建时间
	UpdateBy    string //
	UpdateTime  string // 更新时间
	SearchValue string //
}

// activityRechargeColumns holds the columns for the table t_activity_recharge.
var activityRechargeColumns = ActivityRechargeColumns{
	Id:          "id",
	OnOff:       "on_off",
	RechargePro: "recharge_pro",
	MaxRebate:   "max_rebate",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	SearchValue: "search_value",
}

// NewActivityRechargeDao creates and returns a new DAO object for table data access.
func NewActivityRechargeDao(handlers ...gdb.ModelHandler) *ActivityRechargeDao {
	return &ActivityRechargeDao{
		group:    "default",
		table:    "t_activity_recharge",
		columns:  activityRechargeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityRechargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityRechargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityRechargeDao) Columns() ActivityRechargeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityRechargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityRechargeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityRechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
