// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OptionRulesDao is the data access object for the table t_option_rules.
type OptionRulesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OptionRulesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OptionRulesColumns defines and stores column names for the table t_option_rules.
type OptionRulesColumns struct {
	Id          string //
	Title       string // 标题
	Language    string // 语言 en zh
	Content     string // 内容
	IsShow      string // 是否展示 0展示  2不展示
	CreateTime  string // 创建时间
	Type        string // 0=服务条款 1=秒合约说明 2=币币交易说明 3=代理活动 4=U本位合约说明 5=注册隐私政策 6=注册使用条款 7=贷款规则
	SearchValue string //
}

// optionRulesColumns holds the columns for the table t_option_rules.
var optionRulesColumns = OptionRulesColumns{
	Id:          "id",
	Title:       "title",
	Language:    "language",
	Content:     "content",
	IsShow:      "is_show",
	CreateTime:  "create_time",
	Type:        "type",
	SearchValue: "search_value",
}

// NewOptionRulesDao creates and returns a new DAO object for table data access.
func NewOptionRulesDao(handlers ...gdb.ModelHandler) *OptionRulesDao {
	return &OptionRulesDao{
		group:    "default",
		table:    "t_option_rules",
		columns:  optionRulesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OptionRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OptionRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OptionRulesDao) Columns() OptionRulesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OptionRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OptionRulesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OptionRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
