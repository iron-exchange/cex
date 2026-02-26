// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BotKlineModelInfoDao is the data access object for the table t_bot_kline_model_info.
type BotKlineModelInfoDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  BotKlineModelInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// BotKlineModelInfoColumns defines and stores column names for the table t_bot_kline_model_info.
type BotKlineModelInfoColumns struct {
	Id       string // id
	ModelId  string // t_bot_kline_model 的主键
	DateTime string // 时间戳
	Open     string // 开盘价
	Close    string // 封盘价
	High     string // 最高价
	Low      string // 最低价
	X        string // x轴
	Y        string // y轴
}

// botKlineModelInfoColumns holds the columns for the table t_bot_kline_model_info.
var botKlineModelInfoColumns = BotKlineModelInfoColumns{
	Id:       "id",
	ModelId:  "model_id",
	DateTime: "date_time",
	Open:     "open",
	Close:    "close",
	High:     "high",
	Low:      "low",
	X:        "x",
	Y:        "y",
}

// NewBotKlineModelInfoDao creates and returns a new DAO object for table data access.
func NewBotKlineModelInfoDao(handlers ...gdb.ModelHandler) *BotKlineModelInfoDao {
	return &BotKlineModelInfoDao{
		group:    "default",
		table:    "t_bot_kline_model_info",
		columns:  botKlineModelInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BotKlineModelInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BotKlineModelInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BotKlineModelInfoDao) Columns() BotKlineModelInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BotKlineModelInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BotKlineModelInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BotKlineModelInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
