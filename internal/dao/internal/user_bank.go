// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBankDao is the data access object for the table t_user_bank.
type UserBankDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserBankColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserBankColumns defines and stores column names for the table t_user_bank.
type UserBankColumns struct {
	Id             string //
	UserName       string // 姓名
	CardNumber     string // 银行卡号
	BankName       string // 开户银行名称
	BankAddress    string // 开户省市
	BankBranch     string // 开户网点
	UserId         string // 用户名称
	AdminParentIds string //
	CreateBy       string //
	CreateTime     string // 创建时间
	UpdateBy       string //
	UpdateTime     string // 更新时间
	Remark         string //
	SearchValue    string //
	BankCode       string // 银行编码
	UserAddress    string // 用户地址
}

// userBankColumns holds the columns for the table t_user_bank.
var userBankColumns = UserBankColumns{
	Id:             "id",
	UserName:       "user_name",
	CardNumber:     "card_number",
	BankName:       "bank_name",
	BankAddress:    "bank_address",
	BankBranch:     "bank_branch",
	UserId:         "user_id",
	AdminParentIds: "admin_parent_ids",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	SearchValue:    "search_value",
	BankCode:       "bank_code",
	UserAddress:    "user_address",
}

// NewUserBankDao creates and returns a new DAO object for table data access.
func NewUserBankDao(handlers ...gdb.ModelHandler) *UserBankDao {
	return &UserBankDao{
		group:    "default",
		table:    "t_user_bank",
		columns:  userBankColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserBankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserBankDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserBankDao) Columns() UserBankColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserBankDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserBankDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserBankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
