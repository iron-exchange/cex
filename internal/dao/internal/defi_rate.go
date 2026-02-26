// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DefiRateDao is the data access object for the table t_defi_rate.
type DefiRateDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DefiRateColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DefiRateColumns defines and stores column names for the table t_defi_rate.
type DefiRateColumns struct {
	Id          string // id
	MinAmount   string // 最小金额
	MaxAmount   string // 最大金额
	Rate        string // 利率
	CreateBy    string // 创建者
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Remark      string // 备注
	SearchValue string //
}

// defiRateColumns holds the columns for the table t_defi_rate.
var defiRateColumns = DefiRateColumns{
	Id:          "id",
	MinAmount:   "min_amount",
	MaxAmount:   "max_amount",
	Rate:        "rate",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	SearchValue: "search_value",
}

// NewDefiRateDao creates and returns a new DAO object for table data access.
func NewDefiRateDao(handlers ...gdb.ModelHandler) *DefiRateDao {
	return &DefiRateDao{
		group:    "default",
		table:    "t_defi_rate",
		columns:  defiRateColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DefiRateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DefiRateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DefiRateDao) Columns() DefiRateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DefiRateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DefiRateDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DefiRateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
