// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SymbolManageDao is the data access object for the table t_symbol_manage.
type SymbolManageDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SymbolManageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SymbolManageColumns defines and stores column names for the table t_symbol_manage.
type SymbolManageColumns struct {
	Id           string // 主键id
	Symbol       string // 币种
	MinChargeNum string // 最小兑换数量
	MaxChargeNum string // 最大兑换数量
	Commission   string // 手续费(%)
	Sort         string // 排序
	Enable       string // 1 启用 2 禁用
	Logo         string // 图标
	Market       string // 交易所
	Remark       string // 备注
	CreateBy     string // 创建人
	CreateTime   string // 创建时间
	UpdateBy     string // 修改人
	UpdateTime   string // 修改时间
	DelFlag      string // 0正常  2删除
	SearchValue  string //
}

// symbolManageColumns holds the columns for the table t_symbol_manage.
var symbolManageColumns = SymbolManageColumns{
	Id:           "id",
	Symbol:       "symbol",
	MinChargeNum: "min_charge_num",
	MaxChargeNum: "max_charge_num",
	Commission:   "commission",
	Sort:         "sort",
	Enable:       "enable",
	Logo:         "logo",
	Market:       "market",
	Remark:       "remark",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
	DelFlag:      "del_flag",
	SearchValue:  "search_value",
}

// NewSymbolManageDao creates and returns a new DAO object for table data access.
func NewSymbolManageDao(handlers ...gdb.ModelHandler) *SymbolManageDao {
	return &SymbolManageDao{
		group:    "default",
		table:    "t_symbol_manage",
		columns:  symbolManageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SymbolManageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SymbolManageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SymbolManageDao) Columns() SymbolManageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SymbolManageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SymbolManageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SymbolManageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
