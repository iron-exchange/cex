// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysNoticeDao is the data access object for the table sys_notice.
type SysNoticeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysNoticeColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysNoticeColumns defines and stores column names for the table sys_notice.
type SysNoticeColumns struct {
	NoticeId      string // 公告ID
	NoticeTitle   string // 公告标题
	NoticeType    string // 公告类型（1通知 2公告）
	NoticeContent string // 公告内容
	Status        string // 公告状态（0正常 1关闭）
	CreateBy      string // 创建者
	CreateTime    string // 创建时间
	UpdateBy      string // 更新者
	UpdateTime    string // 更新时间
	Remark        string // 备注
}

// sysNoticeColumns holds the columns for the table sys_notice.
var sysNoticeColumns = SysNoticeColumns{
	NoticeId:      "notice_id",
	NoticeTitle:   "notice_title",
	NoticeType:    "notice_type",
	NoticeContent: "notice_content",
	Status:        "status",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
	Remark:        "remark",
}

// NewSysNoticeDao creates and returns a new DAO object for table data access.
func NewSysNoticeDao(handlers ...gdb.ModelHandler) *SysNoticeDao {
	return &SysNoticeDao{
		group:    "default",
		table:    "sys_notice",
		columns:  sysNoticeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysNoticeDao) Columns() SysNoticeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysNoticeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
