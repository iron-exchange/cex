// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WithdrawDao is the data access object for the table t_withdraw.
type WithdrawDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WithdrawColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WithdrawColumns defines and stores column names for the table t_withdraw.
type WithdrawColumns struct {
	Id                string // 卡ID
	CreateBy          string //
	CreateTime        string // 创建时间
	UpdateBy          string //
	UpdateTime        string // 更新时间
	Remark            string //
	UserId            string // 用户
	Username          string // 用户名
	Address           string // 用户名
	Amount            string // 提现金额
	Status            string // 0审核中1成功2失败
	SerialId          string //
	SearchValue       string //
	FromAddr          string // 用户名
	Type              string // 0审核中1成功2失败
	Coin              string // 用户名
	Ratio             string //
	Fee               string // 手续费
	WithdrawId        string // 用户名
	Host              string // Host
	RealAmount        string // 实际金额
	ToAdress          string // 收款地址
	AdminParentIds    string // 后台用户id
	NoticeFlag        string // 通知字段 0未通知 1通知了
	WithDrawRemark    string // 提现说明
	BankName          string // 银行名称
	BankUserName      string // 银行收款人名称
	BankBranch        string //
	AdminUserIds      string //
	OperateTime       string // 操作时间
	FixedFee          string // 固定手续费
	OrderType         string // 订单类型 1/null 提现  2=彩金扣减
	ExchangeRate      string // 汇率
	ReceiptAmount     string // 应到账金额
	ReceiptRealAmount string // 实际到账金额
	ReceiptCoin       string // 到账币种
}

// withdrawColumns holds the columns for the table t_withdraw.
var withdrawColumns = WithdrawColumns{
	Id:                "id",
	CreateBy:          "create_by",
	CreateTime:        "create_time",
	UpdateBy:          "update_by",
	UpdateTime:        "update_time",
	Remark:            "remark",
	UserId:            "user_id",
	Username:          "username",
	Address:           "address",
	Amount:            "amount",
	Status:            "status",
	SerialId:          "serial_id",
	SearchValue:       "search_value",
	FromAddr:          "from_addr",
	Type:              "type",
	Coin:              "coin",
	Ratio:             "ratio",
	Fee:               "fee",
	WithdrawId:        "withdraw_id",
	Host:              "host",
	RealAmount:        "real_amount",
	ToAdress:          "to_adress",
	AdminParentIds:    "admin_parent_ids",
	NoticeFlag:        "notice_flag",
	WithDrawRemark:    "with_draw_remark",
	BankName:          "bank_name",
	BankUserName:      "bank_user_name",
	BankBranch:        "bank_branch",
	AdminUserIds:      "admin_user_ids",
	OperateTime:       "operate_time",
	FixedFee:          "fixed_fee",
	OrderType:         "order_type",
	ExchangeRate:      "exchange_rate",
	ReceiptAmount:     "receipt_amount",
	ReceiptRealAmount: "receipt_real_amount",
	ReceiptCoin:       "receipt_coin",
}

// NewWithdrawDao creates and returns a new DAO object for table data access.
func NewWithdrawDao(handlers ...gdb.ModelHandler) *WithdrawDao {
	return &WithdrawDao{
		group:    "default",
		table:    "t_withdraw",
		columns:  withdrawColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WithdrawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WithdrawDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WithdrawDao) Columns() WithdrawColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WithdrawDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WithdrawDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
