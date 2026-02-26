// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KlineSymbolDao is the data access object for the table t_kline_symbol.
type KlineSymbolDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  KlineSymbolColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// KlineSymbolColumns defines and stores column names for the table t_kline_symbol.
type KlineSymbolColumns struct {
	Id          string // id
	Market      string // 交易所
	Symbol      string // 币种简称
	Slug        string // 币种名称
	Status      string // 是否开启
	SearchValue string //
	Logo        string //
	Remark      string // 用户备注
	CreateBy    string //
	UpdateBy    string //
	UpdateTime  string // 更新时间
	CreateTime  string //
	ReferMarket string // 参考币种交易所
	ReferCoin   string // 参考币种
	Proportion  string // 价格百分比
}

// klineSymbolColumns holds the columns for the table t_kline_symbol.
var klineSymbolColumns = KlineSymbolColumns{
	Id:          "id",
	Market:      "market",
	Symbol:      "symbol",
	Slug:        "slug",
	Status:      "status",
	SearchValue: "search_value",
	Logo:        "logo",
	Remark:      "remark",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	CreateTime:  "create_time",
	ReferMarket: "refer_market",
	ReferCoin:   "refer_coin",
	Proportion:  "proportion",
}

// NewKlineSymbolDao creates and returns a new DAO object for table data access.
func NewKlineSymbolDao(handlers ...gdb.ModelHandler) *KlineSymbolDao {
	return &KlineSymbolDao{
		group:    "default",
		table:    "t_kline_symbol",
		columns:  klineSymbolColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *KlineSymbolDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *KlineSymbolDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *KlineSymbolDao) Columns() KlineSymbolColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *KlineSymbolDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *KlineSymbolDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *KlineSymbolDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
