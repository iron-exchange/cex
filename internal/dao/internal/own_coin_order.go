// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OwnCoinOrderDao is the data access object for the table t_own_coin_order.
type OwnCoinOrderDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  OwnCoinOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// OwnCoinOrderColumns defines and stores column names for the table t_own_coin_order.
type OwnCoinOrderColumns struct {
	Id             string // 主键
	UserId         string // 用户ID
	OrderId        string // 订单ID
	OwnId          string // 申购币种ID
	OwnCoin        string // 申购币种
	Amount         string // 申购额（usdt）
	Number         string // 申购数量
	Price          string // 申购价
	Status         string // 状态
	AdminUserIds   string // 上级用户IDS
	AdminParentIds string // 上级后台用户IDS
	CreateTime     string //
	UpdateTime     string //
	CreateBy       string //
	UpdateBy       string //
	Remark         string //
}

// ownCoinOrderColumns holds the columns for the table t_own_coin_order.
var ownCoinOrderColumns = OwnCoinOrderColumns{
	Id:             "id",
	UserId:         "user_id",
	OrderId:        "order_id",
	OwnId:          "own_id",
	OwnCoin:        "own_coin",
	Amount:         "amount",
	Number:         "number",
	Price:          "price",
	Status:         "status",
	AdminUserIds:   "admin_user_ids",
	AdminParentIds: "admin_parent_ids",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
	Remark:         "remark",
}

// NewOwnCoinOrderDao creates and returns a new DAO object for table data access.
func NewOwnCoinOrderDao(handlers ...gdb.ModelHandler) *OwnCoinOrderDao {
	return &OwnCoinOrderDao{
		group:    "default",
		table:    "t_own_coin_order",
		columns:  ownCoinOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OwnCoinOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OwnCoinOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OwnCoinOrderDao) Columns() OwnCoinOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OwnCoinOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OwnCoinOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OwnCoinOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
