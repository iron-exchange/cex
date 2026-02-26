// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MingProductUserDao is the data access object for the table t_ming_product_user.
type MingProductUserDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  MingProductUserColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// MingProductUserColumns defines and stores column names for the table t_ming_product_user.
type MingProductUserColumns struct {
	Id          string //
	ProductId   string // 产品id
	AppUserId   string // 玩家用户id
	PledgeNum   string // 限购次数
	CreateBy    string // 创建人
	CreateTime  string // 创建时间
	UpdateBy    string // 更新人员
	UpdateTime  string // 更新时间
	SearchValue string // 币种
	Remark      string // 标签
}

// mingProductUserColumns holds the columns for the table t_ming_product_user.
var mingProductUserColumns = MingProductUserColumns{
	Id:          "id",
	ProductId:   "product_id",
	AppUserId:   "app_user_id",
	PledgeNum:   "pledge_num",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	SearchValue: "search_value",
	Remark:      "remark",
}

// NewMingProductUserDao creates and returns a new DAO object for table data access.
func NewMingProductUserDao(handlers ...gdb.ModelHandler) *MingProductUserDao {
	return &MingProductUserDao{
		group:    "default",
		table:    "t_ming_product_user",
		columns:  mingProductUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MingProductUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MingProductUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MingProductUserDao) Columns() MingProductUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MingProductUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MingProductUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MingProductUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
