package middleware

import (
	"context"
	"time"

	"GoCEX/internal/codes"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	// Auth Auth 中间件实例
	Auth *jwt.GfJWTMiddleware
)

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "cex",
		Key:           []byte("cex_secret_key_change_me_later"),
		Timeout:       time.Hour * 24 * 7, // 7天过期
		MaxRefresh:    time.Hour * 24,     // 1天内可刷新
		IdentityKey:   "userId",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,

		// 这一步在直接调用 TokenGenerator 时是必不可少的，否则 TokenGenerator 内部走不通或者无法给 PayloadFunc 传参
		Authenticator: func(ctx context.Context) (interface{}, error) {
			// 由于我们是在 user 控制器里手动验证后调用的 TokenGenerator(data)
			// TokenGenerator 内部为了补全流程也会调用一次 Authenticator。
			// 解决办法是在业务侧不要直接调 TokenGenerator 而是自己组装，或者在这里给透传的参数放行。
			// 为了防止中间件自带的 login 拦截，这里返回我们在 TokenGenerator 传进来的 map
			r := g.RequestFromCtx(ctx)
			return r.GetCtxVar("jwt_payload").Val(), nil
		},

		// 登录时的 Payload 设置 (这里假定用户登录后返回了用户信息 user)
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					"userId":  v["userId"],
					"account": v["account"],
				}
			}
			return jwt.MapClaims{}
		},

		// 每个带 Token 的请求会走这个方法解析
		IdentityHandler: func(ctx context.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			return claims["userId"]
		},

		// Unauthorized 处理异常返回
		Unauthorized: func(ctx context.Context, code int, message string) {
			r := g.RequestFromCtx(ctx)
			r.Response.WriteJson(g.Map{
				"code": codes.Unauthorized.Code(),
				"msg":  message,
				"data": nil,
			})
			r.ExitAll()
		},
	})

	Auth = auth
}

// 供业务自己二次包装的中间件（例如判断是否黑名单）
func CtxAuth(r *ghttp.Request) {
	// 复用 jwt 的身份校验
	Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
