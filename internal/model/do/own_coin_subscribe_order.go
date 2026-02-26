// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoinSubscribeOrder is the golang structure of table t_own_coin_subscribe_order for DAO operations like Where/Data.
type OwnCoinSubscribeOrder struct {
	g.Meta      `orm:"table:t_own_coin_subscribe_order, do:true"`
	Id          any         // 主键
	SubscribeId any         // 订阅单号ID
	UserId      any         // 用户ID
	OrderId     any         // 订单ID
	OwnId       any         // 申购币种ID
	OwnCoin     any         // 申购币种
	AmountLimit any         // 申购额（usdt）
	NumLimit    any         // 申购数量上限
	Price       any         // 申购价
	Status      any         // 状态，1订阅中、2订阅成功、3成功消息推送完成
	Remark      any         // 备注
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	CreateBy    any         //
}
