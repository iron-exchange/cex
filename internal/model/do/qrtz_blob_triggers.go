// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzBlobTriggers is the golang structure of table qrtz_blob_triggers for DAO operations like Where/Data.
type QrtzBlobTriggers struct {
	g.Meta       `orm:"table:qrtz_blob_triggers, do:true"`
	SchedName    any // 调度名称
	TriggerName  any // qrtz_triggers表trigger_name的外键
	TriggerGroup any // qrtz_triggers表trigger_group的外键
	BlobData     any // 存放持久化Trigger对象
}
