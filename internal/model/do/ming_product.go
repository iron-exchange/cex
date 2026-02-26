// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MingProduct is the golang structure of table t_ming_product for DAO operations like Where/Data.
type MingProduct struct {
	g.Meta      `orm:"table:t_ming_product, do:true"`
	Id          any         //
	Title       any         // 标题
	Icon        any         // 图标
	Status      any         // 启用禁用(展示在前端)1开0关
	Days        any         // 天数(如 7,10,30)
	DefaultOdds any         // 违约利率
	MinOdds     any         // 最小日利率百分比
	MaxOdds     any         // 最大日利率百分比
	TimeLimit   any         // 每人限购次数，0表示不限
	LimitMin    any         // 最小金额
	LimitMax    any         // 最大金额
	Sort        any         // 排序
	CreateBy    any         // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新人员
	UpdateTime  *gtime.Time // 更新时间
	BuyPurchase any         // 购买次数
	Coin        any         // 币种
	Remark      any         // 标签
}
