// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondPeriodConfig is the golang structure for table second_period_config.
type SecondPeriodConfig struct {
	Id          int64       `json:"id"           orm:"id"           description:"id"`
	SecondId    int64       `json:"second_id"    orm:"second_id"    description:"秒合约币种配置id"`
	Period      int         `json:"period"       orm:"period"       description:"时间周期  单位秒"`
	Odds        float64     `json:"odds"         orm:"odds"         description:"赔率"`
	MaxAmount   float64     `json:"max_amount"   orm:"max_amount"   description:"最大金额"`
	MinAmount   float64     `json:"min_amount"   orm:"min_amount"   description:"最小金额"`
	Status      int         `json:"status"       orm:"status"       description:"1开启 2关闭"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建人"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	Flag        int         `json:"flag"         orm:"flag"         description:"全输标识"`
}
