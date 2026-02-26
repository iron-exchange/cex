// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomeSetterDao is the data access object for the table t_home_setter.
type HomeSetterDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HomeSetterColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HomeSetterColumns defines and stores column names for the table t_home_setter.
type HomeSetterColumns struct {
	Id           string //
	Title        string // 标题
	Author       string // 作者
	Content      string // 内容
	CreateTime   string // 创建时间
	ImgUrl       string // 图片地址
	Sort         string // 排序
	IsShow       string // 是否展示 0展示  2不展示
	LanguageName string // 语言
	LikesNum     string // 点赞数
	HomeType     string // 类型（0 首页文本  1 问题列表）
	ModelType    string // 功能（0=首页  1=defi挖矿 2=助力贷）
	SearchValue  string //
}

// homeSetterColumns holds the columns for the table t_home_setter.
var homeSetterColumns = HomeSetterColumns{
	Id:           "id",
	Title:        "title",
	Author:       "author",
	Content:      "content",
	CreateTime:   "create_time",
	ImgUrl:       "img_url",
	Sort:         "sort",
	IsShow:       "is_show",
	LanguageName: "language_name",
	LikesNum:     "likes_num",
	HomeType:     "home_type",
	ModelType:    "model_type",
	SearchValue:  "search_value",
}

// NewHomeSetterDao creates and returns a new DAO object for table data access.
func NewHomeSetterDao(handlers ...gdb.ModelHandler) *HomeSetterDao {
	return &HomeSetterDao{
		group:    "default",
		table:    "t_home_setter",
		columns:  homeSetterColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HomeSetterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HomeSetterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HomeSetterDao) Columns() HomeSetterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HomeSetterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HomeSetterDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HomeSetterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
