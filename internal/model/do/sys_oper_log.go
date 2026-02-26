// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperLog is the golang structure of table sys_oper_log for DAO operations like Where/Data.
type SysOperLog struct {
	g.Meta        `orm:"table:sys_oper_log, do:true"`
	OperId        any         // 日志主键
	Title         any         // 模块标题
	BusinessType  any         // 业务类型（0其它 1新增 2修改 3删除）
	Method        any         // 方法名称
	RequestMethod any         // 请求方式
	OperatorType  any         // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      any         // 操作人员
	DeptName      any         // 部门名称
	OperUrl       any         // 请求URL
	OperIp        any         // 主机地址
	OperLocation  any         // 操作地点
	OperParam     any         // 请求参数
	JsonResult    any         // 返回参数
	Status        any         // 操作状态（0正常 1异常）
	ErrorMsg      any         // 错误消息
	OperTime      *gtime.Time // 操作时间
	CostTime      any         // 消耗时间
}
