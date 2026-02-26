// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BotKlineModelDao is the data access object for the table t_bot_kline_model.
type BotKlineModelDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  BotKlineModelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// BotKlineModelColumns defines and stores column names for the table t_bot_kline_model.
type BotKlineModelColumns struct {
	Id            string // id
	Decline       string // 最大跌幅
	Granularity   string // 控制粒度
	Increase      string // 最大涨幅
	Model         string // 控盘策略
	PricePencent  string // 浮动比例
	Symbol        string // 交易对
	CreateBy      string // 创建人
	CreateTime    string // 创建时间
	UpdateBy      string // 修改人
	UpdateTime    string // 更新时间
	SearchValue   string // 值
	BeginTime     string // 开始时间
	EndTime       string // 结束时间
	LineChartData string //
	Remark        string //
	ConPrice      string //
}

// botKlineModelColumns holds the columns for the table t_bot_kline_model.
var botKlineModelColumns = BotKlineModelColumns{
	Id:            "id",
	Decline:       "decline",
	Granularity:   "granularity",
	Increase:      "increase",
	Model:         "model",
	PricePencent:  "price_pencent",
	Symbol:        "symbol",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
	SearchValue:   "search_value",
	BeginTime:     "begin_time",
	EndTime:       "end_time",
	LineChartData: "line_chart_data",
	Remark:        "remark",
	ConPrice:      "con_price",
}

// NewBotKlineModelDao creates and returns a new DAO object for table data access.
func NewBotKlineModelDao(handlers ...gdb.ModelHandler) *BotKlineModelDao {
	return &BotKlineModelDao{
		group:    "default",
		table:    "t_bot_kline_model",
		columns:  botKlineModelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BotKlineModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BotKlineModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BotKlineModelDao) Columns() BotKlineModelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BotKlineModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BotKlineModelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BotKlineModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
