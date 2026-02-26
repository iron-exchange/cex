// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	UserId      string // 用户ID
	DeptId      string // 部门ID
	UserName    string // 用户账号
	NickName    string // 用户昵称
	UserType    string // 用户类型（0普通用户 1 组长 2 代理）
	Email       string // 用户邮箱
	Phonenumber string // 手机号码
	Sex         string // 用户性别（0男 1女 2未知）
	Avatar      string // 头像地址
	Password    string // 密码
	Status      string // 帐号状态（0正常 1停用）
	GoogleKey   string // 谷歌验证码key
	DelFlag     string // 删除标志（0代表存在 2代表删除）
	LoginIp     string // 最后登录IP
	LoginDate   string // 最后登录时间
	CreateBy    string // 创建者
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Remark      string // 备注
	ParentId    string // 组长ID
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	UserId:      "user_id",
	DeptId:      "dept_id",
	UserName:    "user_name",
	NickName:    "nick_name",
	UserType:    "user_type",
	Email:       "email",
	Phonenumber: "phonenumber",
	Sex:         "sex",
	Avatar:      "avatar",
	Password:    "password",
	Status:      "status",
	GoogleKey:   "google_key",
	DelFlag:     "del_flag",
	LoginIp:     "login_ip",
	LoginDate:   "login_date",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
	ParentId:    "parent_id",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(handlers ...gdb.ModelHandler) *SysUserDao {
	return &SysUserDao{
		group:    "default",
		table:    "sys_user",
		columns:  sysUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
