// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzFiredTriggers is the golang structure for table qrtz_fired_triggers.
type QrtzFiredTriggers struct {
	SchedName        string `json:"sched_name"        orm:"sched_name"        description:"调度名称"`
	EntryId          string `json:"entry_id"          orm:"entry_id"          description:"调度器实例id"`
	TriggerName      string `json:"trigger_name"      orm:"trigger_name"      description:"qrtz_triggers表trigger_name的外键"`
	TriggerGroup     string `json:"trigger_group"     orm:"trigger_group"     description:"qrtz_triggers表trigger_group的外键"`
	InstanceName     string `json:"instance_name"     orm:"instance_name"     description:"调度器实例名"`
	FiredTime        int64  `json:"fired_time"        orm:"fired_time"        description:"触发的时间"`
	SchedTime        int64  `json:"sched_time"        orm:"sched_time"        description:"定时器制定的时间"`
	Priority         int    `json:"priority"          orm:"priority"          description:"优先级"`
	State            string `json:"state"             orm:"state"             description:"状态"`
	JobName          string `json:"job_name"          orm:"job_name"          description:"任务名称"`
	JobGroup         string `json:"job_group"         orm:"job_group"         description:"任务组名"`
	IsNonconcurrent  string `json:"is_nonconcurrent"  orm:"is_nonconcurrent"  description:"是否并发"`
	RequestsRecovery string `json:"requests_recovery" orm:"requests_recovery" description:"是否接受恢复执行"`
}
