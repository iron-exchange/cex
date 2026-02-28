package user

import (
	"context"
	"errors"

	v1 "GoCEX/app/api"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	// 目前通过 goframe 的设计模式，在这里抛出实例注册给 service 层
	// 如果你还没用 `gf gen service`，我们可以先起一个全局变量供 Controller 直接调用
	// 等写完了用 gf gen service 生成接口
}

// New 创建 User 服务的实例
func New() *sUser {
	return &sUser{}
}

// verifyCode 内部通用方法: 校验验证码正确性并销毁
func (s *sUser) verifyCode(ctx context.Context, scene, accountType, account, code string) error {
	if code == "" {
		return gerror.NewCode(codes.ClientError, "请输入验证码")
	}
	redisKey := "CEX:VERIFY_CODE:" + scene + ":" + accountType + ":" + account
	val, err := g.Redis().Do(ctx, "GET", redisKey)
	if err != nil || val.IsEmpty() {
		return gerror.NewCode(codes.ClientError, "验证码已过期或不存在")
	}
	if val.String() != code {
		return gerror.NewCode(codes.ClientError, "验证码输入错误")
	}
	// 验证成功后立即销毁，防止重播攻击
	_, _ = g.Redis().Do(ctx, "DEL", redisKey)
	return nil
}

// Register 注册新用户
func (s *sUser) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterRes, error) {
	var column string
	var accountVal string

	switch in.SignType {
	case "LOGIN":
		column = dao.AppUser.Columns().LoginName
		accountVal = in.LoginName
	case "PHONE":
		column = dao.AppUser.Columns().Phone
		accountVal = in.Phone
	case "EMAIL":
		column = dao.AppUser.Columns().Email
		accountVal = in.Email
	case "ADDRESS":
		column = dao.AppUser.Columns().Address
		accountVal = in.Address
	default:
		return nil, gerror.NewCode(codes.ClientError, "不支持的账户类型")
	}

	if accountVal == "" {
		return nil, gerror.NewCode(codes.ClientError, "账号内容不能为空")
	}

	// 触发手机或邮箱的验证码校验机制 (注册场景)
	if in.SignType == "PHONE" || in.SignType == "EMAIL" {
		if errVerify := s.verifyCode(ctx, "REGISTER", in.SignType, accountVal, in.Code); errVerify != nil {
			return nil, errVerify
		}
	}

	// 校验账号是否已存在
	count, err := dao.AppUser.Ctx(ctx).Where(column, accountVal).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.Failed, "该账号已存在")
	}

	// 密码加密 (Web3 / ADDRESS 注册时通常没有传统密码，这里对非Web3设置默认密码或者拿入参)
	var hashedPassword string
	if in.SignType != "ADDRESS" {
		pwd := in.LoginPassword
		if pwd == "" {
			pwd = "123456" // 原Java代码逻辑：默认123456
		}
		salt := "cex_salt"
		hashedPassword, _ = gmd5.EncryptString(pwd + salt)
	}

	// 处理邀请码与代理层级关系
	var appParentIds string
	var adminParentIds string
	if in.InviteCode != "" {
		var inviter entity.AppUser
		err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().ActiveCode, in.InviteCode).Scan(&inviter)
		if err == nil && inviter.UserId > 0 {
			appParentIds = inviter.AppParentIds + "," + gconv.String(inviter.UserId)
			adminParentIds = inviter.AdminParentIds
		}
	}

	// 生成当前用户自己的邀请码
	newActiveCode := gconv.String(g.Redis().MustDo(ctx, "INCR", "CEX:GLOBAL:INVITE_CODE").Int64() + 10000)

	// 组装入库实体
	userMap := g.Map{
		dao.AppUser.Columns().LoginPassword:  hashedPassword,
		dao.AppUser.Columns().ActiveCode:     newActiveCode,
		dao.AppUser.Columns().AppParentIds:   appParentIds,
		dao.AppUser.Columns().AdminParentIds: adminParentIds,
		dao.AppUser.Columns().Status:         0, // 默认正常
		dao.AppUser.Columns().IsFreeze:       "1",
		dao.AppUser.Columns().IsBlack:        1,
		dao.AppUser.Columns().Level:          0,
		dao.AppUser.Columns().IsTest:         0,
		dao.AppUser.Columns().TotleAmont:     0,
		dao.AppUser.Columns().RechargeAmont:  0,
		dao.AppUser.Columns().Buff:           0,
		dao.AppUser.Columns().LoginName:      "",
		dao.AppUser.Columns().Phone:          "",
		dao.AppUser.Columns().Email:          "",
		dao.AppUser.Columns().Address:        "",
		dao.AppUser.Columns().WalletType:     "",
	}

	// 绑定对应账号字段
	switch in.SignType {
	case "LOGIN":
		userMap[dao.AppUser.Columns().LoginName] = accountVal
	case "PHONE":
		userMap[dao.AppUser.Columns().Phone] = accountVal
	case "EMAIL":
		userMap[dao.AppUser.Columns().Email] = accountVal
	case "ADDRESS":
		userMap[dao.AppUser.Columns().Address] = accountVal
		userMap[dao.AppUser.Columns().WalletType] = "ETH" // 默认为以太坊体系
	}

	// 执行入库
	insertID, err := dao.AppUser.Ctx(ctx).Data(userMap).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &v1.RegisterRes{
		UserId:    uint64(insertID),
		SignType:  in.SignType,
		LoginName: in.LoginName,
		Phone:     in.Phone,
		Email:     in.Email,
		Address:   in.Address,
	}, nil
}

