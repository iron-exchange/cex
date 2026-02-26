// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DefiOrderDao is the data access object for the table t_defi_order.
type DefiOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DefiOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DefiOrderColumns defines and stores column names for the table t_defi_order.
type DefiOrderColumns struct {
	Id             string // id
	Amount         string // 收益金额
	TotleAmount    string // 钱包金额
	Rate           string // 收益率
	CreateBy       string // 创建者
	CreateTime     string // 创建时间
	UpdateBy       string // 更新者
	UpdateTime     string // 更新时间
	Remark         string // 备注
	SearchValue    string //
	UserId         string // 用户id
	AdminParentIds string // 代理ids
}

// defiOrderColumns holds the columns for the table t_defi_order.
var defiOrderColumns = DefiOrderColumns{
	Id:             "id",
	Amount:         "amount",
	TotleAmount:    "totle_amount",
	Rate:           "rate",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	SearchValue:    "search_value",
	UserId:         "user_id",
	AdminParentIds: "admin_parent_ids",
}

// NewDefiOrderDao creates and returns a new DAO object for table data access.
func NewDefiOrderDao(handlers ...gdb.ModelHandler) *DefiOrderDao {
	return &DefiOrderDao{
		group:    "default",
		table:    "t_defi_order",
		columns:  defiOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DefiOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DefiOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DefiOrderDao) Columns() DefiOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DefiOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DefiOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DefiOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
