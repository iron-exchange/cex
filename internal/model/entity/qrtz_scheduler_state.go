// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzSchedulerState is the golang structure for table qrtz_scheduler_state.
type QrtzSchedulerState struct {
	SchedName       string `json:"sched_name"        orm:"sched_name"        description:"调度名称"`
	InstanceName    string `json:"instance_name"     orm:"instance_name"     description:"实例名称"`
	LastCheckinTime int64  `json:"last_checkin_time" orm:"last_checkin_time" description:"上次检查时间"`
	CheckinInterval int64  `json:"checkin_interval"  orm:"checkin_interval"  description:"检查间隔时间"`
}