// Login 校验用户登录并返回对应信息给中间件签发 JWT
func (s *sUser) Login(ctx context.Context, in *v1.LoginReq) (*entity.AppUser, string, error) {
	var column string
	var accountVal string
	switch in.SignType {
	case "LOGIN":
		column = dao.AppUser.Columns().LoginName
		accountVal = in.LoginName
	case "PHONE":
		column = dao.AppUser.Columns().Phone
		accountVal = in.Phone
	case "EMAIL":
		column = dao.AppUser.Columns().Email
		accountVal = in.Email
	case "ADDRESS":
		column = dao.AppUser.Columns().Address
		accountVal = in.Address
	default:
		return nil, "", gerror.NewCode(codes.ClientError, "不支持的账户类型")
	}

	if accountVal == "" {
		return nil, "", gerror.NewCode(codes.ClientError, "账号内容不能为空")
	}

	// 触发手机或邮箱的验证码登录校验 (场景: 传了验证码则走验证码登录)
	if (in.SignType == "PHONE" || in.SignType == "EMAIL") && in.Code != "" {
		if errVerify := s.verifyCode(ctx, "LOGIN", in.SignType, accountVal, in.Code); errVerify != nil {
			return nil, "", errVerify
		}
	}

	var user entity.AppUser
	err := dao.AppUser.Ctx(ctx).Where(column, accountVal).Scan(&user)
	if err != nil {
		return nil, "", gerror.NewCode(codes.UserNotFound, "用户不存在")
	}

	// 判断封禁状态
	if user.Status == 1 || user.IsFreeze == "2" {
		return nil, "", gerror.NewCode(codes.UserDisabled, "账号已被冻结")
	}

	// ADDRESS 登录无需密码比对
	if in.SignType != "ADDRESS" {
		pwd := in.LoginPassword
		if pwd == "" {
			pwd = "123456"
		}
		salt := "cex_salt"
		hashedPassword, _ := gmd5.EncryptString(pwd + salt)
		if user.LoginPassword != hashedPassword {
			return nil, "", errors.New("账号或密码错误")
		}
	}

	return &user, accountVal, nil
}

// PwdSett 修改登录密码
func (s *sUser) PwdSett(ctx context.Context, userId uint64, req *v1.PwdSettReq) error {
	salt := "cex_salt"
	hashedPassword, _ := gmd5.EncryptString(req.Pwd + salt)

	// 更新登录密码
	_, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).
		Data(g.Map{dao.AppUser.Columns().LoginPassword: hashedPassword}).Update()
	return err
}

