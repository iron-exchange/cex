// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCoinDao is the data access object for the table t_user_coin.
type UserCoinDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserCoinColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserCoinColumns defines and stores column names for the table t_user_coin.
type UserCoinColumns struct {
	Id     string // 主键
	UserId string // 用户id
	Coin   string // 币种
	Icon   string // 图标
}

// userCoinColumns holds the columns for the table t_user_coin.
var userCoinColumns = UserCoinColumns{
	Id:     "id",
	UserId: "user_id",
	Coin:   "coin",
	Icon:   "icon",
}

// NewUserCoinDao creates and returns a new DAO object for table data access.
func NewUserCoinDao(handlers ...gdb.ModelHandler) *UserCoinDao {
	return &UserCoinDao{
		group:    "default",
		table:    "t_user_coin",
		columns:  userCoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserCoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserCoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserCoinDao) Columns() UserCoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserCoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserCoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserCoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
