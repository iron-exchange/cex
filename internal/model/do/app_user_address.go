// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserAddress is the golang structure of table t_app_user_address for DAO operations like Where/Data.
type AppUserAddress struct {
	g.Meta       `orm:"table:t_app_user_address, do:true"`
	Id           any //
	UserId       any //
	Symbol       any // 钱包类型
	Address      any // 钱包地址
	BinanceEmail any // 币安子账号地址
}