// TardPwdSet 设置资金密码 (存在 app_user_detail)
func (s *sUser) TardPwdSet(ctx context.Context, userId uint64, req *v1.TardPwdSetReq) error {
	salt := "cex_salt"
	hashedPassword, _ := gmd5.EncryptString(req.Pwd + salt)

	count, err := dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		_, err = dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).
			Data(g.Map{dao.AppUserDetail.Columns().UserTardPwd: hashedPassword}).Update()
	} else {
		_, err = dao.AppUserDetail.Ctx(ctx).Data(g.Map{
			dao.AppUserDetail.Columns().Id:          userId, // 修复无自增序列报错
			dao.AppUserDetail.Columns().UserId:      userId,
			dao.AppUserDetail.Columns().UserTardPwd: hashedPassword,
		}).Insert()
	}
	return err
}

// BindPhone 绑定手机
func (s *sUser) BindPhone(ctx context.Context, userId uint64, req *v1.BindPhoneReq) error {
	// 验证短信验证码
	if errVerify := s.verifyCode(ctx, "BIND", "PHONE", req.Phone, req.Code); errVerify != nil {
		return errVerify
	}

	// 查重该手机号是否被其他账号绑定
	count, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().Phone, req.Phone).
		WhereNot(dao.AppUser.Columns().UserId, userId).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.NewCode(codes.Failed, "该手机号已被绑定")
	}

	_, err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).
		Data(g.Map{dao.AppUser.Columns().Phone: req.Phone}).Update()
	return err
}

// BindEmail 绑定邮箱
func (s *sUser) BindEmail(ctx context.Context, userId uint64, req *v1.BindEmailReq) error {
	// 验证邮箱验证码
	if errVerify := s.verifyCode(ctx, "BIND", "EMAIL", req.Email, req.EmailCode); errVerify != nil {
		return errVerify
	}

	// 查重该邮箱是否被其他账号绑定
	count, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().Email, req.Email).
		WhereNot(dao.AppUser.Columns().UserId, userId).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.NewCode(codes.Failed, "该邮箱已被绑定")
	}

	_, err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).
		Data(g.Map{dao.AppUser.Columns().Email: req.Email}).Update()
	return err
}

// UpdateUserAddress 地址静默绑定
func (s *sUser) UpdateUserAddress(ctx context.Context, currentUserId uint64, req *v1.UpdateUserAddressReq) error {
	targetUserId := currentUserId

	// 1. 更新主表冗余 (快速出账用)
	_, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, targetUserId).
		Data(g.Map{
			dao.AppUser.Columns().Address:    req.Address,
			dao.AppUser.Columns().WalletType: req.Type,
		}).Update()
	if err != nil {
		return err
	}

	// 2. 详细记录到 app_user_address 子表
	count, err := dao.AppUserAddress.Ctx(ctx).
		Where(dao.AppUserAddress.Columns().UserId, targetUserId).
		Where(dao.AppUserAddress.Columns().Symbol, req.Type).Count()
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = dao.AppUserAddress.Ctx(ctx).
			Where(dao.AppUserAddress.Columns().UserId, targetUserId).
			Where(dao.AppUserAddress.Columns().Symbol, req.Type).
			Data(g.Map{dao.AppUserAddress.Columns().Address: req.Address}).Update()
	} else {
		_, err = dao.AppUserAddress.Ctx(ctx).Data(g.Map{
			dao.AppUserAddress.Columns().UserId:  targetUserId,
			dao.AppUserAddress.Columns().Symbol:  req.Type,
			dao.AppUserAddress.Columns().Address: req.Address,
		}).Insert()
	}
	return err
}

// GetUserAddress 获取用户绑定的所有钱包地址
func (s *sUser) GetUserAddress(ctx context.Context, currentUserId uint64) (*v1.GetUserAddressRes, error) {
	var addrs []entity.AppUserAddress
	err := dao.AppUserAddress.Ctx(ctx).
		Where(dao.AppUserAddress.Columns().UserId, currentUserId).
		Scan(&addrs)

	if err != nil {
		return nil, err
	}

	res := &v1.GetUserAddressRes{
		List: make([]v1.UserAddressInfo, 0, len(addrs)),
	}

	for _, addr := range addrs {
		res.List = append(res.List, v1.UserAddressInfo{
			Address: addr.Address,
			Type:    addr.Symbol,
		})
	}

	return res, nil
}

