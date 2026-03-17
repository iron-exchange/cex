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
	"GoCEX/app/controller/owncoin"
	"GoCEX/app/controller/second_contract"
	"GoCEX/app/controller/trading"
	"GoCEX/app/controller/user"
	"GoCEX/internal/service/market"
	"GoCEX/internal/service/middleware"
	"GoCEX/internal/service/websocket"
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
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Bind(
						userCtrl.Register,
						userCtrl.Login,
						userCtrl.SendCode,
					)
				})

				// 放行: 系统大厅开放读取配置、门户公告内容与 WebSocket
				cmsCtrl := cms.New()
				ossCtrl := oss.New()
				defiCtrl := defi.New()
				owncoinCtrl := owncoin.New()
				miningCtrl := mining.New()
				secondCtrl := second_contract.New()
				group.Bind(
					common.New(),
					cmsCtrl.GetAllNoticeList,
					cmsCtrl.GetHelpCenterList,
					ossCtrl.Upload,
					defiCtrl.GetDefiRate,
					owncoinCtrl.GetOwnCoinList, // 自发币大厅展示无需鉴权
					miningCtrl.MingProductList, // 矿机大厅展示无需鉴权
					secondCtrl.GetCoinList,     // 秒合约币种配置展示无需鉴权
				)

				// 针对 App Frontend 提供的全局 WebSocket 接入端点
				group.ALL("/ws/{uuid}", websocket.AppHub.AppConnect)

				// 针对具体币种行情订阅的 WebSocket 独立端点
				group.ALL("/ws/coin/{userId}", websocket.CoinHub.CoinConnect)

				// 秒合约订单开奖结算的单独定向推送端点
				group.ALL("/webSocket/coinOver/{userId}", websocket.CoinOverHub.CoinOverConnect)

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
					group.Bind(
						defiCtrl.SendApproveHash,
						defiCtrl.ShowIncome,
						defiCtrl.ShowOrder,
					)
					group.Bind(
						miningCtrl.MiningShow,
						miningCtrl.MiningSubmit,
						miningCtrl.MiningRedemption,
						miningCtrl.PersonalIncome,
						miningCtrl.FinancialSubmit,
						miningCtrl.MingOrderList,
						miningCtrl.MingOrderDetail,
						miningCtrl.MingOrderRedemptionNew,
					)
					group.Bind(cmsCtrl.GetUserMail)
					group.Bind(
						owncoinCtrl.GetOwnCoinDetail,
						owncoinCtrl.SubscribeOwnCoin,
					)
					group.Bind(
						secondCtrl.GetCoinDetail,
						secondCtrl.CreateOrder,
						secondCtrl.SelectOrderList,
					)
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
