// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MineOrderDayDao is the data access object for the table t_mine_order_day.
type MineOrderDayDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  MineOrderDayColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// MineOrderDayColumns defines and stores column names for the table t_mine_order_day.
type MineOrderDayColumns struct {
	Id          string //
	Amount      string // 投资金额（分）
	Odds        string // 当日利率
	Earn        string // 收益
	PlanId      string // 项目id
	OrderNo     string // 订单编号
	CreateTime  string // 时间
	Address     string // 地址
	Type        string // 0 质押挖矿 1 非质押挖矿
	UpdateTime  string //
	Status      string // 1 待结算  2  结算
	SearchValue string //
	UpdateBy    string //
	CreateBy    string //
	Remark      string //
}

// mineOrderDayColumns holds the columns for the table t_mine_order_day.
var mineOrderDayColumns = MineOrderDayColumns{
	Id:          "id",
	Amount:      "amount",
	Odds:        "odds",
	Earn:        "earn",
	PlanId:      "plan_id",
	OrderNo:     "order_no",
	CreateTime:  "create_time",
	Address:     "address",
	Type:        "type",
	UpdateTime:  "update_time",
	Status:      "status",
	SearchValue: "search_value",
	UpdateBy:    "update_by",
	CreateBy:    "create_by",
	Remark:      "remark",
}

// NewMineOrderDayDao creates and returns a new DAO object for table data access.
func NewMineOrderDayDao(handlers ...gdb.ModelHandler) *MineOrderDayDao {
	return &MineOrderDayDao{
		group:    "default",
		table:    "t_mine_order_day",
		columns:  mineOrderDayColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MineOrderDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MineOrderDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MineOrderDayDao) Columns() MineOrderDayColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MineOrderDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MineOrderDayDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MineOrderDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
