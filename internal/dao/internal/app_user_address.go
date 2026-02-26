// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserAddressDao is the data access object for the table t_app_user_address.
type AppUserAddressDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  AppUserAddressColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// AppUserAddressColumns defines and stores column names for the table t_app_user_address.
type AppUserAddressColumns struct {
	Id           string //
	UserId       string //
	Symbol       string // 钱包类型
	Address      string // 钱包地址
	BinanceEmail string // 币安子账号地址
}

// appUserAddressColumns holds the columns for the table t_app_user_address.
var appUserAddressColumns = AppUserAddressColumns{
	Id:           "id",
	UserId:       "user_id",
	Symbol:       "symbol",
	Address:      "address",
	BinanceEmail: "binance_email",
}

// NewAppUserAddressDao creates and returns a new DAO object for table data access.
func NewAppUserAddressDao(handlers ...gdb.ModelHandler) *AppUserAddressDao {
	return &AppUserAddressDao{
		group:    "default",
		table:    "t_app_user_address",
		columns:  appUserAddressColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppUserAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppUserAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppUserAddressDao) Columns() AppUserAddressColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppUserAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppUserAddressDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppUserAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
