package middleware

import (
	"strings"

	"GoCEX/internal/codes"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

// adminJwtSecret 与 internal/logic/admin/auth/auth.go 中签名 Key 保持完全一致
const adminJwtSecret = "your-256-bit-secret"

// CtxAdminAuth 后台接口拦截哨兵
func CtxAdminAuth(r *ghttp.Request) {
	// 放行登录白名单接口
	path := r.URL.Path
	if path == "/api/admin/v1/login" || path == "/login" || path == "/api/admin/v1/captchaImage" {
		r.Middleware.Next()
		return
	}

	// 从 Header 提取 Bearer Token
	authHeader := r.Header.Get("Authorization")
	tokenString := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	}
	// 也支持 query param: ?token=xxx
	if tokenString == "" {
		tokenString = r.Get("token").String()
	}

	if tokenString == "" {
		r.Response.WriteJson(g.Map{
			"code": codes.Unauthorized.Code(),
			"msg":  "认证失败，Token 不能为空",
			"data": nil,
		})
		r.ExitAll()
		return
	}

	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(adminJwtSecret), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteJson(g.Map{
			"code": codes.Unauthorized.Code(),
			"msg":  "认证失败，Token 无效或已过期",
			"data": nil,
		})
		r.ExitAll()
		return
	}

	// 将 claims 注入 ctx，供后续业务层取用
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		r.SetCtxVar("adminId", claims["userId"])
		r.SetCtxVar("adminAccount", claims["account"])
	}

	r.Middleware.Next()
}
