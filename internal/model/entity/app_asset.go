// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppAsset is the golang structure for table app_asset.
type AppAsset struct {
	Id                   int         `json:"id"                     orm:"id"                     description:""`
	UserId               int64       `json:"user_id"                orm:"user_id"                description:""`
	Adress               string      `json:"adress"                 orm:"adress"                 description:"地址"`
	Symbol               string      `json:"symbol"                 orm:"symbol"                 description:"币种"`
	Amout                float64     `json:"amout"                  orm:"amout"                  description:"资产总额"`
	OccupiedAmount       float64     `json:"occupied_amount"        orm:"occupied_amount"        description:"占用资产"`
	AvailableAmount      float64     `json:"available_amount"       orm:"available_amount"       description:"可用资产"`
	AvailableAmountDaily float64     `json:"available_amount_daily" orm:"available_amount_daily" description:"每日余额（0点时分的余额，提现会减少）"`
	CodingVolumeDaily    float64     `json:"coding_volume_daily"    orm:"coding_volume_daily"    description:"每日打码量（24点之前，提现会减少）"`
	Type                 string      `json:"type"                   orm:"type"                   description:"资产类型 1=平台资产 2=理财资产 3=合约账户"`
	CreateBy             string      `json:"create_by"              orm:"create_by"              description:"创建人"`
	CreateTime           *gtime.Time `json:"create_time"            orm:"create_time"            description:"创建时间"`
	UpdateBy             string      `json:"update_by"              orm:"update_by"              description:"更新人"`
	UpdateTime           *gtime.Time `json:"update_time"            orm:"update_time"            description:"更新时间"`
	Remark               string      `json:"remark"                 orm:"remark"                 description:"备注"`
	SearchValue          string      `json:"search_value"           orm:"search_value"           description:""`
}
