// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpontaneousCoinDao is the data access object for the table t_spontaneous_coin.
type SpontaneousCoinDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  SpontaneousCoinColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// SpontaneousCoinColumns defines and stores column names for the table t_spontaneous_coin.
type SpontaneousCoinColumns struct {
	Id          string // 主键ID
	Coin        string // 币种
	Logo        string // 图标
	ReferCoin   string // 参考币种
	ReferMarket string // 参考币种交易所
	ShowSymbol  string // 展示名称
	Price       string // 初始价格（单位USDT）
	Proportion  string // 价格百分比
	CreateBy    string // 创建人
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Remark      string // 备注
}

// spontaneousCoinColumns holds the columns for the table t_spontaneous_coin.
var spontaneousCoinColumns = SpontaneousCoinColumns{
	Id:          "id",
	Coin:        "coin",
	Logo:        "logo",
	ReferCoin:   "refer_coin",
	ReferMarket: "refer_market",
	ShowSymbol:  "show_symbol",
	Price:       "price",
	Proportion:  "proportion",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
}

// NewSpontaneousCoinDao creates and returns a new DAO object for table data access.
func NewSpontaneousCoinDao(handlers ...gdb.ModelHandler) *SpontaneousCoinDao {
	return &SpontaneousCoinDao{
		group:    "default",
		table:    "t_spontaneous_coin",
		columns:  spontaneousCoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpontaneousCoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpontaneousCoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpontaneousCoinDao) Columns() SpontaneousCoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpontaneousCoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpontaneousCoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SpontaneousCoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
