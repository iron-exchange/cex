// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SecondCoinConfigDao is the data access object for the table t_second_coin_config.
type SecondCoinConfigDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  SecondCoinConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// SecondCoinConfigColumns defines and stores column names for the table t_second_coin_config.
type SecondCoinConfigColumns struct {
	Id          string // id
	Symbol      string // 合约交易对
	Market      string // 所属交易所
	Status      string // 是否启用 2关闭 1启用
	ShowFlag    string // 是否展示 2不展示 1展示
	Coin        string // 币种
	Sort        string // 排序
	CreateBy    string // 创建人
	CreateTime  string // 创建时间
	UpdateBy    string // 更新人
	UpdateTime  string // 更新时间
	Remark      string // 备注
	SearchValue string //
	Logo        string // 图标
	BaseCoin    string // 结算币种
	ShowSymbol  string // 展示币种
	Type        string // 币种类型 1 外汇  2虚拟币
}

// secondCoinConfigColumns holds the columns for the table t_second_coin_config.
var secondCoinConfigColumns = SecondCoinConfigColumns{
	Id:          "id",
	Symbol:      "symbol",
	Market:      "market",
	Status:      "status",
	ShowFlag:    "show_flag",
	Coin:        "coin",
	Sort:        "sort",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	SearchValue: "search_value",
	Logo:        "logo",
	BaseCoin:    "base_coin",
	ShowSymbol:  "show_symbol",
	Type:        "type",
}

// NewSecondCoinConfigDao creates and returns a new DAO object for table data access.
func NewSecondCoinConfigDao(handlers ...gdb.ModelHandler) *SecondCoinConfigDao {
	return &SecondCoinConfigDao{
		group:    "default",
		table:    "t_second_coin_config",
		columns:  secondCoinConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SecondCoinConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SecondCoinConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SecondCoinConfigDao) Columns() SecondCoinConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SecondCoinConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SecondCoinConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SecondCoinConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
