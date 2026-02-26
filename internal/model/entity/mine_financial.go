// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MineFinancial is the golang structure for table mine_financial.
type MineFinancial struct {
	Id                  int         `json:"id"                   orm:"id"                   description:""`
	Title               string      `json:"title"                orm:"title"                description:"标题"`
	Icon                string      `json:"icon"                 orm:"icon"                 description:"图标"`
	Status              int         `json:"status"               orm:"status"               description:"启用禁用(展示在前端)1开0关"`
	Days                string      `json:"days"                 orm:"days"                 description:"天数(如 7,10,30)"`
	DefaultOdds         float64     `json:"default_odds"         orm:"default_odds"         description:"违约利率"`
	MinOdds             float64     `json:"min_odds"             orm:"min_odds"             description:"最小日利率百分比"`
	MaxOdds             float64     `json:"max_odds"             orm:"max_odds"             description:"最大日利率百分比"`
	TimeLimit           int         `json:"time_limit"           orm:"time_limit"           description:"每人限购次数，0表示不限"`
	LimitMin            float64     `json:"limit_min"            orm:"limit_min"            description:"最小金额"`
	LimitMax            float64     `json:"limit_max"            orm:"limit_max"            description:"最大金额"`
	IsHot               int         `json:"is_hot"               orm:"is_hot"               description:"是否热销1是0否"`
	Sort                int         `json:"sort"                 orm:"sort"                 description:"排序"`
	CreateBy            string      `json:"create_by"            orm:"create_by"            description:"创建人"`
	CreateTime          *gtime.Time `json:"create_time"          orm:"create_time"          description:"创建时间"`
	UpdateBy            string      `json:"update_by"            orm:"update_by"            description:"更新人员"`
	UpdateTime          *gtime.Time `json:"update_time"          orm:"update_time"          description:"更新时间"`
	BuyPurchase         int         `json:"buy_purchase"         orm:"buy_purchase"         description:"购买次数"`
	AvgRate             float64     `json:"avg_rate"             orm:"avg_rate"             description:"日平均利率"`
	Coin                string      `json:"coin"                 orm:"coin"                 description:"币种"`
	Classify            string      `json:"classify"             orm:"classify"             description:"分类（0 普通  1 vip  2 增值）"`
	BasicInvestAmount   float64     `json:"basic_invest_amount"  orm:"basic_invest_amount"  description:"平台基础投资金额"`
	TotalInvestAmount   float64     `json:"total_invest_amount"  orm:"total_invest_amount"  description:"平台总投资额"`
	Level               int         `json:"level"                orm:"level"                description:"VIP等级"`
	Process             float64     `json:"process"              orm:"process"              description:"项目进度"`
	RemainAmount        float64     `json:"remain_amount"        orm:"remain_amount"        description:"剩余金额"`
	Remark              string      `json:"remark"               orm:"remark"               description:"标签"`
	PurchasedAmount     float64     `json:"purchased_amount"     orm:"purchased_amount"     description:"易购金额"`
	Problem             string      `json:"problem"              orm:"problem"              description:"常见问题"`
	ProdectIntroduction string      `json:"prodect_introduction" orm:"prodect_introduction" description:"产品介绍"`
}
