// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppSubTransferDao is the data access object for the table t_app_sub_transfer.
type AppSubTransferDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  AppSubTransferColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// AppSubTransferColumns defines and stores column names for the table t_app_sub_transfer.
type AppSubTransferColumns struct {
	Id         string //
	UserId     string // 用户id
	Email      string // 子账号邮箱
	EmailType  string // 1转入，2转出
	Type       string // 划转类型："SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	Tranid     string // 订单号
	Amount     string // 价格
	Asset      string // 货币类型
	Status     string // 1正常 2失败
	CreateTime string // 创建时间
}

// appSubTransferColumns holds the columns for the table t_app_sub_transfer.
var appSubTransferColumns = AppSubTransferColumns{
	Id:         "id",
	UserId:     "user_id",
	Email:      "email",
	EmailType:  "email_type",
	Type:       "type",
	Tranid:     "tranid",
	Amount:     "amount",
	Asset:      "asset",
	Status:     "status",
	CreateTime: "create_time",
}

// NewAppSubTransferDao creates and returns a new DAO object for table data access.
func NewAppSubTransferDao(handlers ...gdb.ModelHandler) *AppSubTransferDao {
	return &AppSubTransferDao{
		group:    "default",
		table:    "t_app_sub_transfer",
		columns:  appSubTransferColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppSubTransferDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppSubTransferDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppSubTransferDao) Columns() AppSubTransferColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppSubTransferDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppSubTransferDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppSubTransferDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
