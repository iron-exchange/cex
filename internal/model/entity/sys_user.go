// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	UserId      int64       `json:"user_id"     orm:"user_id"     description:"用户ID"`
	DeptId      int64       `json:"dept_id"     orm:"dept_id"     description:"部门ID"`
	UserName    string      `json:"user_name"   orm:"user_name"   description:"用户账号"`
	NickName    string      `json:"nick_name"   orm:"nick_name"   description:"用户昵称"`
	UserType    string      `json:"user_type"   orm:"user_type"   description:"用户类型（0普通用户 1 组长 2 代理）"`
	Email       string      `json:"email"       orm:"email"       description:"用户邮箱"`
	Phonenumber string      `json:"phonenumber" orm:"phonenumber" description:"手机号码"`
	Sex         string      `json:"sex"         orm:"sex"         description:"用户性别（0男 1女 2未知）"`
	Avatar      string      `json:"avatar"      orm:"avatar"      description:"头像地址"`
	Password    string      `json:"password"    orm:"password"    description:"密码"`
	Status      string      `json:"status"      orm:"status"      description:"帐号状态（0正常 1停用）"`
	GoogleKey   string      `json:"google_key"  orm:"google_key"  description:"谷歌验证码key"`
	DelFlag     string      `json:"del_flag"    orm:"del_flag"    description:"删除标志（0代表存在 2代表删除）"`
	LoginIp     string      `json:"login_ip"    orm:"login_ip"    description:"最后登录IP"`
	LoginDate   *gtime.Time `json:"login_date"  orm:"login_date"  description:"最后登录时间"`
	CreateBy    string      `json:"create_by"   orm:"create_by"   description:"创建者"`
	CreateTime  *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
	UpdateBy    string      `json:"update_by"   orm:"update_by"   description:"更新者"`
	UpdateTime  *gtime.Time `json:"update_time" orm:"update_time" description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`
	ParentId    int64       `json:"parent_id"   orm:"parent_id"   description:"组长ID"`
}
