// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SymbolManage is the golang structure for table symbol_manage.
type SymbolManage struct {
	Id           int64       `json:"id"             orm:"id"             description:"主键id"`
	Symbol       string      `json:"symbol"         orm:"symbol"         description:"币种"`
	MinChargeNum float64     `json:"min_charge_num" orm:"min_charge_num" description:"最小兑换数量"`
	MaxChargeNum float64     `json:"max_charge_num" orm:"max_charge_num" description:"最大兑换数量"`
	Commission   float64     `json:"commission"     orm:"commission"     description:"手续费(%)"`
	Sort         int         `json:"sort"           orm:"sort"           description:"排序"`
	Enable       string      `json:"enable"         orm:"enable"         description:"1 启用 2 禁用"`
	Logo         string      `json:"logo"           orm:"logo"           description:"图标"`
	Market       string      `json:"market"         orm:"market"         description:"交易所"`
	Remark       string      `json:"remark"         orm:"remark"         description:"备注"`
	CreateBy     string      `json:"create_by"      orm:"create_by"      description:"创建人"`
	CreateTime   *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
	UpdateBy     string      `json:"update_by"      orm:"update_by"      description:"修改人"`
	UpdateTime   *gtime.Time `json:"update_time"    orm:"update_time"    description:"修改时间"`
	DelFlag      string      `json:"del_flag"       orm:"del_flag"       description:"0正常  2删除"`
	SearchValue  string      `json:"search_value"   orm:"search_value"   description:""`
}
