// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoin is the golang structure of table t_own_coin for DAO operations like Where/Data.
type OwnCoin struct {
	g.Meta          `orm:"table:t_own_coin, do:true"`
	Id              any         // 主键ID
	Coin            any         // 币种
	Logo            any         // 图标
	ReferCoin       any         // 参考币种
	ReferMarket     any         // 参考币种交易所
	ShowSymbol      any         // 展示名称
	Price           any         // 初始价格（单位USDT）
	Proportion      any         // 价格百分比
	RaisingAmount   any         // 私募发行量
	RaisedAmount    any         // 已筹集额度
	PurchaseLimit   any         // 预购上限
	TotalAmount     any         // 总发行量
	ParticipantsNum any         // 参与人数
	RaisingTime     any         // 筹集期限
	BeginTime       *gtime.Time // 开始时间
	EndTime         *gtime.Time // 结束时间
	Introduce       any         // 介绍
	Status          any         // 1.未发布  2.筹集中 3 筹集成功 4.筹集失败
	CreateBy        any         // 创建人
	CreateTime      *gtime.Time // 创建时间
	UpdateBy        any         // 更新者
	UpdateTime      *gtime.Time // 更新时间
	Remark          any         // 备注
}
