// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppRecharge is the golang structure of table t_app_recharge for DAO operations like Where/Data.
type AppRecharge struct {
	g.Meta         `orm:"table:t_app_recharge, do:true"`
	Id             any         // 卡ID
	CreateBy       any         //
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         //
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         //
	UserId         any         // 所有者ID
	Username       any         // 用户名
	Amount         any         // 充值金额
	Bonus          any         //
	Status         any         // 状态
	SerialId       any         // 订单号
	TxId           any         // 第三方支付订单号
	Type           any         // 类型
	SearchValue    any         //
	Address        any         // 充值地址
	Tree           any         //
	Coin           any         // 币总
	ToAddress      any         // 入款地址
	BlockTime      *gtime.Time // 区块时间
	Host           any         //
	RealAmount     any         // 实际到账金额
	FileName       any         // 充值凭证
	RechargeRemark any         //
	NoticeFlag     any         // 通知字段 0未通知 1通知了
	AppParentIds   any         // app代理ids
	AdminParentIds any         // 后台代理ids
	OperateTime    *gtime.Time // 操作时间
	OrderType      any         // 订单类型 1/null=充值  2=彩金赠送
}
