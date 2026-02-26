// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzLocks is the golang structure of table qrtz_locks for DAO operations like Where/Data.
type QrtzLocks struct {
	g.Meta    `orm:"table:qrtz_locks, do:true"`
	SchedName any // 调度名称
	LockName  any // 悲观锁名称
}
