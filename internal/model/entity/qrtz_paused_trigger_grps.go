// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzPausedTriggerGrps is the golang structure for table qrtz_paused_trigger_grps.
type QrtzPausedTriggerGrps struct {
	SchedName    string `json:"sched_name"    orm:"sched_name"    description:"调度名称"`
	TriggerGroup string `json:"trigger_group" orm:"trigger_group" description:"qrtz_triggers表trigger_group的外键"`
}
