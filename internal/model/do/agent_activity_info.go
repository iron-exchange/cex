// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentActivityInfo is the golang structure of table t_agent_activity_info for DAO operations like Where/Data.
type AgentActivityInfo struct {
	g.Meta     `orm:"table:t_agent_activity_info, do:true"`
	Id         any         // id
	Type       any         // 1 充值返利 2挖矿返利
	Amount     any         // 返利金额
	CoinType   any         // 币种
	FromId     any         // 返利用户
	UserId     any         // 用户id
	CreateBy   any         //
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         //
	UpdateTime *gtime.Time // 更新时间
	Status     any         // 1  待返  2  已返
	LoginName  any         //
	SerialId   any         //
}
