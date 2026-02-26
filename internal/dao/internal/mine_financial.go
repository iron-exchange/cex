// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MineFinancialDao is the data access object for the table t_mine_financial.
type MineFinancialDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  MineFinancialColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// MineFinancialColumns defines and stores column names for the table t_mine_financial.
type MineFinancialColumns struct {
	Id                  string //
	Title               string // 标题
	Icon                string // 图标
	Status              string // 启用禁用(展示在前端)1开0关
	Days                string // 天数(如 7,10,30)
	DefaultOdds         string // 违约利率
	MinOdds             string // 最小日利率百分比
	MaxOdds             string // 最大日利率百分比
	TimeLimit           string // 每人限购次数，0表示不限
	LimitMin            string // 最小金额
	LimitMax            string // 最大金额
	IsHot               string // 是否热销1是0否
	Sort                string // 排序
	CreateBy            string // 创建人
	CreateTime          string // 创建时间
	UpdateBy            string // 更新人员
	UpdateTime          string // 更新时间
	BuyPurchase         string // 购买次数
	AvgRate             string // 日平均利率
	Coin                string // 币种
	Classify            string // 分类（0 普通  1 vip  2 增值）
	BasicInvestAmount   string // 平台基础投资金额
	TotalInvestAmount   string // 平台总投资额
	Level               string // VIP等级
	Process             string // 项目进度
	RemainAmount        string // 剩余金额
	Remark              string // 标签
	PurchasedAmount     string // 易购金额
	Problem             string // 常见问题
	ProdectIntroduction string // 产品介绍
}

// mineFinancialColumns holds the columns for the table t_mine_financial.
var mineFinancialColumns = MineFinancialColumns{
	Id:                  "id",
	Title:               "title",
	Icon:                "icon",
	Status:              "status",
	Days:                "days",
	DefaultOdds:         "default_odds",
	MinOdds:             "min_odds",
	MaxOdds:             "max_odds",
	TimeLimit:           "time_limit",
	LimitMin:            "limit_min",
	LimitMax:            "limit_max",
	IsHot:               "is_hot",
	Sort:                "sort",
	CreateBy:            "create_by",
	CreateTime:          "create_time",
	UpdateBy:            "update_by",
	UpdateTime:          "update_time",
	BuyPurchase:         "buy_purchase",
	AvgRate:             "avg_rate",
	Coin:                "coin",
	Classify:            "classify",
	BasicInvestAmount:   "basic_invest_amount",
	TotalInvestAmount:   "total_invest_amount",
	Level:               "level",
	Process:             "process",
	RemainAmount:        "remain_amount",
	Remark:              "remark",
	PurchasedAmount:     "purchased_amount",
	Problem:             "problem",
	ProdectIntroduction: "prodect_introduction",
}

// NewMineFinancialDao creates and returns a new DAO object for table data access.
func NewMineFinancialDao(handlers ...gdb.ModelHandler) *MineFinancialDao {
	return &MineFinancialDao{
		group:    "default",
		table:    "t_mine_financial",
		columns:  mineFinancialColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MineFinancialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MineFinancialDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MineFinancialDao) Columns() MineFinancialColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MineFinancialDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MineFinancialDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MineFinancialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
