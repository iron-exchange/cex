package user

import (
	"context"
	"errors"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

// Register 注册新用户
func (s *sUser) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterRes, error) {
	// 校验用户名是否已存在
	count, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().LoginName, in.Username).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.NewCode(codes.Failed, "用户名已存在")
	}

	// 密码加密 (这里模拟原 Java 系统的简单 md5 或者你可以换成 bcrypt)
	salt := "cex_salt" // 生产环境请放在配置文件
	hashedPassword, _ := gmd5.EncryptString(in.Password + salt)

	// 处理邀请码与代理层级关系
	var appParentIds string
	var adminParentIds string
	if in.InviteCode != "" {
		// 查询邀请人
		var inviter entity.AppUser
		err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().ActiveCode, in.InviteCode).Scan(&inviter)
		if err == nil && inviter.UserId > 0 {
			// 将邀请人的链路继承下来并附加上邀请人自己的 ID
			appParentIds = inviter.AppParentIds + "," + gconv.String(inviter.UserId)
			adminParentIds = inviter.AdminParentIds // 如果有后台代理链路，同理附加
		}
	}

	// 生成当前用户自己的邀请码 (简单用时间戳或随机字符串，这里简化处理)
	newActiveCode := gconv.String(g.Redis().MustDo(ctx, "INCR", "CEX:GLOBAL:INVITE_CODE").Int64() + 10000)

	// 组装入库实体
	user := entity.AppUser{
		LoginName:      in.Username,
		LoginPassword:  hashedPassword,
		ActiveCode:     newActiveCode,
		AppParentIds:   appParentIds,
		AdminParentIds: adminParentIds,
		Status:         0, // 默认正常
		IsFreeze:       "1",
		IsBlack:        1,
		Level:          0,
	}

	// 执行入库
	insertID, err := dao.AppUser.Ctx(ctx).Data(user).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &v1.RegisterRes{
		UserId:   uint64(insertID),
		Username: in.Username,
	}, nil
}

// Login 校验用户登录并返回对应信息给中间件签发 JWT
func (s *sUser) Login(ctx context.Context, in *v1.LoginReq) (*entity.AppUser, error) {
	var user entity.AppUser
	err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().LoginName, in.Username).Scan(&user)
	if err != nil {
		// 数据库错误或记录不存在
		return nil, gerror.NewCode(codes.UserNotFound, "用户不存在")
	}

	// 判断封禁状态
	if user.Status == 1 || user.IsFreeze == "2" {
		return nil, gerror.NewCode(codes.UserDisabled, "账号已被冻结")
	}

	// 密码校验
	salt := "cex_salt"
	hashedPassword, _ := gmd5.EncryptString(in.Password + salt)
	if user.LoginPassword != hashedPassword {
		return nil, errors.New("用户名或密码错误")
	}

	return &user, nil
}
