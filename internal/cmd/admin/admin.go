package admin

import (
	"context"

	"GoCEX/admin/controller/address"
	"GoCEX/admin/controller/announcement"
	"GoCEX/admin/controller/auth"
	"GoCEX/admin/controller/bank"
	"GoCEX/admin/controller/bot"
	"GoCEX/admin/controller/collection"
	"GoCEX/admin/controller/contract"
	"GoCEX/admin/controller/currency_trading"
	"GoCEX/admin/controller/dashboard"
	"GoCEX/admin/controller/defi"
	"GoCEX/admin/controller/exchange"
	"GoCEX/admin/controller/financial"
	"GoCEX/admin/controller/funding"
	"GoCEX/admin/controller/ieo"
	"GoCEX/admin/controller/kyc"
	"GoCEX/admin/controller/loan"
	"GoCEX/admin/controller/log"
	"GoCEX/admin/controller/mining"
	"GoCEX/admin/controller/monitor"
	"GoCEX/admin/controller/rbac"
	"GoCEX/admin/controller/report"
	"GoCEX/admin/controller/second_contract"
	"GoCEX/admin/controller/swap"
	"GoCEX/admin/controller/system"
	"GoCEX/admin/controller/user"
	"GoCEX/admin/controller/wallet_record"
	"GoCEX/internal/controller/admin/asset"
	"GoCEX/internal/service/middleware"
	"GoCEX/internal/service/websocket"

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

			// 很多前端默认请求放在 / 或者 /dev-api 下，这里将登录相关绑定在根路由下，或者 v1 中皆可，保持兼容
			authCtrl := auth.New()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(authCtrl)

				// 兼容原版 /webSocket/notice/{userId}/{uuId} 消息推送路由 (双路绑定，防止前端没改)
				group.ALL("/admin/webSocket/notice/{userId}/{uuId}", websocket.NoticeHub.Connect)
				group.ALL("/webSocket/notice/{userId}/{uuId}", websocket.NoticeHub.Connect)
			})

			s.Group("/api/admin/v1", func(group *ghttp.RouterGroup) {
				// JWT 阻断网关
				group.Middleware(middleware.CtxAdminAuth)
				assetCtrl := asset.New()
				userCtrl := user.New()
				kycCtrl := kyc.New()
				addressCtrl := address.New()
				bankCtrl := bank.New()
				logCtrl := log.New()
				collectionCtrl := collection.New()
				walletRecordCtrl := wallet_record.New()
				reportCtrl := report.New()
				dashboardCtrl := dashboard.New()
				fundingCtrl := funding.New()
				ieoCtrl := ieo.New()
				secondContractCtrl := second_contract.New()
				exchangeCtrl := exchange.New()
				contractCtrl := contract.New()
				currencyTradingCtrl := currency_trading.New()
				miningCtrl := mining.New()
				defiCtrl := defi.New()
				swapCtrl := swap.New()
				botCtrl := bot.New()
				financialCtrl := financial.New()
				announcementCtrl := announcement.New()
				loanCtrl := loan.New()
				systemCtrl := system.New()
				rbacCtrl := rbac.New()
				monitorCtrl := monitor.New()

				// 如果前端也带前缀请求 Auth，也同时挂载一份在这里
				group.Bind(
					authCtrl,
					assetCtrl,
					userCtrl,
					kycCtrl,
					addressCtrl,
					bankCtrl,
					logCtrl,
					collectionCtrl,
					walletRecordCtrl,
					reportCtrl,
					dashboardCtrl,
					fundingCtrl,
					ieoCtrl,
					secondContractCtrl,
					exchangeCtrl,
					contractCtrl,
					currencyTradingCtrl,
					miningCtrl,
					defiCtrl,
					swapCtrl,
					botCtrl,
					financialCtrl,
					announcementCtrl,
					loanCtrl,
					systemCtrl,
					rbacCtrl,
					monitorCtrl,
				)
			})

			s.Run()
			return nil
		},
	}
)
