// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzJobDetails is the golang structure of table qrtz_job_details for DAO operations like Where/Data.
type QrtzJobDetails struct {
	g.Meta           `orm:"table:qrtz_job_details, do:true"`
	SchedName        any // 调度名称
	JobName          any // 任务名称
	JobGroup         any // 任务组名
	Description      any // 相关介绍
	JobClassName     any // 执行任务类名称
	IsDurable        any // 是否持久化
	IsNonconcurrent  any // 是否并发
	IsUpdateData     any // 是否更新数据
	RequestsRecovery any // 是否接受恢复执行
	JobData          any // 存放持久化job对象
}
