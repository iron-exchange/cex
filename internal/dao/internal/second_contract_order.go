// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SecondContractOrderDao is the data access object for the table t_second_contract_order.
type SecondContractOrderDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  SecondContractOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// SecondContractOrderColumns defines and stores column names for the table t_second_contract_order.
type SecondContractOrderColumns struct {
	Id                 string //
	OrderNo            string // 订单号
	Symbol             string // 交易对
	Type               string // 类型
	UserId             string // 用户id
	UserAddress        string // 用户地址
	BetContent         string // 预测方向:0 涨 1跌
	OpenResult         string // 开奖结果
	Status             string // 订单状态 0参与中 1已开奖 2已撤销
	RateFlag           string // 是否全输
	BetAmount          string // 投注金额
	RewardAmount       string // 获奖金额
	CompensationAmount string // 赔偿金额
	CreateTime         string // 创建时间
	OpenPrice          string // 开盘价格
	ClosePrice         string // 关盘价格
	OpenTime           string // 开盘时间
	CloseTime          string // 关盘时间
	CoinSymbol         string // 交易币符号
	BaseSymbol         string // 结算币符号
	Sign               string // 订单标记 0正常  1包赢  2包输
	ManualIntervention string // 是否人工干预 0是 1否
	Rate               string //
	CreateBy           string //
	UpdateTime         string //
	UpdateBy           string //
	SearchValue        string //
	Remark             string //
	AdminParentIds     string // 后台代理ID
	IsHandling         string // 行锁
}

// secondContractOrderColumns holds the columns for the table t_second_contract_order.
var secondContractOrderColumns = SecondContractOrderColumns{
	Id:                 "id",
	OrderNo:            "order_no",
	Symbol:             "symbol",
	Type:               "type",
	UserId:             "user_id",
	UserAddress:        "user_address",
	BetContent:         "bet_content",
	OpenResult:         "open_result",
	Status:             "status",
	RateFlag:           "rate_flag",
	BetAmount:          "bet_amount",
	RewardAmount:       "reward_amount",
	CompensationAmount: "compensation_amount",
	CreateTime:         "create_time",
	OpenPrice:          "open_price",
	ClosePrice:         "close_price",
	OpenTime:           "open_time",
	CloseTime:          "close_time",
	CoinSymbol:         "coin_symbol",
	BaseSymbol:         "base_symbol",
	Sign:               "sign",
	ManualIntervention: "manual_intervention",
	Rate:               "rate",
	CreateBy:           "create_by",
	UpdateTime:         "update_time",
	UpdateBy:           "update_by",
	SearchValue:        "search_value",
	Remark:             "remark",
	AdminParentIds:     "admin_parent_ids",
	IsHandling:         "is_handling",
}

// NewSecondContractOrderDao creates and returns a new DAO object for table data access.
func NewSecondContractOrderDao(handlers ...gdb.ModelHandler) *SecondContractOrderDao {
	return &SecondContractOrderDao{
		group:    "default",
		table:    "t_second_contract_order",
		columns:  secondContractOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SecondContractOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SecondContractOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SecondContractOrderDao) Columns() SecondContractOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SecondContractOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SecondContractOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SecondContractOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
