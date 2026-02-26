// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractLoss is the golang structure for table contract_loss.
type ContractLoss struct {
	Id                int64       `json:"id"                  orm:"id"                  description:"主键"`
	DelegateType      int         `json:"delegate_type"       orm:"delegate_type"       description:"委托类型（0 限价 1 市价）"`
	Status            int         `json:"status"              orm:"status"              description:"状态  0  正常 1 删除  2 撤销"`
	PositionId        string      `json:"position_id"         orm:"position_id"         description:"仓位ID"`
	UserId            int64       `json:"user_id"             orm:"user_id"             description:"用户id"`
	EarnPrice         float64     `json:"earn_price"          orm:"earn_price"          description:"止盈触发价"`
	LosePrice         float64     `json:"lose_price"          orm:"lose_price"          description:"止损触发价"`
	CreateTime        *gtime.Time `json:"create_time"         orm:"create_time"         description:"创建时间"`
	EarnDelegatePrice float64     `json:"earn_delegate_price" orm:"earn_delegate_price" description:"止盈委托价"`
	LoseDelegatePrice float64     `json:"lose_delegate_price" orm:"lose_delegate_price" description:"止损委托价"`
	EarnNumber        float64     `json:"earn_number"         orm:"earn_number"         description:"止盈数量"`
	LoseNumber        float64     `json:"lose_number"         orm:"lose_number"         description:"止损数量"`
	LossType          int         `json:"loss_type"           orm:"loss_type"           description:"0 止盈    1止损"`
	UpdateTime        *gtime.Time `json:"update_time"         orm:"update_time"         description:"更新时间"`
	Type              int         `json:"type"                orm:"type"                description:""`
	Leverage          float64     `json:"leverage"            orm:"leverage"            description:""`
	Symbol            string      `json:"symbol"              orm:"symbol"              description:""`
}
