// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondContractOrder is the golang structure of table t_second_contract_order for DAO operations like Where/Data.
type SecondContractOrder struct {
	g.Meta             `orm:"table:t_second_contract_order, do:true"`
	Id                 any         //
	OrderNo            any         // 订单号
	Symbol             any         // 交易对
	Type               any         // 类型
	UserId             any         // 用户id
	UserAddress        any         // 用户地址
	BetContent         any         // 预测方向:0 涨 1跌
	OpenResult         any         // 开奖结果
	Status             any         // 订单状态 0参与中 1已开奖 2已撤销
	RateFlag           any         // 是否全输
	BetAmount          any         // 投注金额
	RewardAmount       any         // 获奖金额
	CompensationAmount any         // 赔偿金额
	CreateTime         *gtime.Time // 创建时间
	OpenPrice          any         // 开盘价格
	ClosePrice         any         // 关盘价格
	OpenTime           any         // 开盘时间
	CloseTime          any         // 关盘时间
	CoinSymbol         any         // 交易币符号
	BaseSymbol         any         // 结算币符号
	Sign               any         // 订单标记 0正常  1包赢  2包输
	ManualIntervention any         // 是否人工干预 0是 1否
	Rate               any         //
	CreateBy           any         //
	UpdateTime         *gtime.Time //
	UpdateBy           any         //
	SearchValue        any         //
	Remark             any         //
	AdminParentIds     any         // 后台代理ID
	IsHandling         any         // 行锁
}
