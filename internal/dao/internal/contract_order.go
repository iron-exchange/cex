// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractOrderDao is the data access object for the table t_contract_order.
type ContractOrderDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ContractOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ContractOrderColumns defines and stores column names for the table t_contract_order.
type ContractOrderColumns struct {
	Id             string // 主键
	Type           string // (0 买多 1卖空)
	DelegateType   string // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status         string // 状态  0 （等待成交  1 完全成交  3已撤销）
	DelegateTotal  string // 委托总量
	DelegatePrice  string // 委托价格
	DealNum        string // 已成交量
	DealPrice      string // 成交价
	DelegateValue  string // 委托价值
	DealValue      string // 成交价值
	DelegateTime   string // 委托时间
	DealTime       string // 成交时间
	CoinSymbol     string // 交易币种
	CreateTime     string // 创建时间
	OrderNo        string // 订单编号
	UserId         string // 用户id
	UpdateTime     string // 更新时间
	Fee            string // 手续费
	BaseCoin       string // 基础币种（USDT）
	Leverage       string // 杠杆
	Symbol         string // 交易对
	AdminUserIds   string // 代理IDS
	AdminParentIds string //
}

// contractOrderColumns holds the columns for the table t_contract_order.
var contractOrderColumns = ContractOrderColumns{
	Id:             "id",
	Type:           "type",
	DelegateType:   "delegate_type",
	Status:         "status",
	DelegateTotal:  "delegate_total",
	DelegatePrice:  "delegate_price",
	DealNum:        "deal_num",
	DealPrice:      "deal_price",
	DelegateValue:  "delegate_value",
	DealValue:      "deal_value",
	DelegateTime:   "delegate_time",
	DealTime:       "deal_time",
	CoinSymbol:     "coin_symbol",
	CreateTime:     "create_time",
	OrderNo:        "order_no",
	UserId:         "user_id",
	UpdateTime:     "update_time",
	Fee:            "fee",
	BaseCoin:       "base_coin",
	Leverage:       "leverage",
	Symbol:         "symbol",
	AdminUserIds:   "admin_user_ids",
	AdminParentIds: "admin_parent_ids",
}

// NewContractOrderDao creates and returns a new DAO object for table data access.
func NewContractOrderDao(handlers ...gdb.ModelHandler) *ContractOrderDao {
	return &ContractOrderDao{
		group:    "default",
		table:    "t_contract_order",
		columns:  contractOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContractOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContractOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContractOrderDao) Columns() ContractOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContractOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContractOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ContractOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
