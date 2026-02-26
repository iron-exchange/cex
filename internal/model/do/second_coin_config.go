// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondCoinConfig is the golang structure of table t_second_coin_config for DAO operations like Where/Data.
type SecondCoinConfig struct {
	g.Meta      `orm:"table:t_second_coin_config, do:true"`
	Id          any         // id
	Symbol      any         // 合约交易对
	Market      any         // 所属交易所
	Status      any         // 是否启用 2关闭 1启用
	ShowFlag    any         // 是否展示 2不展示 1展示
	Coin        any         // 币种
	Sort        any         // 排序
	CreateBy    any         // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新人
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	SearchValue any         //
	Logo        any         // 图标
	BaseCoin    any         // 结算币种
	ShowSymbol  any         // 展示币种
	Type        any         // 币种类型 1 外汇  2虚拟币
}
