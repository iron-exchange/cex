// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectionOrderDao is the data access object for the table t_collection_order.
type CollectionOrderDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  CollectionOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// CollectionOrderColumns defines and stores column names for the table t_collection_order.
type CollectionOrderColumns struct {
	Id          string // 主键ID
	OrderId     string // 订单号
	UserId      string // 用户ID
	Address     string // 归集地址
	Chain       string // 地址类型
	Hash        string // hash
	Coin        string // 币种
	Amount      string // 归集金额
	Status      string // 1  进行中   2 归集成功  3 归集失败
	ClientName  string // 客户端名称
	CreateTime  string // 创建时间
	CreateBy    string // 创建人
	UpdateTime  string // 修改时间
	UpdateBy    string // 修改人
	Remark      string // 备注
	SearchValue string //
}

// collectionOrderColumns holds the columns for the table t_collection_order.
var collectionOrderColumns = CollectionOrderColumns{
	Id:          "id",
	OrderId:     "order_id",
	UserId:      "user_id",
	Address:     "address",
	Chain:       "chain",
	Hash:        "hash",
	Coin:        "coin",
	Amount:      "amount",
	Status:      "status",
	ClientName:  "client_name",
	CreateTime:  "create_time",
	CreateBy:    "create_by",
	UpdateTime:  "update_time",
	UpdateBy:    "update_by",
	Remark:      "remark",
	SearchValue: "search_value",
}

// NewCollectionOrderDao creates and returns a new DAO object for table data access.
func NewCollectionOrderDao(handlers ...gdb.ModelHandler) *CollectionOrderDao {
	return &CollectionOrderDao{
		group:    "default",
		table:    "t_collection_order",
		columns:  collectionOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CollectionOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CollectionOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CollectionOrderDao) Columns() CollectionOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CollectionOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CollectionOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CollectionOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
