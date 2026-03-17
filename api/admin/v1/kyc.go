package v1

import "github.com/gogf/gf/v2/frame/g"

// AuditKycReq 审核/驳回实名认证
type AuditKycReq struct {
	g.Meta     `path:"/kyc/audit" tags:"AdminUser" method:"post" summary:"审核实名认证"`
	UserId     int    `json:"userId" v:"required#用户ID不能为空"`
	Action     string `json:"action" v:"required|in:pass,reject#操作仅支持 pass 或 reject"`
	AuditLevel int    `json:"auditLevel" v:"required|in:1,2#1初级审核 2高级审核"`
	Remark     string `json:"remark" dc:"驳回原因等备注信息"`
}

type AuditKycRes struct{}
