// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserDetailDao is the data access object for the table t_app_user_detail.
type AppUserDetailDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  AppUserDetailColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// AppUserDetailColumns defines and stores column names for the table t_app_user_detail.
type AppUserDetailColumns struct {
	Id                  string //
	UserId              string //
	RealName            string // 真实姓名
	IdCard              string // 身份证号码
	FrontUrl            string // 身份证正面照片
	Country             string // 国际
	CardType            string //
	HandelUrl           string // 手持身份证照片
	BackUrl             string // 身份证反面照片
	UserTardPwd         string // 用户交易密码
	CreateBy            string //
	CreateTime          string //
	UpdateBy            string //
	UpdateTime          string //
	Remark              string //
	SearchValue         string //
	AuditStatusPrimary  string // 初级验证状态
	AuditStatusAdvanced string // 高级验证状态
	Credits             string // 信用分
	UserRechargeAddress string // 用户充值地址
	WinNum              string // 连赢场次
	LoseNum             string // 连输场次
	TradeFlag           string // 交易是否被限制 1 为限制
	AmountFlag          string // 金额是否被限制 1 为限制
	PushMessage         string // 金额限制提示语
	TradeMessage        string // 交易限制提示语
	OperateTime         string // 实名认证时间
}

// appUserDetailColumns holds the columns for the table t_app_user_detail.
var appUserDetailColumns = AppUserDetailColumns{
	Id:                  "id",
	UserId:              "user_id",
	RealName:            "real_name",
	IdCard:              "id_card",
	FrontUrl:            "front_url",
	Country:             "country",
	CardType:            "card_type",
	HandelUrl:           "handel_url",
	BackUrl:             "back_url",
	UserTardPwd:         "user_tard_pwd",
	CreateBy:            "create_by",
	CreateTime:          "create_time",
	UpdateBy:            "update_by",
	UpdateTime:          "update_time",
	Remark:              "remark",
	SearchValue:         "search_value",
	AuditStatusPrimary:  "audit_status_primary",
	AuditStatusAdvanced: "audit_status_advanced",
	Credits:             "credits",
	UserRechargeAddress: "user_recharge_address",
	WinNum:              "win_num",
	LoseNum:             "lose_num",
	TradeFlag:           "trade_flag",
	AmountFlag:          "amount_flag",
	PushMessage:         "push_message",
	TradeMessage:        "trade_message",
	OperateTime:         "operate_time",
}

// NewAppUserDetailDao creates and returns a new DAO object for table data access.
func NewAppUserDetailDao(handlers ...gdb.ModelHandler) *AppUserDetailDao {
	return &AppUserDetailDao{
		group:    "default",
		table:    "t_app_user_detail",
		columns:  appUserDetailColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppUserDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppUserDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppUserDetailDao) Columns() AppUserDetailColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppUserDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppUserDetailDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppUserDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
