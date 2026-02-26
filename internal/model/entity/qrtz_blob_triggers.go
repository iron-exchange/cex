// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzBlobTriggers is the golang structure for table qrtz_blob_triggers.
type QrtzBlobTriggers struct {
	SchedName    string `json:"sched_name"    orm:"sched_name"    description:"调度名称"`
	TriggerName  string `json:"trigger_name"  orm:"trigger_name"  description:"qrtz_triggers表trigger_name的外键"`
	TriggerGroup string `json:"trigger_group" orm:"trigger_group" description:"qrtz_triggers表trigger_group的外键"`
	BlobData     string `json:"blob_data"     orm:"blob_data"     description:"存放持久化Trigger对象"`
}
