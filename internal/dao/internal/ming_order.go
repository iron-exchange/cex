// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MingOrderDao is the data access object for the table t_ming_order.
type MingOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MingOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MingOrderColumns defines and stores column names for the table t_ming_order.
type MingOrderColumns struct {
	Id              string //
	Amount          string // 投资金额（分）
	Days            string // 投资期限（天）
	Status          string // 0 收益  1 结算
	PlanId          string // 项目id
	PlanTitle       string // 项目名称
	OrderNo         string // 订单编号
	CreateTime      string // 投资时间
	EndTime         string // 到期时间
	SettleTime      string // 结算时间
	AccumulaEarn    string // 累计收益
	UpdateTime      string //
	MinOdds         string // 最小利率
	MaxOdds         string // 最大利率
	AdminUserIds    string // 后台用户id
	CollectionOrder string //
	UserId          string //
	OrderAmount     string //
}

// mingOrderColumns holds the columns for the table t_ming_order.
var mingOrderColumns = MingOrderColumns{
	Id:              "id",
	Amount:          "amount",
	Days:            "days",
	Status:          "status",
	PlanId:          "plan_id",
	PlanTitle:       "plan_title",
	OrderNo:         "order_no",
	CreateTime:      "create_time",
	EndTime:         "end_time",
	SettleTime:      "settle_time",
	AccumulaEarn:    "accumula_earn",
	UpdateTime:      "update_time",
	MinOdds:         "min_odds",
	MaxOdds:         "max_odds",
	AdminUserIds:    "admin_user_ids",
	CollectionOrder: "collection_order",
	UserId:          "user_id",
	OrderAmount:     "order_amount",
}

// NewMingOrderDao creates and returns a new DAO object for table data access.
func NewMingOrderDao(handlers ...gdb.ModelHandler) *MingOrderDao {
	return &MingOrderDao{
		group:    "default",
		table:    "t_ming_order",
		columns:  mingOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MingOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MingOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MingOrderDao) Columns() MingOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MingOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MingOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MingOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
