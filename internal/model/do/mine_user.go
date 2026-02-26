// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MineUser is the golang structure of table t_mine_user for DAO operations like Where/Data.
type MineUser struct {
	g.Meta    `orm:"table:t_mine_user, do:true"`
	UserId    any // 用户id
	Id        any // 挖矿产品id
	TimeLimit any // 限购次数
}
