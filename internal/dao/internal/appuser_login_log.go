// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppuserLoginLogDao is the data access object for the table t_appuser_login_log.
type AppuserLoginLogDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AppuserLoginLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AppuserLoginLogColumns defines and stores column names for the table t_appuser_login_log.
type AppuserLoginLogColumns struct {
	Id            string // 主键ID
	UserId        string // 登录用户ID
	Username      string // 登录用户名
	Ipaddr        string // 访问IP
	LoginLocation string // 访问位置
	Browser       string // 浏览器
	Os            string // 系统OS
	Status        string // 登录状态（0成功 1失败）
	Msg           string //
	LoginTime     string // 访问时间
}

// appuserLoginLogColumns holds the columns for the table t_appuser_login_log.
var appuserLoginLogColumns = AppuserLoginLogColumns{
	Id:            "id",
	UserId:        "user_id",
	Username:      "username",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginTime:     "login_time",
}

// NewAppuserLoginLogDao creates and returns a new DAO object for table data access.
func NewAppuserLoginLogDao(handlers ...gdb.ModelHandler) *AppuserLoginLogDao {
	return &AppuserLoginLogDao{
		group:    "default",
		table:    "t_appuser_login_log",
		columns:  appuserLoginLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppuserLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppuserLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppuserLoginLogDao) Columns() AppuserLoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppuserLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppuserLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppuserLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
