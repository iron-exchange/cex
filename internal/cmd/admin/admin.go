package admin

import (
	"context"

	"GoCEX/internal/controller/admin/asset"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "admin",
		Usage: "admin",
		Brief: "start admin backend http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server("admin")

			// 统一开启跨域与响应格式化
			s.Use(ghttp.MiddlewareCORS)
			s.Use(middleware.HandlerResponse)

			s.Group("/api/admin/v1", func(group *ghttp.RouterGroup) {
				assetCtrl := asset.New()

				// 资产管理相关路由
				group.Group("/asset", func(group *ghttp.RouterGroup) {
					group.Bind(
						assetCtrl.SubAmount,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
