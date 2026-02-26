package app

import (
	"context"

	"GoCEX/internal/controller/app/user"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var (
	Main = gcmd.Command{
		Name:  "app",
		Usage: "app",
		Brief: "start app http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 统一开启跨域
			s.Use(ghttp.MiddlewareCORS)

			// 全局响应的 JSON 包装器
			s.Use(middleware.HandlerResponse)

			// V1 接口 API 路由注册
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				// 注入用户相关 Controller
				userCtrl := user.New()

				// 放行: 注册与登录
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						userCtrl.Register,
						userCtrl.Login,
					)
				})

				// 需要鉴权的接口，放在下面
				// group.Group("/", func(group *ghttp.RouterGroup) {
				// 	group.Middleware(middleware.CtxAuth)
				// 	// 例如：group.Bind(userCtrl.Info)
				// })
			})

			s.Run()
			return nil
		},
	}
)
