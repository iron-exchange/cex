// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SecondPeriodConfigDao is the data access object for the table t_second_period_config.
type SecondPeriodConfigDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  SecondPeriodConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// SecondPeriodConfigColumns defines and stores column names for the table t_second_period_config.
type SecondPeriodConfigColumns struct {
	Id          string // id
	SecondId    string // 秒合约币种配置id
	Period      string // 时间周期  单位秒
	Odds        string // 赔率
	MaxAmount   string // 最大金额
	MinAmount   string // 最小金额
	Status      string // 1开启 2关闭
	CreateBy    string // 创建人
	CreateTime  string // 创建时间
	UpdateBy    string // 更新人
	UpdateTime  string // 更新时间
	Remark      string // 备注
	SearchValue string //
	Flag        string // 全输标识
}

// secondPeriodConfigColumns holds the columns for the table t_second_period_config.
var secondPeriodConfigColumns = SecondPeriodConfigColumns{
	Id:          "id",
	SecondId:    "second_id",
	Period:      "period",
	Odds:        "odds",
	MaxAmount:   "max_amount",
	MinAmount:   "min_amount",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	SearchValue: "search_value",
	Flag:        "flag",
}

// NewSecondPeriodConfigDao creates and returns a new DAO object for table data access.
func NewSecondPeriodConfigDao(handlers ...gdb.ModelHandler) *SecondPeriodConfigDao {
	return &SecondPeriodConfigDao{
		group:    "default",
		table:    "t_second_period_config",
		columns:  secondPeriodConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SecondPeriodConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SecondPeriodConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SecondPeriodConfigDao) Columns() SecondPeriodConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SecondPeriodConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SecondPeriodConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SecondPeriodConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
