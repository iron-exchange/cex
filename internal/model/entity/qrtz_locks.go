// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzLocks is the golang structure for table qrtz_locks.
type QrtzLocks struct {
	SchedName string `json:"sched_name" orm:"sched_name" description:"调度名称"`
	LockName  string `json:"lock_name"  orm:"lock_name"  description:"悲观锁名称"`
}
