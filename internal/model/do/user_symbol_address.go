// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserSymbolAddress is the golang structure of table t_user_symbol_address for DAO operations like Where/Data.
type UserSymbolAddress struct {
	g.Meta      `orm:"table:t_user_symbol_address, do:true"`
	Id          any // 主键id
	UserId      any // 用户id
	Symbol      any // 币种
	Address     any // 充值地址
	SearchValue any //
}
