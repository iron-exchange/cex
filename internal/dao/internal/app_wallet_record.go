// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppWalletRecordDao is the data access object for the table t_app_wallet_record.
type AppWalletRecordDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AppWalletRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AppWalletRecordColumns defines and stores column names for the table t_app_wallet_record.
type AppWalletRecordColumns struct {
	Id             string // 卡ID
	Amount         string // 余额
	UAmount        string // 换算U金额
	CreateBy       string //
	CreateTime     string // 创建时间
	UpdateBy       string //
	UpdateTime     string // 更新时间
	Remark         string //
	UserId         string // 用户id
	SearchValue    string //
	BeforeAmount   string // 前值
	AfterAmount    string // 后值
	SerialId       string //
	Type           string // 余额
	Symbol         string // 币种
	AdminParentIds string // 代理ID
	OperateTime    string // 操作时间
}

// appWalletRecordColumns holds the columns for the table t_app_wallet_record.
var appWalletRecordColumns = AppWalletRecordColumns{
	Id:             "id",
	Amount:         "amount",
	UAmount:        "u_amount",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	UserId:         "user_id",
	SearchValue:    "search_value",
	BeforeAmount:   "before_amount",
	AfterAmount:    "after_amount",
	SerialId:       "serial_id",
	Type:           "type",
	Symbol:         "symbol",
	AdminParentIds: "admin_parent_ids",
	OperateTime:    "operate_time",
}

// NewAppWalletRecordDao creates and returns a new DAO object for table data access.
func NewAppWalletRecordDao(handlers ...gdb.ModelHandler) *AppWalletRecordDao {
	return &AppWalletRecordDao{
		group:    "default",
		table:    "t_app_wallet_record",
		columns:  appWalletRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppWalletRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppWalletRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppWalletRecordDao) Columns() AppWalletRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppWalletRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppWalletRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppWalletRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
