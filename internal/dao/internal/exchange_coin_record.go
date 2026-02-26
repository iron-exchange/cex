// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ExchangeCoinRecordDao is the data access object for the table t_exchange_coin_record.
type ExchangeCoinRecordDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  ExchangeCoinRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// ExchangeCoinRecordColumns defines and stores column names for the table t_exchange_coin_record.
type ExchangeCoinRecordColumns struct {
	Id             string //
	FromCoin       string //
	ToCoin         string //
	UserId         string // 用户id
	Username       string // 用户名称
	Address        string // 用户地址
	Status         string // 兑换状态0:已提交;1:成功;2失败
	Amount         string // 金额
	ThirdRate      string // 三方汇率
	SystemRate     string // 系统汇率
	AdminParentIds string //
	CreateBy       string //
	CreateTime     string //
	UpdateBy       string //
	UpdateTime     string //
	Remark         string //
	SearchValue    string //
}

// exchangeCoinRecordColumns holds the columns for the table t_exchange_coin_record.
var exchangeCoinRecordColumns = ExchangeCoinRecordColumns{
	Id:             "id",
	FromCoin:       "from_coin",
	ToCoin:         "to_coin",
	UserId:         "user_id",
	Username:       "username",
	Address:        "address",
	Status:         "status",
	Amount:         "amount",
	ThirdRate:      "third_rate",
	SystemRate:     "system_rate",
	AdminParentIds: "admin_parent_ids",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	SearchValue:    "search_value",
}

// NewExchangeCoinRecordDao creates and returns a new DAO object for table data access.
func NewExchangeCoinRecordDao(handlers ...gdb.ModelHandler) *ExchangeCoinRecordDao {
	return &ExchangeCoinRecordDao{
		group:    "default",
		table:    "t_exchange_coin_record",
		columns:  exchangeCoinRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ExchangeCoinRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ExchangeCoinRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ExchangeCoinRecordDao) Columns() ExchangeCoinRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ExchangeCoinRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ExchangeCoinRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ExchangeCoinRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
