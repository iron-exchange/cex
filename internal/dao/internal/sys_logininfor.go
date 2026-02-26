// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLogininforDao is the data access object for the table sys_logininfor.
type SysLogininforDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysLogininforColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysLogininforColumns defines and stores column names for the table sys_logininfor.
type SysLogininforColumns struct {
	InfoId        string // 访问ID
	UserName      string // 用户账号
	Ipaddr        string // 登录IP地址
	LoginLocation string // 登录地点
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Status        string // 登录状态（0成功 1失败）
	Msg           string // 提示消息
	LoginTime     string // 访问时间
}

// sysLogininforColumns holds the columns for the table sys_logininfor.
var sysLogininforColumns = SysLogininforColumns{
	InfoId:        "info_id",
	UserName:      "user_name",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginTime:     "login_time",
}

// NewSysLogininforDao creates and returns a new DAO object for table data access.
func NewSysLogininforDao(handlers ...gdb.ModelHandler) *SysLogininforDao {
	return &SysLogininforDao{
		group:    "default",
		table:    "sys_logininfor",
		columns:  sysLogininforColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLogininforDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLogininforDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLogininforDao) Columns() SysLogininforColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLogininforDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLogininforDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysLogininforDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
