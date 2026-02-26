// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoin is the golang structure for table own_coin.
type OwnCoin struct {
	Id              int64       `json:"id"               orm:"id"               description:"主键ID"`
	Coin            string      `json:"coin"             orm:"coin"             description:"币种"`
	Logo            string      `json:"logo"             orm:"logo"             description:"图标"`
	ReferCoin       string      `json:"refer_coin"       orm:"refer_coin"       description:"参考币种"`
	ReferMarket     string      `json:"refer_market"     orm:"refer_market"     description:"参考币种交易所"`
	ShowSymbol      string      `json:"show_symbol"      orm:"show_symbol"      description:"展示名称"`
	Price           float64     `json:"price"            orm:"price"            description:"初始价格（单位USDT）"`
	Proportion      float64     `json:"proportion"       orm:"proportion"       description:"价格百分比"`
	RaisingAmount   float64     `json:"raising_amount"   orm:"raising_amount"   description:"私募发行量"`
	RaisedAmount    float64     `json:"raised_amount"    orm:"raised_amount"    description:"已筹集额度"`
	PurchaseLimit   int         `json:"purchase_limit"   orm:"purchase_limit"   description:"预购上限"`
	TotalAmount     float64     `json:"total_amount"     orm:"total_amount"     description:"总发行量"`
	ParticipantsNum int         `json:"participants_num" orm:"participants_num" description:"参与人数"`
	RaisingTime     int         `json:"raising_time"     orm:"raising_time"     description:"筹集期限"`
	BeginTime       *gtime.Time `json:"begin_time"       orm:"begin_time"       description:"开始时间"`
	EndTime         *gtime.Time `json:"end_time"         orm:"end_time"         description:"结束时间"`
	Introduce       string      `json:"introduce"        orm:"introduce"        description:"介绍"`
	Status          int         `json:"status"           orm:"status"           description:"1.未发布  2.筹集中 3 筹集成功 4.筹集失败"`
	CreateBy        string      `json:"create_by"        orm:"create_by"        description:"创建人"`
	CreateTime      *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy        string      `json:"update_by"        orm:"update_by"        description:"更新者"`
	UpdateTime      *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark          string      `json:"remark"           orm:"remark"           description:"备注"`
}
