// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractLossDao is the data access object for the table t_contract_loss.
type ContractLossDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ContractLossColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ContractLossColumns defines and stores column names for the table t_contract_loss.
type ContractLossColumns struct {
	Id                string // 主键
	DelegateType      string // 委托类型（0 限价 1 市价）
	Status            string // 状态  0  正常 1 删除  2 撤销
	PositionId        string // 仓位ID
	UserId            string // 用户id
	EarnPrice         string // 止盈触发价
	LosePrice         string // 止损触发价
	CreateTime        string // 创建时间
	EarnDelegatePrice string // 止盈委托价
	LoseDelegatePrice string // 止损委托价
	EarnNumber        string // 止盈数量
	LoseNumber        string // 止损数量
	LossType          string // 0 止盈    1止损
	UpdateTime        string // 更新时间
	Type              string //
	Leverage          string //
	Symbol            string //
}

// contractLossColumns holds the columns for the table t_contract_loss.
var contractLossColumns = ContractLossColumns{
	Id:                "id",
	DelegateType:      "delegate_type",
	Status:            "status",
	PositionId:        "position_id",
	UserId:            "user_id",
	EarnPrice:         "earn_price",
	LosePrice:         "lose_price",
	CreateTime:        "create_time",
	EarnDelegatePrice: "earn_delegate_price",
	LoseDelegatePrice: "lose_delegate_price",
	EarnNumber:        "earn_number",
	LoseNumber:        "lose_number",
	LossType:          "loss_type",
	UpdateTime:        "update_time",
	Type:              "type",
	Leverage:          "leverage",
	Symbol:            "symbol",
}

// NewContractLossDao creates and returns a new DAO object for table data access.
func NewContractLossDao(handlers ...gdb.ModelHandler) *ContractLossDao {
	return &ContractLossDao{
		group:    "default",
		table:    "t_contract_loss",
		columns:  contractLossColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContractLossDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContractLossDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContractLossDao) Columns() ContractLossColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContractLossDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContractLossDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ContractLossDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
