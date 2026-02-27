package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 注册请求结构体
type RegisterReq struct {
	g.Meta        `path:"/user/register" tags:"User" method:"post" summary:"用户注册"`
	SignType      string `json:"signType" v:"required|in:LOGIN,PHONE,EMAIL,ADDRESS#请指定注册类型" dc:"注册类型: LOGIN/PHONE/EMAIL/ADDRESS"`
	LoginName     string `json:"loginName" dc:"用户名 (LOGIN 必填)"`
	LoginPassword string `json:"loginPassword" dc:"登录密码 (未传默认 123456)"`
	Phone         string `json:"phone" dc:"手机号 (PHONE 必填)"`
	Email         string `json:"email" dc:"邮箱 (EMAIL 必填)"`
	Address       string `json:"address" dc:"ERC20/TRC20 钱包地址直签授权注册 (ADDRESS 必填)"`
	Code          string `json:"code" dc:"验证码"`
	InviteCode    string `json:"inviteCode" dc:"邀请码/推广码"` // 原系统的推荐机制可能依赖
}

// RegisterRes 注册返回结构体
type RegisterRes struct {
	UserId    uint64 `json:"userId"`
	SignType  string `json:"signType"`
	LoginName string `json:"loginName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

// LoginReq 登录请求结构体
type LoginReq struct {
	g.Meta        `path:"/user/login" tags:"User" method:"post" summary:"用户登录"`
	SignType      string `json:"signType" v:"required|in:LOGIN,PHONE,EMAIL,ADDRESS#请指定登录类型" dc:"登录类型: LOGIN/PHONE/EMAIL/ADDRESS"`
	LoginName     string `json:"loginName" dc:"用户名"`
	Phone         string `json:"phone" dc:"手机号"`
	Email         string `json:"email" dc:"邮箱"`
	Address       string `json:"address" dc:"钱包地址"`
	LoginPassword string `json:"loginPassword" dc:"密码"`
	Code          string `json:"code" dc:"验证码"`
}

// LoginRes 登录返回结构体
type LoginRes struct {
	Token    string `json:"token" dc:"JWT 授权 Token"`
	Expire   string `json:"expire" dc:"Token 过期时间"`
	UserId   uint64 `json:"userId" dc:"用户ID"`
	SignType string `json:"signType" dc:"登录类型"`
}

// PwdSettReq 修改/设置登录密码
type PwdSettReq struct {
	g.Meta `path:"/user/pwdSett" tags:"User" method:"post" summary:"设置登录密码"`
	Pwd    string `json:"pwd" v:"required#请输入密码"`
}

// PwdSettRes 返回
type PwdSettRes struct{}

// TardPwdSetReq 修改交易密码（资金密码）
type TardPwdSetReq struct {
	g.Meta `path:"/user/tardPwdSet" tags:"User" method:"post" summary:"设置交易密码"`
	Pwd    string `json:"pwd" v:"required#请输入密码"`
}

// TardPwdSetRes 返回
type TardPwdSetRes struct{}

// BindPhoneReq 绑定手机
type BindPhoneReq struct {
	g.Meta `path:"/user/bindPhone" tags:"User" method:"post" summary:"绑定手机"`
	Phone  string `json:"phone" v:"required#请输入手机号"`
	Code   string `json:"code" v:"required#请输入验证码"`
}
type BindPhoneRes struct{}

// BindEmailReq 绑定邮箱
type BindEmailReq struct {
	g.Meta    `path:"/user/bindEmail" tags:"User" method:"post" summary:"绑定邮箱"`
	Email     string `json:"email" v:"required#请输入邮箱"`
	EmailCode string `json:"emailCode" v:"required#请输入验证码"`
}
type BindEmailRes struct{}

// UpdateUserAddressReq 地址静默绑定
type UpdateUserAddressReq struct {
	g.Meta  `path:"/user/updateUserAddress" tags:"User" method:"post" summary:"绑定钱包地址"`
	Address string `json:"address" v:"required#请输入钱包地址"`
	Type    string `json:"type" v:"required|in:ETH,TRON#钱包类型错误"` // 钱包类型 (ETH, TRON)
}
type UpdateUserAddressRes struct{}

// UploadKYCReq KYC 实名认证
type UploadKYCReq struct {
	g.Meta    `path:"/user/uploadKYC" tags:"User" method:"post" summary:"上传KYC实名认证"`
	RealName  string `json:"realName" v:"required#真实姓名必填"`
	IdCard    string `json:"idCard" v:"required#证件号码必填"`
	FrontUrl  string `json:"frontUrl" v:"required#正面照必填"`
	BackUrl   string `json:"backUrl" v:"required#反面照必填"`
	HandelUrl string `json:"handelUrl" v:"required#手持证件照必填"`
	Country   string `json:"country" dc:"国家代码"`
	CardType  string `json:"cardType" dc:"证件类型"`
	Flag      string `json:"flag" dc:"2 会触发风控开关校验"`
}
type UploadKYCRes struct{}

// SendCodeReq 发送验证码请求
type SendCodeReq struct {
	g.Meta `path:"/user/sendCode" tags:"User" method:"post" summary:"发送短信或邮箱验证码"`
	To     string `json:"to" v:"required#请输入手机号或邮箱"`
	Type   string `json:"type" v:"required|in:PHONE,EMAIL#验证码通道: PHONE/EMAIL"`
	Scene  string `json:"scene" v:"required|in:REGISTER,LOGIN,BIND,FORGET#业务场景: REGISTER/LOGIN/BIND/FORGET"`
}

type SendCodeRes struct {
	// 出于安全考虑，验证码绝不能在接口返回。此处仅作占位。
	// 实际验证码已通过后台日志/短信网关发送。
}
