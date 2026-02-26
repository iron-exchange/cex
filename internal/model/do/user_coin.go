// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserCoin is the golang structure of table t_user_coin for DAO operations like Where/Data.
type UserCoin struct {
	g.Meta `orm:"table:t_user_coin, do:true"`
	Id     any // 主键
	UserId any // 用户id
	Coin   any // 币种
	Icon   any // 图标
}
