// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityDao is the data access object for the table t_activity.
type ActivityDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ActivityColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ActivityColumns defines and stores column names for the table t_activity.
type ActivityColumns struct {
	Id           string // 卡ID
	CreateBy     string //
	CreateTime   string // 创建时间
	UpdateBy     string //
	UpdateTime   string // 更新时间
	SearchValue  string //
	Status       string // 状态0编辑1上线
	Remark       string //
	DisplayUrl   string //
	DetailUrl    string // 详情页
	DetailStatus string //
	LoadingUrl   string //
	Tags         string //
	Type         string // 1广告2活动
	Source       string // 展示端,1:PC,2:H5
	LanguageId   string // zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13
	Agent        string // 0代理无关，1相关
	JumpType     string // 跳转类型 0 内链 1外链
	SortNum      string // 排序
}

// activityColumns holds the columns for the table t_activity.
var activityColumns = ActivityColumns{
	Id:           "id",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
	SearchValue:  "search_value",
	Status:       "status",
	Remark:       "remark",
	DisplayUrl:   "display_url",
	DetailUrl:    "detail_url",
	DetailStatus: "detail_status",
	LoadingUrl:   "loading_url",
	Tags:         "tags",
	Type:         "type",
	Source:       "source",
	LanguageId:   "language_id",
	Agent:        "agent",
	JumpType:     "jump_type",
	SortNum:      "sort_num",
}

// NewActivityDao creates and returns a new DAO object for table data access.
func NewActivityDao(handlers ...gdb.ModelHandler) *ActivityDao {
	return &ActivityDao{
		group:    "default",
		table:    "t_activity",
		columns:  activityColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ActivityDao) Columns() ActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ActivityDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
