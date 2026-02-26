package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/user"
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
	u, err := user.New().Login(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2. 签发 JWT (使用 jwt/v4 基础库手动签发，兼容我们的中间件解析)
	expire := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.MapClaims{
		"userId":   u.UserId,
		"username": u.LoginName,
		"exp":      expire.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("cex_secret_key_change_me_later"))

	return &v1.LoginRes{
		Token:    tokenString,
		Expire:   expire.Format("2006-01-02 15:04:05"),
		UserId:   uint64(u.UserId),
		Username: u.LoginName,
	}, nil
}
