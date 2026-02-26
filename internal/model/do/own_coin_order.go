// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoinOrder is the golang structure of table t_own_coin_order for DAO operations like Where/Data.
type OwnCoinOrder struct {
	g.Meta         `orm:"table:t_own_coin_order, do:true"`
	Id             any         // 主键
	UserId         any         // 用户ID
	OrderId        any         // 订单ID
	OwnId          any         // 申购币种ID
	OwnCoin        any         // 申购币种
	Amount         any         // 申购额（usdt）
	Number         any         // 申购数量
	Price          any         // 申购价
	Status         any         // 状态
	AdminUserIds   any         // 上级用户IDS
	AdminParentIds any         // 上级后台用户IDS
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
	CreateBy       any         //
	UpdateBy       any         //
	Remark         any         //
}
