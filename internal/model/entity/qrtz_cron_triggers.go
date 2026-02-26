// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzCronTriggers is the golang structure for table qrtz_cron_triggers.
type QrtzCronTriggers struct {
	SchedName      string `json:"sched_name"      orm:"sched_name"      description:"调度名称"`
	TriggerName    string `json:"trigger_name"    orm:"trigger_name"    description:"qrtz_triggers表trigger_name的外键"`
	TriggerGroup   string `json:"trigger_group"   orm:"trigger_group"   description:"qrtz_triggers表trigger_group的外键"`
	CronExpression string `json:"cron_expression" orm:"cron_expression" description:"cron表达式"`
	TimeZoneId     string `json:"time_zone_id"    orm:"time_zone_id"    description:"时区"`
}
