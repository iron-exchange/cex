// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MarketsDao is the data access object for the table t_markets.
type MarketsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MarketsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MarketsColumns defines and stores column names for the table t_markets.
type MarketsColumns struct {
	Id          string //
	Slug        string // 交易所名称(ID)
	Fullname    string // 交易所全称
	WebsiteUrl  string // 交易所官网链接
	Status      string // 状态: [enable, disable]. disable为停止更新数据
	Kline       string // 是否接入K线数据
	Spot        string // 是否支持现货
	Futures     string // 是否支持期货
	SearchValue string //
}

// marketsColumns holds the columns for the table t_markets.
var marketsColumns = MarketsColumns{
	Id:          "id",
	Slug:        "slug",
	Fullname:    "fullname",
	WebsiteUrl:  "website_url",
	Status:      "status",
	Kline:       "kline",
	Spot:        "spot",
	Futures:     "futures",
	SearchValue: "search_value",
}

// NewMarketsDao creates and returns a new DAO object for table data access.
func NewMarketsDao(handlers ...gdb.ModelHandler) *MarketsDao {
	return &MarketsDao{
		group:    "default",
		table:    "t_markets",
		columns:  marketsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MarketsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MarketsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MarketsDao) Columns() MarketsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MarketsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MarketsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MarketsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
