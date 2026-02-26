// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzTriggers is the golang structure for table qrtz_triggers.
type QrtzTriggers struct {
	SchedName    string `json:"sched_name"     orm:"sched_name"     description:"调度名称"`
	TriggerName  string `json:"trigger_name"   orm:"trigger_name"   description:"触发器的名字"`
	TriggerGroup string `json:"trigger_group"  orm:"trigger_group"  description:"触发器所属组的名字"`
	JobName      string `json:"job_name"       orm:"job_name"       description:"qrtz_job_details表job_name的外键"`
	JobGroup     string `json:"job_group"      orm:"job_group"      description:"qrtz_job_details表job_group的外键"`
	Description  string `json:"description"    orm:"description"    description:"相关介绍"`
	NextFireTime int64  `json:"next_fire_time" orm:"next_fire_time" description:"上一次触发时间（毫秒）"`
	PrevFireTime int64  `json:"prev_fire_time" orm:"prev_fire_time" description:"下一次触发时间（默认为-1表示不触发）"`
	Priority     int    `json:"priority"       orm:"priority"       description:"优先级"`
	TriggerState string `json:"trigger_state"  orm:"trigger_state"  description:"触发器状态"`
	TriggerType  string `json:"trigger_type"   orm:"trigger_type"   description:"触发器的类型"`
	StartTime    int64  `json:"start_time"     orm:"start_time"     description:"开始时间"`
	EndTime      int64  `json:"end_time"       orm:"end_time"       description:"结束时间"`
	CalendarName string `json:"calendar_name"  orm:"calendar_name"  description:"日程表名称"`
	MisfireInstr int    `json:"misfire_instr"  orm:"misfire_instr"  description:"补偿执行的策略"`
	JobData      string `json:"job_data"       orm:"job_data"       description:"存放持久化job对象"`
}
