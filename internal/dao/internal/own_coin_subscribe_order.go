// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OwnCoinSubscribeOrderDao is the data access object for the table t_own_coin_subscribe_order.
type OwnCoinSubscribeOrderDao struct {
	table    string                       // table is the underlying table name of the DAO.
	group    string                       // group is the database configuration group name of the current DAO.
	columns  OwnCoinSubscribeOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler           // handlers for customized model modification.
}

// OwnCoinSubscribeOrderColumns defines and stores column names for the table t_own_coin_subscribe_order.
type OwnCoinSubscribeOrderColumns struct {
	Id          string // 主键
	SubscribeId string // 订阅单号ID
	UserId      string // 用户ID
	OrderId     string // 订单ID
	OwnId       string // 申购币种ID
	OwnCoin     string // 申购币种
	AmountLimit string // 申购额（usdt）
	NumLimit    string // 申购数量上限
	Price       string // 申购价
	Status      string // 状态，1订阅中、2订阅成功、3成功消息推送完成
	Remark      string // 备注
	CreateTime  string //
	UpdateTime  string //
	CreateBy    string //
}

// ownCoinSubscribeOrderColumns holds the columns for the table t_own_coin_subscribe_order.
var ownCoinSubscribeOrderColumns = OwnCoinSubscribeOrderColumns{
	Id:          "id",
	SubscribeId: "subscribe_id",
	UserId:      "user_id",
	OrderId:     "order_id",
	OwnId:       "own_id",
	OwnCoin:     "own_coin",
	AmountLimit: "amount_limit",
	NumLimit:    "num_limit",
	Price:       "price",
	Status:      "status",
	Remark:      "remark",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	CreateBy:    "create_by",
}

// NewOwnCoinSubscribeOrderDao creates and returns a new DAO object for table data access.
func NewOwnCoinSubscribeOrderDao(handlers ...gdb.ModelHandler) *OwnCoinSubscribeOrderDao {
	return &OwnCoinSubscribeOrderDao{
		group:    "default",
		table:    "t_own_coin_subscribe_order",
		columns:  ownCoinSubscribeOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OwnCoinSubscribeOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OwnCoinSubscribeOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OwnCoinSubscribeOrderDao) Columns() OwnCoinSubscribeOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OwnCoinSubscribeOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OwnCoinSubscribeOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OwnCoinSubscribeOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
