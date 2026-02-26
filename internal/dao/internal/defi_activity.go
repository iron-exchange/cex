// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DefiActivityDao is the data access object for the table t_defi_activity.
type DefiActivityDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  DefiActivityColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// DefiActivityColumns defines and stores column names for the table t_defi_activity.
type DefiActivityColumns struct {
	Id          string // id
	TotleAmount string // 需要金额
	UserId      string // 用户id
	EndTime     string // 结束时间
	Amount      string // 奖励金额
	Type        string // 0-usdt 1-eth
	Status      string // 0未领取 1已读 2已领取
	CreateBy    string // 创建者
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Remark      string // 备注
	SearchValue string //
}

// defiActivityColumns holds the columns for the table t_defi_activity.
var defiActivityColumns = DefiActivityColumns{
	Id:          "id",
	TotleAmount: "totle_amount",
	UserId:      "user_id",
	EndTime:     "end_time",
	Amount:      "amount",
	Type:        "type",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	SearchValue: "search_value",
}

// NewDefiActivityDao creates and returns a new DAO object for table data access.
func NewDefiActivityDao(handlers ...gdb.ModelHandler) *DefiActivityDao {
	return &DefiActivityDao{
		group:    "default",
		table:    "t_defi_activity",
		columns:  defiActivityColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DefiActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DefiActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DefiActivityDao) Columns() DefiActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DefiActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DefiActivityDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DefiActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
