package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/user"
	"GoCEX/internal/service/middleware"
)

// Controller v1 版本的用户接口
type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Register 注册接口
func (c *Controller) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 调用 logic 层，由于暂未用 gf gen service，直接使用 logic 下的单例
	res, err = user.New().Register(ctx, req)
	return
}

// Login 登录接口
func (c *Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 1. 验证用户名和密码，查库
	u, accountVal, err := user.New().Login(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2. 签发 JWT (利用 ctx 透传数据给 gf-jwt 内部的 Authenticator 和 PayloadFunc)
	payloadData := g.Map{
		"userId":  u.UserId,
		"account": accountVal,
	}
	g.RequestFromCtx(ctx).SetCtxVar("jwt_payload", payloadData)
	tokenString, expire, errToken := middleware.Auth.TokenGenerator(payloadData)
	if errToken != nil {
		return nil, errToken
	}

	return &v1.LoginRes{
		Token:    tokenString,
		Expire:   expire.Format("2006-01-02 15:04:05"),
		UserId:   uint64(u.UserId),
		SignType: req.SignType,
	}, nil
}

// PwdSett 修改/设置登录密码
func (c *Controller) PwdSett(ctx context.Context, req *v1.PwdSettReq) (res *v1.PwdSettRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().PwdSett(ctx, userId, req)
	return &v1.PwdSettRes{}, err
}

// TardPwdSet 修改交易密码（资金密码）
func (c *Controller) TardPwdSet(ctx context.Context, req *v1.TardPwdSetReq) (res *v1.TardPwdSetRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().TardPwdSet(ctx, userId, req)
	return &v1.TardPwdSetRes{}, err
}

// BindPhone 绑定手机
func (c *Controller) BindPhone(ctx context.Context, req *v1.BindPhoneReq) (res *v1.BindPhoneRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().BindPhone(ctx, userId, req)
	return &v1.BindPhoneRes{}, err
}

// BindEmail 绑定邮箱
func (c *Controller) BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().BindEmail(ctx, userId, req)
	return &v1.BindEmailRes{}, err
}

// UpdateUserAddress 地址静默绑定
func (c *Controller) UpdateUserAddress(ctx context.Context, req *v1.UpdateUserAddressReq) (res *v1.UpdateUserAddressRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().UpdateUserAddress(ctx, userId, req)
	return &v1.UpdateUserAddressRes{}, err
}

// GetUserInfo 获取当前登录用户信息
func (c *Controller) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return user.New().GetUserInfo(ctx, userId)
}

// GetUserAddress 获取已绑定的钱包地址列表
func (c *Controller) GetUserAddress(ctx context.Context, req *v1.GetUserAddressReq) (res *v1.GetUserAddressRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return user.New().GetUserAddress(ctx, userId)
}

// UploadKYC KYC 实名认证
func (c *Controller) UploadKYC(ctx context.Context, req *v1.UploadKYCReq) (res *v1.UploadKYCRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	err = user.New().UploadKYC(ctx, userId, req)
	return &v1.UploadKYCRes{}, err
}

// SendCode 发送验证码
func (c *Controller) SendCode(ctx context.Context, req *v1.SendCodeReq) (res *v1.SendCodeRes, err error) {
	return user.New().SendCode(ctx, req)
}
