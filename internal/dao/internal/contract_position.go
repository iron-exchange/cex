// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractPositionDao is the data access object for the table t_contract_position.
type ContractPositionDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ContractPositionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ContractPositionColumns defines and stores column names for the table t_contract_position.
type ContractPositionColumns struct {
	Id               string // 主键
	Type             string // (0 买多 1卖空)
	DelegateType     string // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status           string // 状态  0 （等待成交  1 完全成交
	Amount           string // 保证金
	OpenNum          string // 持仓数量
	OpenPrice        string // 开仓均价
	ClosePrice       string // 预计强平价
	OrderNo          string // 仓位编号
	UserId           string // 用户id
	OpenFee          string // 开仓手续费
	Leverage         string // 杠杆
	Symbol           string // 交易对
	CreateTime       string // 创建时间
	AdjustAmount     string // 调整保证金
	Earn             string // 收益
	DealPrice        string // 成交价
	DealNum          string // 成交量
	DealTime         string // 成交时间
	SellFee          string // 卖出手续费
	RemainMargin     string // 剩余保证金
	AssetFee         string // 周期手续费
	EntrustmentValue string //
	DealValue        string //
	UpdateTime       string //
	AdminParentIds   string // 代理IDS
	AuditStatus      string // 审核
	DeliveryDays     string // 交割时间
	MinMargin        string // 最小保证金
	LossRate         string // 止损率
	EarnRate         string // 止盈率
	SubTime          string // 提交时间
}

// contractPositionColumns holds the columns for the table t_contract_position.
var contractPositionColumns = ContractPositionColumns{
	Id:               "id",
	Type:             "type",
	DelegateType:     "delegate_type",
	Status:           "status",
	Amount:           "amount",
	OpenNum:          "open_num",
	OpenPrice:        "open_price",
	ClosePrice:       "close_price",
	OrderNo:          "order_no",
	UserId:           "user_id",
	OpenFee:          "open_fee",
	Leverage:         "leverage",
	Symbol:           "symbol",
	CreateTime:       "create_time",
	AdjustAmount:     "adjust_amount",
	Earn:             "earn",
	DealPrice:        "deal_price",
	DealNum:          "deal_num",
	DealTime:         "deal_time",
	SellFee:          "sell_fee",
	RemainMargin:     "remain_margin",
	AssetFee:         "asset_fee",
	EntrustmentValue: "entrustment_value",
	DealValue:        "deal_value",
	UpdateTime:       "update_time",
	AdminParentIds:   "admin_parent_ids",
	AuditStatus:      "audit_status",
	DeliveryDays:     "delivery_days",
	MinMargin:        "min_margin",
	LossRate:         "loss_rate",
	EarnRate:         "earn_rate",
	SubTime:          "sub_time",
}

// NewContractPositionDao creates and returns a new DAO object for table data access.
func NewContractPositionDao(handlers ...gdb.ModelHandler) *ContractPositionDao {
	return &ContractPositionDao{
		group:    "default",
		table:    "t_contract_position",
		columns:  contractPositionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContractPositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContractPositionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContractPositionDao) Columns() ContractPositionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContractPositionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContractPositionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ContractPositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
