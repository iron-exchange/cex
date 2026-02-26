// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzJobDetails is the golang structure for table qrtz_job_details.
type QrtzJobDetails struct {
	SchedName        string `json:"sched_name"        orm:"sched_name"        description:"调度名称"`
	JobName          string `json:"job_name"          orm:"job_name"          description:"任务名称"`
	JobGroup         string `json:"job_group"         orm:"job_group"         description:"任务组名"`
	Description      string `json:"description"       orm:"description"       description:"相关介绍"`
	JobClassName     string `json:"job_class_name"    orm:"job_class_name"    description:"执行任务类名称"`
	IsDurable        string `json:"is_durable"        orm:"is_durable"        description:"是否持久化"`
	IsNonconcurrent  string `json:"is_nonconcurrent"  orm:"is_nonconcurrent"  description:"是否并发"`
	IsUpdateData     string `json:"is_update_data"    orm:"is_update_data"    description:"是否更新数据"`
	RequestsRecovery string `json:"requests_recovery" orm:"requests_recovery" description:"是否接受恢复执行"`
	JobData          string `json:"job_data"          orm:"job_data"          description:"存放持久化job对象"`
}
