// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MineUserDao is the data access object for the table t_mine_user.
type MineUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MineUserColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MineUserColumns defines and stores column names for the table t_mine_user.
type MineUserColumns struct {
	UserId    string // 用户id
	Id        string // 挖矿产品id
	TimeLimit string // 限购次数
}

// mineUserColumns holds the columns for the table t_mine_user.
var mineUserColumns = MineUserColumns{
	UserId:    "user_id",
	Id:        "id",
	TimeLimit: "time_limit",
}

// NewMineUserDao creates and returns a new DAO object for table data access.
func NewMineUserDao(handlers ...gdb.ModelHandler) *MineUserDao {
	return &MineUserDao{
		group:    "default",
		table:    "t_mine_user",
		columns:  mineUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MineUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MineUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MineUserDao) Columns() MineUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MineUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MineUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MineUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
