// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondPeriodConfig is the golang structure of table t_second_period_config for DAO operations like Where/Data.
type SecondPeriodConfig struct {
	g.Meta      `orm:"table:t_second_period_config, do:true"`
	Id          any         // id
	SecondId    any         // 秒合约币种配置id
	Period      any         // 时间周期  单位秒
	Odds        any         // 赔率
	MaxAmount   any         // 最大金额
	MinAmount   any         // 最小金额
	Status      any         // 1开启 2关闭
	CreateBy    any         // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新人
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	SearchValue any         //
	Flag        any         // 全输标识
}
