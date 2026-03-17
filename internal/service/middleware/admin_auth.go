package middleware

import (
	"context"
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
	if path == "/api/admin/v1/login" || path == "/login" || path == "/api/admin/v1/captchaImage" || path == "/api/admin/v1/common/getAllSetting" {
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
		uid := claims["userId"]
		account := claims["account"]
		g.Log().Debugf(r.Context(), "[CtxAdminAuth] JWT 解析成功: userId=%v(%T), account=%v", uid, uid, account)

		// GF 原生的 SetCtxVar 是存在 Request 里的，逻辑层用 ctx.Value 拿不到
		// 这里必须显式注入到 Standard Context
		newCtx := context.WithValue(r.Context(), "adminId", uid)
		newCtx = context.WithValue(newCtx, "adminAccount", account)
		r.SetCtx(newCtx)

		// 为了兼容之前的代码，SetCtxVar 也留着
		r.SetCtxVar("adminId", uid)
		r.SetCtxVar("adminAccount", account)

		g.Log().Debugf(r.Context(), "[CtxAdminAuth] Context 已注入: adminId=%v", r.Context().Value("adminId"))
	} else {
		g.Log().Error(r.Context(), "[CtxAdminAuth] JWT Claims 类型断言失败")
	}

	r.Middleware.Next()
}
