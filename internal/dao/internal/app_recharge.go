// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppRechargeDao is the data access object for the table t_app_recharge.
type AppRechargeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppRechargeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppRechargeColumns defines and stores column names for the table t_app_recharge.
type AppRechargeColumns struct {
	Id             string // 卡ID
	CreateBy       string //
	CreateTime     string // 创建时间
	UpdateBy       string //
	UpdateTime     string // 更新时间
	Remark         string //
	UserId         string // 所有者ID
	Username       string // 用户名
	Amount         string // 充值金额
	Bonus          string //
	Status         string // 状态
	SerialId       string // 订单号
	TxId           string // 第三方支付订单号
	Type           string // 类型
	SearchValue    string //
	Address        string // 充值地址
	Tree           string //
	Coin           string // 币总
	ToAddress      string // 入款地址
	BlockTime      string // 区块时间
	Host           string //
	RealAmount     string // 实际到账金额
	FileName       string // 充值凭证
	RechargeRemark string //
	NoticeFlag     string // 通知字段 0未通知 1通知了
	AppParentIds   string // app代理ids
	AdminParentIds string // 后台代理ids
	OperateTime    string // 操作时间
	OrderType      string // 订单类型 1/null=充值  2=彩金赠送
}

// appRechargeColumns holds the columns for the table t_app_recharge.
var appRechargeColumns = AppRechargeColumns{
	Id:             "id",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
	UserId:         "user_id",
	Username:       "username",
	Amount:         "amount",
	Bonus:          "bonus",
	Status:         "status",
	SerialId:       "serial_id",
	TxId:           "tx_id",
	Type:           "type",
	SearchValue:    "search_value",
	Address:        "address",
	Tree:           "tree",
	Coin:           "coin",
	ToAddress:      "to_address",
	BlockTime:      "block_time",
	Host:           "host",
	RealAmount:     "real_amount",
	FileName:       "file_name",
	RechargeRemark: "recharge_remark",
	NoticeFlag:     "notice_flag",
	AppParentIds:   "app_parent_ids",
	AdminParentIds: "admin_parent_ids",
	OperateTime:    "operate_time",
	OrderType:      "order_type",
}

// NewAppRechargeDao creates and returns a new DAO object for table data access.
func NewAppRechargeDao(handlers ...gdb.ModelHandler) *AppRechargeDao {
	return &AppRechargeDao{
		group:    "default",
		table:    "t_app_recharge",
		columns:  appRechargeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppRechargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppRechargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppRechargeDao) Columns() AppRechargeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppRechargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppRechargeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppRechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
