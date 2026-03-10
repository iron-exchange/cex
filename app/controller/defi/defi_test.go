package defi_test

import (
	"context"
	"testing"
	"time"

	v1 "GoCEX/api/app/v1"
	"GoCEX/app/controller/defi"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func Test_DefiController(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(gconv.String(time.Now().UnixNano()))

		// 统一处理返回 JSON
		s.Use(middleware.HandlerResponse)

		// 注入路由
		s.Group("/api", func(group *ghttp.RouterGroup) {
			// 公开路由 (GetDefiRate)
			group.Group("/public", func(pubGroup *ghttp.RouterGroup) {
				pubGroup.Bind(defi.New().GetDefiRate)
			})

			// 鉴权路由
			group.Group("/auth", func(authGroup *ghttp.RouterGroup) {
				authGroup.Middleware(middleware.CtxAuth)
				defiCtrl := defi.New()
				authGroup.Bind(
					defiCtrl.SendApproveHash,
					defiCtrl.ShowIncome,
					defiCtrl.ShowOrder,
				)
			})
		})

		s.SetDumpRouterMap(false)
		s.SetPort(0) // Use random available port
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond) // wait for server start

		client := g.Client()
		client.SetPrefix("http://127.0.0.1:" + gconv.String(s.GetListenedPort()))

		// 生成本地测试用的 JWT Token (模拟用户 ID = 1)
		// 强制生成一个 Payload 给 Token
		tokenString, _, err := middleware.Auth.TokenGenerator(g.Map{"userId": uint64(1)})
		t.AssertNil(err)

		// 将 Token 设置进通用头部
		client.SetHeader("Authorization", "Bearer "+tokenString)

		// GetDefiRate
		{
			resp, err := client.Post(context.TODO(), "/api/public/apiDefi/getDefiRate")
			t.AssertNil(err)
			defer resp.Close()
			t.Assert(resp.StatusCode, 200)

			t.Log("GetDefiRate Response:", resp.ReadAllString())
		}

		// SendApproveHash
		{
			resp, err := client.Post(context.TODO(), "/api/auth/apiDefi/sendApproveHash", &v1.SendApproveHashReq{
				Address: "0xTestAddress123",
				Hash:    "0xTestHash456",
			})
			t.AssertNil(err)
			defer resp.Close()
			t.Assert(resp.StatusCode, 200)

			t.Log("SendApproveHash Response:", resp.ReadAllString())
		}

		// ShowIncome
		{
			resp, err := client.Post(context.TODO(), "/api/auth/apiDefi/showIncome", &v1.ShowIncomeReq{})
			t.AssertNil(err)
			defer resp.Close()
			t.Assert(resp.StatusCode, 200)

			t.Log("ShowIncome Response:", resp.ReadAllString())
		}

		// ShowOrder
		{
			resp, err := client.Get(context.TODO(), "/api/auth/apiDefi/showOrder")
			t.AssertNil(err)
			defer resp.Close()
			t.Assert(resp.StatusCode, 200)

			t.Log("ShowOrder Response:", resp.ReadAllString())
		}
	})
}
