package app

import (
	"context"

	"GoCEX/app/controller/asset"
	"GoCEX/app/controller/cms"
	"GoCEX/app/controller/common"
	"GoCEX/app/controller/defi"
	"GoCEX/app/controller/funding"
	"GoCEX/app/controller/mining"
	"GoCEX/app/controller/oss"
	"GoCEX/app/controller/trading"
	"GoCEX/app/controller/user"
	"GoCEX/internal/service/market"
	"GoCEX/internal/service/middleware"
	taskCtrl "GoCEX/task/controller/task"

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

			// 启动币安 WebSocket 后台行情拉取服务 (取代老的 Http Polling Cron)
			market.StartBinanceWSDaemon(ctx)

			// 统一开启跨域
			s.Use(ghttp.MiddlewareCORS)

			// 全局响应的 JSON 包装器
			s.Use(middleware.HandlerResponse)

			// V1 接口 API 路由注册
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				// 注入相关 Controller
				userCtrl := user.New()

				// 放行: 注册、登录、与优盾第三方支付回调 Webhook
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						userCtrl.Register,
						userCtrl.Login,
						userCtrl.SendCode,
					)
				})

				// 放行: 系统大厅开放读取配置、门户公告内容与 WebSocket
				cmsCtrl := cms.New()
				ossCtrl := oss.New()
				group.Bind(
					common.New(),
					cmsCtrl.GetAllNoticeList,
					cmsCtrl.GetHelpCenterList,
					ossCtrl.Upload,
					defi.New().GetDefiRate,
				)

				// 需要鉴权的接口 (包含充提、交易、与用户设置)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.CtxAuth)

					group.Group("/user", func(userGroup *ghttp.RouterGroup) {
						userGroup.Bind(
							userCtrl.PwdSett,
							userCtrl.TardPwdSet,
							userCtrl.BindPhone,
							userCtrl.BindEmail,
							userCtrl.UpdateUserAddress,
							userCtrl.GetUserAddress,
							userCtrl.GetUserInfo,
							userCtrl.UploadKYC,
						)
					})

					group.Bind(funding.New())
					group.Bind(trading.New())
					group.Bind(asset.New())
					group.Bind(defi.New())
					group.Bind(mining.New())
					group.Bind(cmsCtrl.GetUserMail)
				})
			})

			// 挂载独立的后台定时器 (Quartz Tasks) 触发网关
			// 这些接口本身并不属于面向 C 端的 /api 业务，所以注册为隔离的 /task 域名后缀
			// 未来可以在此处新增一层特殊的 Token Validator 或只允许 localhost 回环地址调用的风控层。
			s.Group("/task", func(taskGroup *ghttp.RouterGroup) {
				taskGroup.Middleware(ghttp.MiddlewareHandlerResponse) // 返回标准 JSON 化
				taskGroup.Bind(
					taskCtrl.New(),
				)
			})

			s.Run()
			return nil
		},
	}
)
