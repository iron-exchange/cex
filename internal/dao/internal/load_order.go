// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoadOrderDao is the data access object for the table t_load_order.
type LoadOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LoadOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LoadOrderColumns defines and stores column names for the table t_load_order.
type LoadOrderColumns struct {
	Id             string // 主键
	ProId          string // 贷款商品表id
	UserId         string // 用户id
	CreateTime     string // 购买时间
	Amount         string // 贷款金额
	Rate           string // 贷款利率
	Interest       string // 利息
	Status         string // 0=待审核 1=审核通过  2=审核拒绝  3=已结清  4=已逾期
	FinalRepayTime string // 最后还款日
	DisburseTime   string // 放款日期
	ReturnTime     string // 还款日期
	DisburseAmount string // 审批金额
	AdminParentIds string // 后台代理ids
	CardUrl        string // 手持身份证
	CardBackUrl    string // 身份证正面
	CapitalUrl     string // 身份证反面
	LicenseUrl     string //
	OrderNo        string //
	CycleType      string // 还款周期  多少天
	Remark         string // 用户备注
	CreateBy       string //
	UpdateBy       string //
	UpdateTime     string // 更新时间
	SearchValue    string //
}

// loadOrderColumns holds the columns for the table t_load_order.
var loadOrderColumns = LoadOrderColumns{
	Id:             "id",
	ProId:          "pro_id",
	UserId:         "user_id",
	CreateTime:     "create_time",
	Amount:         "amount",
	Rate:           "rate",
	Interest:       "interest",
	Status:         "status",
	FinalRepayTime: "final_repay_time",
	DisburseTime:   "disburse_time",
	ReturnTime:     "return_time",
	DisburseAmount: "disburse_amount",
	AdminParentIds: "admin_parent_ids",
	CardUrl:        "card_url",
	CardBackUrl:    "card_back_url",
	CapitalUrl:     "capital_url",
	LicenseUrl:     "license_url",
	OrderNo:        "order_no",
	CycleType:      "cycle_type",
	Remark:         "remark",
	CreateBy:       "create_by",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	SearchValue:    "search_value",
}

// NewLoadOrderDao creates and returns a new DAO object for table data access.
func NewLoadOrderDao(handlers ...gdb.ModelHandler) *LoadOrderDao {
	return &LoadOrderDao{
		group:    "default",
		table:    "t_load_order",
		columns:  loadOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LoadOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LoadOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LoadOrderDao) Columns() LoadOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LoadOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LoadOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LoadOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
