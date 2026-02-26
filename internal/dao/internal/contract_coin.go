// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractCoinDao is the data access object for the table t_contract_coin.
type ContractCoinDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ContractCoinColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ContractCoinColumns defines and stores column names for the table t_contract_coin.
type ContractCoinColumns struct {
	Id               string //
	Symbol           string // 交易对
	Coin             string // 币种
	BaseCoin         string // 基础币种
	ShareNumber      string // 合约面值（1手多少 如 1手=0.01BTC）
	Leverage         string // 杠杆倍数
	Enable           string // 0 启用  1 禁止
	Visible          string // 前端显示0启用 1 禁止
	Exchangeable     string // 是否可交易（0 可以 1 禁止）
	EnableOpenSell   string // 开空  （0  是  1 否）
	EnableOpenBuy    string // 开多  （0  是  1 否）
	EnableMarketSell string // 市价开空（0 是 1否）
	EnableMarketBuy  string // 市价开多（0 是 1否）
	OpenFee          string // 开仓手续费
	CloseFee         string // 平仓手续费
	UsdtRate         string // 资金费率
	IntervalHour     string // 资金周期
	CoinScale        string // 币种小数精度
	BaseScale        string // 基础币小数精度
	MinShare         string // 最小数（以手为单位 ）
	MaxShare         string // 最大数（以手为单位 ）
	TotalProfit      string // 平台收益
	Sort             string // 排序字段
	CreateTime       string //
	UpdateTime       string //
	ShowSymbol       string // 显示币种
	Logo             string //
	Market           string //
	DeliveryDays     string // 交割时间
	MinMargin        string // 最小保证金
	LossRate         string // 止损率
	EarnRate         string // 止盈率
	FloatProfit      string // 浮动盈利点
	ProfitLoss       string // 浮动盈亏
}

// contractCoinColumns holds the columns for the table t_contract_coin.
var contractCoinColumns = ContractCoinColumns{
	Id:               "id",
	Symbol:           "symbol",
	Coin:             "coin",
	BaseCoin:         "base_coin",
	ShareNumber:      "share_number",
	Leverage:         "leverage",
	Enable:           "enable",
	Visible:          "visible",
	Exchangeable:     "exchangeable",
	EnableOpenSell:   "enable_open_sell",
	EnableOpenBuy:    "enable_open_buy",
	EnableMarketSell: "enable_market_sell",
	EnableMarketBuy:  "enable_market_buy",
	OpenFee:          "open_fee",
	CloseFee:         "close_fee",
	UsdtRate:         "usdt_rate",
	IntervalHour:     "interval_hour",
	CoinScale:        "coin_scale",
	BaseScale:        "base_scale",
	MinShare:         "min_share",
	MaxShare:         "max_share",
	TotalProfit:      "total_profit",
	Sort:             "sort",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	ShowSymbol:       "show_symbol",
	Logo:             "logo",
	Market:           "market",
	DeliveryDays:     "delivery_days",
	MinMargin:        "min_margin",
	LossRate:         "loss_rate",
	EarnRate:         "earn_rate",
	FloatProfit:      "float_profit",
	ProfitLoss:       "profit_loss",
}

// NewContractCoinDao creates and returns a new DAO object for table data access.
func NewContractCoinDao(handlers ...gdb.ModelHandler) *ContractCoinDao {
	return &ContractCoinDao{
		group:    "default",
		table:    "t_contract_coin",
		columns:  contractCoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContractCoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContractCoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContractCoinDao) Columns() ContractCoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContractCoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContractCoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ContractCoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
