// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeDao is the data access object for the table t_notice.
type NoticeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  NoticeColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// NoticeColumns defines and stores column names for the table t_notice.
type NoticeColumns struct {
	NoticeId      string // 公告ID
	NoticeTitle   string // 标题
	NoticeType    string // 公告类型 1=公告信息 2=活动公告 3=首页滚动公告
	ModelType     string // 模块类型 1=公告信息 2=活动公告 3=首页滚动公告   1={1=链接弹窗 2=图文弹窗},   2={1=首页轮播活动 2=Defi挖矿活动图},  3={1=首页滚动公告}  注:没有二级的默认给1  二级联动
	NoticeContent string // 内容
	CommentsNum   string // 评论数
	Cover         string // 图片
	ViewNum       string // 浏览数
	ExpireTime    string // 公告截止时间
	ImgUrl        string // 图片地址
	ChainedUrl    string // 链接地址
	DetailUrl     string // 详情页
	LanguageId    string // zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13
	Status        string // 公告状态（0正常 1关闭）
	Source        string // 展示端1=pc 2=h5
	Sort          string // 排序
	Remark        string // 备注
	CreateBy      string // 创建人
	CreateTime    string // 创建时间
	UpdateBy      string // 修改人
	UpdateTime    string // 更新时间
}

// noticeColumns holds the columns for the table t_notice.
var noticeColumns = NoticeColumns{
	NoticeId:      "notice_id",
	NoticeTitle:   "notice_title",
	NoticeType:    "notice_type",
	ModelType:     "model_type",
	NoticeContent: "notice_content",
	CommentsNum:   "comments_num",
	Cover:         "cover",
	ViewNum:       "view_num",
	ExpireTime:    "expire_time",
	ImgUrl:        "img_url",
	ChainedUrl:    "chained_url",
	DetailUrl:     "detail_url",
	LanguageId:    "language_id",
	Status:        "status",
	Source:        "source",
	Sort:          "sort",
	Remark:        "remark",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
}

// NewNoticeDao creates and returns a new DAO object for table data access.
func NewNoticeDao(handlers ...gdb.ModelHandler) *NoticeDao {
	return &NoticeDao{
		group:    "default",
		table:    "t_notice",
		columns:  noticeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NoticeDao) Columns() NoticeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NoticeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *NoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
