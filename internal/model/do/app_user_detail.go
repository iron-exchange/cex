// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUserDetail is the golang structure of table t_app_user_detail for DAO operations like Where/Data.
type AppUserDetail struct {
	g.Meta              `orm:"table:t_app_user_detail, do:true"`
	Id                  any         //
	UserId              any         //
	RealName            any         // 真实姓名
	IdCard              any         // 身份证号码
	FrontUrl            any         // 身份证正面照片
	Country             any         // 国际
	CardType            any         //
	HandelUrl           any         // 手持身份证照片
	BackUrl             any         // 身份证反面照片
	UserTardPwd         any         // 用户交易密码
	CreateBy            any         //
	CreateTime          *gtime.Time //
	UpdateBy            any         //
	UpdateTime          *gtime.Time //
	Remark              any         //
	SearchValue         any         //
	AuditStatusPrimary  any         // 初级验证状态
	AuditStatusAdvanced any         // 高级验证状态
	Credits             any         // 信用分
	UserRechargeAddress any         // 用户充值地址
	WinNum              any         // 连赢场次
	LoseNum             any         // 连输场次
	TradeFlag           any         // 交易是否被限制 1 为限制
	AmountFlag          any         // 金额是否被限制 1 为限制
	PushMessage         any         // 金额限制提示语
	TradeMessage        any         // 交易限制提示语
	OperateTime         *gtime.Time // 实名认证时间
}
