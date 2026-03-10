package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type sAdminAuth struct{}

func New() *sAdminAuth {
	return &sAdminAuth{}
}

// checkPassword 支持 BCrypt（RuoYi 默认）和 MD5 双模式匹配
func checkPassword(plain, hashed string) bool {
	// BCrypt 格式以 $2a$ 或 $2b$ 开头，长度通常 60 位
	if len(hashed) > 4 && (hashed[:4] == "$2a$" || hashed[:4] == "$2b$") {
		err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
		return err == nil
	}
	// 兼容老库 MD5（32位 hex）
	h := md5.New()
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil)) == hashed
}

// Login 后台管理员登录
func (s *sAdminAuth) Login(ctx context.Context, req *v1.AdminLoginReq) (*v1.AdminLoginRes, error) {
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where("user_name", req.Username).Where("del_flag", "0").Scan(&user)
	if err != nil {
		// 直接暴露真实 DB 错误信息，方便调试
		return nil, gerror.Newf("数据库查询失败: %v", err)
	}

	if user.UserId == 0 {
		return nil, gerror.New("用户不存在或密码错误")
	}
	// 兼容 RuoYi 两种激活状态：有些库用 "0" 正常，有些是空字符串
	// 只要非 "1"(停用) 就放行
	if user.Status == "1" {
		return nil, gerror.New("账户已停用，请联系管理员")
	}
	if !checkPassword(req.Password, user.Password) {
		// 临时调试：输出数据库里存的密码前20位字符（调试完山删除）
		hashPreview := user.Password
		if len(hashPreview) > 20 {
			hashPreview = hashPreview[:20]
		}
		return nil, gerror.Newf("密码错误, DB哈希字头: [%s...]长度:%d", hashPreview, len(user.Password))
	}

	// Sign JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  user.UserId,
		"account": user.UserName, // reuse existing gf-jwt logic if any, else define standalone
	})
	tokenString, err := token.SignedString([]byte("your-256-bit-secret")) // Update with real loaded config
	if err != nil {
		return nil, err
	}

	return &v1.AdminLoginRes{Token: tokenString}, nil
}

// GetInfo 获取管理员权限信息
func (s *sAdminAuth) GetInfo(ctx context.Context, req *v1.AdminGetInfoReq) (*v1.AdminGetInfoRes, error) {
	// Dummy logic: admin gets all
	permissions := []string{"*:*:*"}
	roles := []string{"admin"}
	userMap := g.Map{
		"userId":   1,
		"userName": "admin",
		"nickName": "GoCEX Admin",
		"avatar":   "",
	}
	return &v1.AdminGetInfoRes{
		User:        userMap,
		Roles:       roles,
		Permissions: permissions,
	}, nil
}

// GetRouters 获取动态菜单树
func (s *sAdminAuth) GetRouters(ctx context.Context, req *v1.AdminGetRoutersReq) (*v1.AdminGetRoutersRes, error) {
	// Simple stub for getting Routers, typically derived from sys_menu built into a tree matching user's sys_role.
	// For now, delivering basic mock array to let FE boot.
	var list []v1.AdminRouterInfo
	return &v1.AdminGetRoutersRes{Data: list}, nil
}
