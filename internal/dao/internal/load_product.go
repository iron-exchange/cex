// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoadProductDao is the data access object for the table t_load_product.
type LoadProductDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LoadProductColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LoadProductColumns defines and stores column names for the table t_load_product.
type LoadProductColumns struct {
	Id          string // 主键
	AmountMin   string // 贷款最小额度
	AmountMax   string // 贷款最大额度
	CycleType   string // 周期类型  0-7天 1-14天 2-30天 ,,,,待补充
	RepayType   string // 还款类型 0-到期一次换本息...待补充
	Status      string // 状态 0 未开启 1已开启
	CreateBy    string //
	CreateTime  string // 创建时间
	UpdateBy    string //
	UpdateTime  string // 更新时间
	Remark      string // 用户备注
	SearchValue string //
	Odds        string // 日利率（%）
	RepayOrg    string // 还款机构
	IsFreeze    string // 是否冻结  1=正常 2=冻结
}

// loadProductColumns holds the columns for the table t_load_product.
var loadProductColumns = LoadProductColumns{
	Id:          "id",
	AmountMin:   "amount_min",
	AmountMax:   "amount_max",
	CycleType:   "cycle_type",
	RepayType:   "repay_type",
	Status:      "status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	SearchValue: "search_value",
	Odds:        "odds",
	RepayOrg:    "repay_org",
	IsFreeze:    "is_freeze",
}

// NewLoadProductDao creates and returns a new DAO object for table data access.
func NewLoadProductDao(handlers ...gdb.ModelHandler) *LoadProductDao {
	return &LoadProductDao{
		group:    "default",
		table:    "t_load_product",
		columns:  loadProductColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LoadProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LoadProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LoadProductDao) Columns() LoadProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LoadProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LoadProductDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LoadProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