// UploadKYC 实名认证上传
func (s *sUser) UploadKYC(ctx context.Context, userId uint64, req *v1.UploadKYCReq) error {
	// 组装 detail 表数据
	data := g.Map{
		dao.AppUserDetail.Columns().RealName:  req.RealName,
		dao.AppUserDetail.Columns().IdCard:    req.IdCard,
		dao.AppUserDetail.Columns().FrontUrl:  req.FrontUrl,
		dao.AppUserDetail.Columns().BackUrl:   req.BackUrl,
		dao.AppUserDetail.Columns().HandelUrl: req.HandelUrl,
		dao.AppUserDetail.Columns().Country:   req.Country,
		dao.AppUserDetail.Columns().CardType:  req.CardType,
		// 设置审核初级状态为待审核 (这里暂定 1: 待审核)
		dao.AppUserDetail.Columns().AuditStatusPrimary: 1,
	}

	// Flag 为 2 时可能强制同步送审高级认证
	if req.Flag == "2" {
		data[dao.AppUserDetail.Columns().AuditStatusAdvanced] = 1
	}

	count, err := dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).Count()
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).Data(data).Update()
	} else {
		data[dao.AppUserDetail.Columns().Id] = userId // 修复无自增序列报错
		data[dao.AppUserDetail.Columns().UserId] = userId
		_, err = dao.AppUserDetail.Ctx(ctx).Data(data).Insert()
	}

	return err
}

// SendCode 发送验证码 (短信/邮箱)
func (s *sUser) SendCode(ctx context.Context, in *v1.SendCodeReq) (*v1.SendCodeRes, error) {
	// 简单的防刷：根据场景 (Scene)、类型 (Type) 和目标 (To) 作为 Redis Key
	redisKey := "CEX:VERIFY_CODE:" + in.Scene + ":" + in.Type + ":" + in.To
	exists, _ := g.Redis().Do(ctx, "EXISTS", redisKey)
	if exists.Int() == 1 {
		return nil, gerror.NewCode(codes.ClientError, "请求过于频繁，请稍后再试")
	}

	// 随机生成 6 位数字验证码
	code := grand.Digits(6)

	// 存入 Redis，5分钟 (300秒) 有效
	_, err := g.Redis().Do(ctx, "SETEX", redisKey, 300, code)
	if err != nil {
		return nil, gerror.NewCode(codes.Failed, "验证码下发失败，请重试")
	}

	// ----- 模拟发送动作 -----
	// 实际生产中，这里应该调用阿里云短信/腾讯云短信/AWS邮件服务等 SDK 来发送 `code`。
	// 目前我们将其“发送”到服务端的控制台日志中，您可以通过查看终端日志来获取收到的验证码。
	g.Log().Infof(ctx, "=======> [%s] 向 %s 发送 %s 验证码: %s <=======", in.Scene, in.To, in.Type, code)

	// 绝对不能把 Code 放在 Res 里返回给前端，防止被抓包越权！
	return &v1.SendCodeRes{}, nil
}

// GetUserInfo 获取当前登录用户的基础与实名等全量信息
func (s *sUser) GetUserInfo(ctx context.Context, userId uint64) (*v1.GetUserInfoRes, error) {
	var user entity.AppUser
	err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).Scan(&user)
	if err != nil {
		return nil, gerror.NewCode(codes.ClientError, "用户不存在或已被删除")
	}

	var detail entity.AppUserDetail
	// 可能有很多老用户尚未记录 detail 表，所以不报错跳过，用默认空值即可
	_ = dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).Scan(&detail)

	return &v1.GetUserInfoRes{
		UserId:              user.UserId,
		LoginName:           user.LoginName,
		Phone:               user.Phone,
		Email:               user.Email,
		Address:             user.Address,
		WalletType:          user.WalletType,
		Level:               user.Level,
		Status:              user.Status,
		IsFreeze:            user.IsFreeze,
		RealName:            detail.RealName,
		AuditStatusPrimary:  detail.AuditStatusPrimary,
		AuditStatusAdvanced: detail.AuditStatusAdvanced,
	}, nil
}
