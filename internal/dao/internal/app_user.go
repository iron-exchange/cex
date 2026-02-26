// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserDao is the data access object for the table t_app_user.
type AppUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppUserColumns defines and stores column names for the table t_app_user.
type AppUserColumns struct {
	UserId         string //
	LoginName      string // 姓名
	LoginPassword  string // 登陆密码
	Phone          string // 手机号
	IsTest         string // 0-正常 1-测试
	Address        string // 地址
	WalletType     string // 地址类型 ETH TRC
	Status         string // 0正常1冻结
	TotleAmont     string // 总打码量
	RechargeAmont  string // 充值打码量
	Buff           string // 0正常 1包赢 2包输
	AppParentIds   string // app代理ids
	AdminParentIds string // 后台代理ids
	ActiveCode     string // 邀请码
	RegisterIp     string // 注册ip
	Host           string // 注册域名
	Email          string // 邮箱
	Level          string // vip等级
	IsFreeze       string // 是否冻结  1=正常 2=冻结
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	UpdateBy       string // 更新人
	UpdateTime     string // 更新时间
	Remark         string // 备注
	SearchValue    string //
	IsBlack        string // 黑名单 1=正常 2拉黑
	BinanceEmail   string // 币安子账号邮箱
}

// appUserColumns holds the columns for the table t_app_user.
var appUserColumns = AppUserColumns{
	UserId:         "user_id",
	LoginName:      "login_name",
	LoginPassword:  "login_password",
	Phone:          "phone",
	IsTest:         "is_test",
	Address:        "address",
	WalletType:     "wallet_type",
	Status:         "status",
	TotleAmont:     "totle_amont",
	RechargeAmont:  "recharge_amont",
	Buff:           "buff",
	AppParentIds:   "app_parent_ids",
	AdminParentIds: "admin_parent_ids",
	ActiveCode:     "active_code",
	RegisterIp:     "register_ip",
	Host:           "host",
	Email:          "email",
	Level:          "level",
	IsFreeze:       "is_freeze",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	SearchValue:    "search_value",
	IsBlack:        "is_black",
	BinanceEmail:   "binance_email",
}

// NewAppUserDao creates and returns a new DAO object for table data access.
func NewAppUserDao(handlers ...gdb.ModelHandler) *AppUserDao {
	return &AppUserDao{
		group:    "default",
		table:    "t_app_user",
		columns:  appUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppUserDao) Columns() AppUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
