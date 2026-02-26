// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OptionRules is the golang structure of table t_option_rules for DAO operations like Where/Data.
type OptionRules struct {
	g.Meta      `orm:"table:t_option_rules, do:true"`
	Id          any         //
	Title       any         // 标题
	Language    any         // 语言 en zh
	Content     any         // 内容
	IsShow      any         // 是否展示 0展示  2不展示
	CreateTime  *gtime.Time // 创建时间
	Type        any         // 0=服务条款 1=秒合约说明 2=币币交易说明 3=代理活动 4=U本位合约说明 5=注册隐私政策 6=注册使用条款 7=贷款规则
	SearchValue any         //
}
