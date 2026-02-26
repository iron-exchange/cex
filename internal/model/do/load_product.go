// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoadProduct is the golang structure of table t_load_product for DAO operations like Where/Data.
type LoadProduct struct {
	g.Meta      `orm:"table:t_load_product, do:true"`
	Id          any         // 主键
	AmountMin   any         // 贷款最小额度
	AmountMax   any         // 贷款最大额度
	CycleType   any         // 周期类型  0-7天 1-14天 2-30天 ,,,,待补充
	RepayType   any         // 还款类型 0-到期一次换本息...待补充
	Status      any         // 状态 0 未开启 1已开启
	CreateBy    any         //
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         //
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 用户备注
	SearchValue any         //
	Odds        any         // 日利率（%）
	RepayOrg    any         // 还款机构
	IsFreeze    any         // 是否冻结  1=正常 2=冻结
}
