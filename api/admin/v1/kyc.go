package v1

import "github.com/gogf/gf/v2/frame/g"

type KycInfo struct {
	UserId              int    `json:"userId"`
	RealName            string `json:"realName"`
	IdCard              string `json:"idCard"`
	FrontUrl            string `json:"frontUrl"`
	BackUrl             string `json:"backUrl"`
	HandelUrl           string `json:"handelUrl"`
	AuditStatusPrimary  int    `json:"auditStatusPrimary"`
	AuditStatusAdvanced int    `json:"auditStatusAdvanced"`
	CreateTime          string `json:"createTime"`
}

// GetKycListReq 获取实名认证列表
type GetKycListReq struct {
	g.Meta      `path:"/kyc/list" tags:"AdminUser" method:"get" summary:"获取实名认证列表"`
	Page        int  `json:"page" d:"1"`
	Size        int  `json:"size" d:"20"`
	UserId      int  `json:"userId" dc:"根据用户ID查询"`
	AuditStatus *int `json:"auditStatus" dc:"审核状态 (0:待审 1:通过 2:拒绝)"`
}

type GetKycListRes struct {
	List  []KycInfo `json:"list"`
	Total int       `json:"total"`
}

// AuditKycReq 审核/驳回实名认证
type AuditKycReq struct {
	g.Meta     `path:"/kyc/audit" tags:"AdminUser" method:"post" summary:"审核实名认证"`
	UserId     int    `json:"userId" v:"required#用户ID不能为空"`
	Action     string `json:"action" v:"required|in:pass,reject#操作仅支持 pass 或 reject"`
	AuditLevel int    `json:"auditLevel" v:"required|in:1,2#1初级审核 2高级审核"`
	Remark     string `json:"remark" dc:"驳回原因等备注信息"`
}

type AuditKycRes struct{}
