// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppAssetDao is the data access object for the table t_app_asset.
type AppAssetDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppAssetColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppAssetColumns defines and stores column names for the table t_app_asset.
type AppAssetColumns struct {
	Id                   string //
	UserId               string //
	Adress               string // 地址
	Symbol               string // 币种
	Amout                string // 资产总额
	OccupiedAmount       string // 占用资产
	AvailableAmount      string // 可用资产
	AvailableAmountDaily string // 每日余额（0点时分的余额，提现会减少）
	CodingVolumeDaily    string // 每日打码量（24点之前，提现会减少）
	Type                 string // 资产类型 1=平台资产 2=理财资产 3=合约账户
	CreateBy             string // 创建人
	CreateTime           string // 创建时间
	UpdateBy             string // 更新人
	UpdateTime           string // 更新时间
	Remark               string // 备注
	SearchValue          string //
}

// appAssetColumns holds the columns for the table t_app_asset.
var appAssetColumns = AppAssetColumns{
	Id:                   "id",
	UserId:               "user_id",
	Adress:               "adress",
	Symbol:               "symbol",
	Amout:                "amout",
	OccupiedAmount:       "occupied_amount",
	AvailableAmount:      "available_amount",
	AvailableAmountDaily: "available_amount_daily",
	CodingVolumeDaily:    "coding_volume_daily",
	Type:                 "type",
	CreateBy:             "create_by",
	CreateTime:           "create_time",
	UpdateBy:             "update_by",
	UpdateTime:           "update_time",
	Remark:               "remark",
	SearchValue:          "search_value",
}

// NewAppAssetDao creates and returns a new DAO object for table data access.
func NewAppAssetDao(handlers ...gdb.ModelHandler) *AppAssetDao {
	return &AppAssetDao{
		group:    "default",
		table:    "t_app_asset",
		columns:  appAssetColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppAssetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppAssetDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppAssetDao) Columns() AppAssetColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppAssetDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppAssetDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppAssetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
