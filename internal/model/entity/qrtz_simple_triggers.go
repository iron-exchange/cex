// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzSimpleTriggers is the golang structure for table qrtz_simple_triggers.
type QrtzSimpleTriggers struct {
	SchedName      string `json:"sched_name"      orm:"sched_name"      description:"调度名称"`
	TriggerName    string `json:"trigger_name"    orm:"trigger_name"    description:"qrtz_triggers表trigger_name的外键"`
	TriggerGroup   string `json:"trigger_group"   orm:"trigger_group"   description:"qrtz_triggers表trigger_group的外键"`
	RepeatCount    int64  `json:"repeat_count"    orm:"repeat_count"    description:"重复的次数统计"`
	RepeatInterval int64  `json:"repeat_interval" orm:"repeat_interval" description:"重复的间隔时间"`
	TimesTriggered int64  `json:"times_triggered" orm:"times_triggered" description:"已经触发的次数"`
}
