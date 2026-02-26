// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ExchangeCoinRecord is the golang structure of table t_exchange_coin_record for DAO operations like Where/Data.
type ExchangeCoinRecord struct {
	g.Meta         `orm:"table:t_exchange_coin_record, do:true"`
	Id             any         //
	FromCoin       any         //
	ToCoin         any         //
	UserId         any         // 用户id
	Username       any         // 用户名称
	Address        any         // 用户地址
	Status         any         // 兑换状态0:已提交;1:成功;2失败
	Amount         any         // 金额
	ThirdRate      any         // 三方汇率
	SystemRate     any         // 系统汇率
	AdminParentIds any         //
	CreateBy       any         //
	CreateTime     *gtime.Time //
	UpdateBy       any         //
	UpdateTime     *gtime.Time //
	Remark         any         //
	SearchValue    any         //
}
