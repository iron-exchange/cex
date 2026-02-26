// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppMailDao is the data access object for the table t_app_mail.
type AppMailDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppMailColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppMailColumns defines and stores column names for the table t_app_mail.
type AppMailColumns struct {
	Id          string //
	UserId      string //
	Title       string // 标题
	Content     string // 内容
	Type        string // 消息类型 1=普通消息 2=全站消息
	Status      string // 状态（0 未读 1已读）
	OpertorId   string // 操作人
	CreateTime  string //
	UpdateTime  string //
	SearchValue string //
	DelFlag     string // 0正常 2删除
}

// appMailColumns holds the columns for the table t_app_mail.
var appMailColumns = AppMailColumns{
	Id:          "id",
	UserId:      "user_id",
	Title:       "title",
	Content:     "content",
	Type:        "type",
	Status:      "status",
	OpertorId:   "opertor_id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	SearchValue: "search_value",
	DelFlag:     "del_flag",
}

// NewAppMailDao creates and returns a new DAO object for table data access.
func NewAppMailDao(handlers ...gdb.ModelHandler) *AppMailDao {
	return &AppMailDao{
		group:    "default",
		table:    "t_app_mail",
		columns:  appMailColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppMailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppMailDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppMailDao) Columns() AppMailColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppMailDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppMailDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppMailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
