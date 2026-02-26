// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CurrencyOrderDao is the data access object for the table t_currency_order.
type CurrencyOrderDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  CurrencyOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// CurrencyOrderColumns defines and stores column names for the table t_currency_order.
type CurrencyOrderColumns struct {
	Id             string // 主键
	UserId         string // 用户id
	Type           string // (0 买入 1卖出)
	DelegateType   string // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status         string // 状态  0 （等待成交  1 完全成交  3已撤销）
	OrderNo        string // 订单编号
	Symbol         string // 交易币种
	Coin           string // 结算币种
	DelegateTotal  string // 委托总量
	DelegatePrice  string // 委托价格
	DelegateValue  string // 委托价值
	DelegateTime   string // 委托时间
	DealNum        string // 成交总量
	DealPrice      string // 成交价格
	DealValue      string // 成交价值
	DealTime       string // 成交时间
	Fee            string // 手续费
	AdminParentIds string // 后台代理ids
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
	SearchValue    string //
	CreateBy       string //
	UpdateBy       string //
	Remark         string //
}

// currencyOrderColumns holds the columns for the table t_currency_order.
var currencyOrderColumns = CurrencyOrderColumns{
	Id:             "id",
	UserId:         "user_id",
	Type:           "type",
	DelegateType:   "delegate_type",
	Status:         "status",
	OrderNo:        "order_no",
	Symbol:         "symbol",
	Coin:           "coin",
	DelegateTotal:  "delegate_total",
	DelegatePrice:  "delegate_price",
	DelegateValue:  "delegate_value",
	DelegateTime:   "delegate_time",
	DealNum:        "deal_num",
	DealPrice:      "deal_price",
	DealValue:      "deal_value",
	DealTime:       "deal_time",
	Fee:            "fee",
	AdminParentIds: "admin_parent_ids",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	SearchValue:    "search_value",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
	Remark:         "remark",
}

// NewCurrencyOrderDao creates and returns a new DAO object for table data access.
func NewCurrencyOrderDao(handlers ...gdb.ModelHandler) *CurrencyOrderDao {
	return &CurrencyOrderDao{
		group:    "default",
		table:    "t_currency_order",
		columns:  currencyOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CurrencyOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CurrencyOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CurrencyOrderDao) Columns() CurrencyOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CurrencyOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CurrencyOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CurrencyOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
