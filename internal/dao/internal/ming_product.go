// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MingProductDao is the data access object for the table t_ming_product.
type MingProductDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MingProductColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MingProductColumns defines and stores column names for the table t_ming_product.
type MingProductColumns struct {
	Id          string //
	Title       string // 标题
	Icon        string // 图标
	Status      string // 启用禁用(展示在前端)1开0关
	Days        string // 天数(如 7,10,30)
	DefaultOdds string // 违约利率
	MinOdds     string // 最小日利率百分比
	MaxOdds     string // 最大日利率百分比
	TimeLimit   string // 每人限购次数，0表示不限
	LimitMin    string // 最小金额
	LimitMax    string // 最大金额
	Sort        string // 排序
	CreateBy    string // 创建人
	CreateTime  string // 创建时间
	UpdateBy    string // 更新人员
	UpdateTime  string // 更新时间
	BuyPurchase string // 购买次数
	Coin        string // 币种
	Remark      string // 标签
}

// mingProductColumns holds the columns for the table t_ming_product.
var mingProductColumns = MingProductColumns{
	Id:          "id",
	Title:       "title",
	Icon:        "icon",
	Status:      "status",
	Days:        "days",
	DefaultOdds: "default_odds",
	MinOdds:     "min_odds",
	MaxOdds:     "max_odds",
	TimeLimit:   "time_limit",
	LimitMin:    "limit_min",
	LimitMax:    "limit_max",
	Sort:        "sort",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	BuyPurchase: "buy_purchase",
	Coin:        "coin",
	Remark:      "remark",
}

// NewMingProductDao creates and returns a new DAO object for table data access.
func NewMingProductDao(handlers ...gdb.ModelHandler) *MingProductDao {
	return &MingProductDao{
		group:    "default",
		table:    "t_ming_product",
		columns:  mingProductColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MingProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MingProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MingProductDao) Columns() MingProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MingProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MingProductDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MingProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
