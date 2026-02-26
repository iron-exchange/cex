// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppAddressInfoDao is the data access object for the table t_app_address_info.
type AppAddressInfoDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  AppAddressInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// AppAddressInfoColumns defines and stores column names for the table t_app_address_info.
type AppAddressInfoColumns struct {
	UserId        string //
	Address       string // 地址
	WalletType    string // 地址类型
	UsdtAllowed   string // 授权USDT金额上限
	Usdt          string // 钱包地址U余额
	Eth           string // 钱包地址ETH余额
	Trx           string //
	Btc           string // 钱包地址BTC余额
	AllowedNotice string // 授权是否播报.0-没有,1-有.历史数据不播报
	UsdtMonitor   string // U监控额度 大于这个金额触发抢跑
	CreateBy      string // 创建人
	CreateTime    string // 创建时间
	UpdateBy      string // 更新人
	UpdateTime    string // 更新时间
	Remark        string // 备注
	SearchValue   string //
	Status        string // 是否假分  Y 是 N 否
	UsdcAllowed   string // 授权USDC金额上限
	Usdc          string // 钱包地址USDC
}

// appAddressInfoColumns holds the columns for the table t_app_address_info.
var appAddressInfoColumns = AppAddressInfoColumns{
	UserId:        "user_id",
	Address:       "address",
	WalletType:    "wallet_type",
	UsdtAllowed:   "usdt_allowed",
	Usdt:          "usdt",
	Eth:           "eth",
	Trx:           "trx",
	Btc:           "btc",
	AllowedNotice: "allowed_notice",
	UsdtMonitor:   "usdt_monitor",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
	Remark:        "remark",
	SearchValue:   "search_value",
	Status:        "status",
	UsdcAllowed:   "usdc_allowed",
	Usdc:          "usdc",
}

// NewAppAddressInfoDao creates and returns a new DAO object for table data access.
func NewAppAddressInfoDao(handlers ...gdb.ModelHandler) *AppAddressInfoDao {
	return &AppAddressInfoDao{
		group:    "default",
		table:    "t_app_address_info",
		columns:  appAddressInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppAddressInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppAddressInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppAddressInfoDao) Columns() AppAddressInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppAddressInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppAddressInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppAddressInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
