// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CurrencySymbolDao is the data access object for the table t_currency_symbol.
type CurrencySymbolDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  CurrencySymbolColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// CurrencySymbolColumns defines and stores column names for the table t_currency_symbol.
type CurrencySymbolColumns struct {
	Id            string // 主键id
	Symbol        string // 交易对
	ShowSymbol    string // 展示交易对
	Coin          string // 交易币种
	BaseCoin      string // 结算币种
	FeeRate       string // 手续费率
	CoinPrecision string // 交易币种精度
	BasePrecision string // 结算币种精度
	SellMin       string // 最低卖单价
	BuyMax        string // 最高买单价
	OrderMin      string // 最小下单量
	OrderMax      string // 最大下单量
	Enable        string // 启用禁用  1=启用 2=禁用
	IsShow        string // 前端是否显示 1=显示  2=隐藏
	IsDeal        string // 是否可交易 1=是 2=否
	MarketBuy     string // 市价买 1=可以 2=不可以
	MarketSell    string // 市价卖 1=可以 2=不可以
	LimitedBuy    string // 限价买 1=可以 2=不可以
	LimitedSell   string // 限价卖 1=可以 2=不可以
	Logo          string // 图标
	Market        string // 交易所
	CreateBy      string //
	CreateTime    string //
	UpdateBy      string //
	UpdateTime    string //
	SearchValue   string //
	Remark        string //
	MinSell       string // 最低卖出量
}

// currencySymbolColumns holds the columns for the table t_currency_symbol.
var currencySymbolColumns = CurrencySymbolColumns{
	Id:            "id",
	Symbol:        "symbol",
	ShowSymbol:    "show_symbol",
	Coin:          "coin",
	BaseCoin:      "base_coin",
	FeeRate:       "fee_rate",
	CoinPrecision: "coin_precision",
	BasePrecision: "base_precision",
	SellMin:       "sell_min",
	BuyMax:        "buy_max",
	OrderMin:      "order_min",
	OrderMax:      "order_max",
	Enable:        "enable",
	IsShow:        "is_show",
	IsDeal:        "is_deal",
	MarketBuy:     "market_buy",
	MarketSell:    "market_sell",
	LimitedBuy:    "limited_buy",
	LimitedSell:   "limited_sell",
	Logo:          "logo",
	Market:        "market",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
	SearchValue:   "search_value",
	Remark:        "remark",
	MinSell:       "min_sell",
}

// NewCurrencySymbolDao creates and returns a new DAO object for table data access.
func NewCurrencySymbolDao(handlers ...gdb.ModelHandler) *CurrencySymbolDao {
	return &CurrencySymbolDao{
		group:    "default",
		table:    "t_currency_symbol",
		columns:  currencySymbolColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CurrencySymbolDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CurrencySymbolDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CurrencySymbolDao) Columns() CurrencySymbolColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CurrencySymbolDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CurrencySymbolDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CurrencySymbolDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
