// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OptionRules is the golang structure for table option_rules.
type OptionRules struct {
	Id          int64       `json:"id"           orm:"id"           description:""`
	Title       string      `json:"title"        orm:"title"        description:"标题"`
	Language    string      `json:"language"     orm:"language"     description:"语言 en zh"`
	Content     string      `json:"content"      orm:"content"      description:"内容"`
	IsShow      string      `json:"is_show"      orm:"is_show"      description:"是否展示 0展示  2不展示"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	Type        int         `json:"type"         orm:"type"         description:"0=服务条款 1=秒合约说明 2=币币交易说明 3=代理活动 4=U本位合约说明 5=注册隐私政策 6=注册使用条款 7=贷款规则"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
}
