// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SymbolsDao is the data access object for the table t_symbols.
type SymbolsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SymbolsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SymbolsColumns defines and stores column names for the table t_symbols.
type SymbolsColumns struct {
	Id       string //
	Slug     string // 币种名称（ID）
	Symbol   string // 币种符号
	Fullname string // 币种全称
	LogoUrl  string // 图标链接
	Fiat     string // 是否法定货币
}

// symbolsColumns holds the columns for the table t_symbols.
var symbolsColumns = SymbolsColumns{
	Id:       "id",
	Slug:     "slug",
	Symbol:   "symbol",
	Fullname: "fullname",
	LogoUrl:  "logo_Url",
	Fiat:     "fiat",
}

// NewSymbolsDao creates and returns a new DAO object for table data access.
func NewSymbolsDao(handlers ...gdb.ModelHandler) *SymbolsDao {
	return &SymbolsDao{
		group:    "default",
		table:    "t_symbols",
		columns:  symbolsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SymbolsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SymbolsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SymbolsDao) Columns() SymbolsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SymbolsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SymbolsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SymbolsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
