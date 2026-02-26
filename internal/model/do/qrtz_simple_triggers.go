// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzSimpleTriggers is the golang structure of table qrtz_simple_triggers for DAO operations like Where/Data.
type QrtzSimpleTriggers struct {
	g.Meta         `orm:"table:qrtz_simple_triggers, do:true"`
	SchedName      any // 调度名称
	TriggerName    any // qrtz_triggers表trigger_name的外键
	TriggerGroup   any // qrtz_triggers表trigger_group的外键
	RepeatCount    any // 重复的次数统计
	RepeatInterval any // 重复的间隔时间
	TimesTriggered any // 已经触发的次数
}
