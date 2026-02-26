// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OwnCoinDao is the data access object for the table t_own_coin.
type OwnCoinDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OwnCoinColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OwnCoinColumns defines and stores column names for the table t_own_coin.
type OwnCoinColumns struct {
	Id              string // 主键ID
	Coin            string // 币种
	Logo            string // 图标
	ReferCoin       string // 参考币种
	ReferMarket     string // 参考币种交易所
	ShowSymbol      string // 展示名称
	Price           string // 初始价格（单位USDT）
	Proportion      string // 价格百分比
	RaisingAmount   string // 私募发行量
	RaisedAmount    string // 已筹集额度
	PurchaseLimit   string // 预购上限
	TotalAmount     string // 总发行量
	ParticipantsNum string // 参与人数
	RaisingTime     string // 筹集期限
	BeginTime       string // 开始时间
	EndTime         string // 结束时间
	Introduce       string // 介绍
	Status          string // 1.未发布  2.筹集中 3 筹集成功 4.筹集失败
	CreateBy        string // 创建人
	CreateTime      string // 创建时间
	UpdateBy        string // 更新者
	UpdateTime      string // 更新时间
	Remark          string // 备注
}

// ownCoinColumns holds the columns for the table t_own_coin.
var ownCoinColumns = OwnCoinColumns{
	Id:              "id",
	Coin:            "coin",
	Logo:            "logo",
	ReferCoin:       "refer_coin",
	ReferMarket:     "refer_market",
	ShowSymbol:      "show_symbol",
	Price:           "price",
	Proportion:      "proportion",
	RaisingAmount:   "raising_amount",
	RaisedAmount:    "raised_amount",
	PurchaseLimit:   "purchase_limit",
	TotalAmount:     "total_amount",
	ParticipantsNum: "participants_num",
	RaisingTime:     "raising_time",
	BeginTime:       "begin_time",
	EndTime:         "end_time",
	Introduce:       "introduce",
	Status:          "status",
	CreateBy:        "create_by",
	CreateTime:      "create_time",
	UpdateBy:        "update_by",
	UpdateTime:      "update_time",
	Remark:          "remark",
}

// NewOwnCoinDao creates and returns a new DAO object for table data access.
func NewOwnCoinDao(handlers ...gdb.ModelHandler) *OwnCoinDao {
	return &OwnCoinDao{
		group:    "default",
		table:    "t_own_coin",
		columns:  ownCoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OwnCoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OwnCoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OwnCoinDao) Columns() OwnCoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OwnCoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OwnCoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OwnCoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
