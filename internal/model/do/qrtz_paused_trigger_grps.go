// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzPausedTriggerGrps is the golang structure of table qrtz_paused_trigger_grps for DAO operations like Where/Data.
type QrtzPausedTriggerGrps struct {
	g.Meta       `orm:"table:qrtz_paused_trigger_grps, do:true"`
	SchedName    any // 调度名称
	TriggerGroup any // qrtz_triggers表trigger_group的外键
}
