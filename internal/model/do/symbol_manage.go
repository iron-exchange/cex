// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SymbolManage is the golang structure of table t_symbol_manage for DAO operations like Where/Data.
type SymbolManage struct {
	g.Meta       `orm:"table:t_symbol_manage, do:true"`
	Id           any         // 主键id
	Symbol       any         // 币种
	MinChargeNum any         // 最小兑换数量
	MaxChargeNum any         // 最大兑换数量
	Commission   any         // 手续费(%)
	Sort         any         // 排序
	Enable       any         // 1 启用 2 禁用
	Logo         any         // 图标
	Market       any         // 交易所
	Remark       any         // 备注
	CreateBy     any         // 创建人
	CreateTime   *gtime.Time // 创建时间
	UpdateBy     any         // 修改人
	UpdateTime   *gtime.Time // 修改时间
	DelFlag      any         // 0正常  2删除
	SearchValue  any         //
}
