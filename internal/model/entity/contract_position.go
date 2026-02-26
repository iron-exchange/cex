// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractPosition is the golang structure for table contract_position.
type ContractPosition struct {
	Id               int64       `json:"id"                orm:"id"                description:"主键"`
	Type             int         `json:"type"              orm:"type"              description:"(0 买多 1卖空)"`
	DelegateType     int         `json:"delegate_type"     orm:"delegate_type"     description:"委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）"`
	Status           int         `json:"status"            orm:"status"            description:"状态  0 （等待成交  1 完全成交"`
	Amount           float64     `json:"amount"            orm:"amount"            description:"保证金"`
	OpenNum          float64     `json:"open_num"          orm:"open_num"          description:"持仓数量"`
	OpenPrice        float64     `json:"open_price"        orm:"open_price"        description:"开仓均价"`
	ClosePrice       float64     `json:"close_price"       orm:"close_price"       description:"预计强平价"`
	OrderNo          string      `json:"order_no"          orm:"order_no"          description:"仓位编号"`
	UserId           int64       `json:"user_id"           orm:"user_id"           description:"用户id"`
	OpenFee          float64     `json:"open_fee"          orm:"open_fee"          description:"开仓手续费"`
	Leverage         float64     `json:"leverage"          orm:"leverage"          description:"杠杆"`
	Symbol           string      `json:"symbol"            orm:"symbol"            description:"交易对"`
	CreateTime       *gtime.Time `json:"create_time"       orm:"create_time"       description:"创建时间"`
	AdjustAmount     float64     `json:"adjust_amount"     orm:"adjust_amount"     description:"调整保证金"`
	Earn             float64     `json:"earn"              orm:"earn"              description:"收益"`
	DealPrice        float64     `json:"deal_price"        orm:"deal_price"        description:"成交价"`
	DealNum          float64     `json:"deal_num"          orm:"deal_num"          description:"成交量"`
	DealTime         *gtime.Time `json:"deal_time"         orm:"deal_time"         description:"成交时间"`
	SellFee          float64     `json:"sell_fee"          orm:"sell_fee"          description:"卖出手续费"`
	RemainMargin     float64     `json:"remain_margin"     orm:"remain_margin"     description:"剩余保证金"`
	AssetFee         float64     `json:"asset_fee"         orm:"asset_fee"         description:"周期手续费"`
	EntrustmentValue float64     `json:"entrustment_value" orm:"entrustment_value" description:""`
	DealValue        float64     `json:"deal_value"        orm:"deal_value"        description:""`
	UpdateTime       *gtime.Time `json:"update_time"       orm:"update_time"       description:""`
	AdminParentIds   string      `json:"admin_parent_ids"  orm:"admin_parent_ids"  description:"代理IDS"`
	AuditStatus      int         `json:"audit_status"      orm:"audit_status"      description:"审核"`
	DeliveryDays     int         `json:"delivery_days"     orm:"delivery_days"     description:"交割时间"`
	MinMargin        float64     `json:"min_margin"        orm:"min_margin"        description:"最小保证金"`
	LossRate         float64     `json:"loss_rate"         orm:"loss_rate"         description:"止损率"`
	EarnRate         float64     `json:"earn_rate"         orm:"earn_rate"         description:"止盈率"`
	SubTime          *gtime.Time `json:"sub_time"          orm:"sub_time"          description:"提交时间"`
}
