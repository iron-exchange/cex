// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppSubTransfer is the golang structure of table t_app_sub_transfer for DAO operations like Where/Data.
type AppSubTransfer struct {
	g.Meta     `orm:"table:t_app_sub_transfer, do:true"`
	Id         any         //
	UserId     any         // 用户id
	Email      any         // 子账号邮箱
	EmailType  any         // 1转入，2转出
	Type       any         // 划转类型："SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	Tranid     any         // 订单号
	Amount     any         // 价格
	Asset      any         // 货币类型
	Status     any         // 1正常 2失败
	CreateTime *gtime.Time // 创建时间
}
