package common

import (
	"context"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/common"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/lib/pq"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Config 获取系统大厅配置
func (c *Controller) Config(ctx context.Context, req *v1.CommonConfigReq) (res *v1.CommonConfigRes, err error) {
	return common.New().GetConfig(ctx)
} // GetAllSetting 获取系统全站配置参数
func (c *Controller) GetAllSetting(ctx context.Context, req *v1.GetAllSettingReq) (res *v1.GetAllSettingRes, err error) {
	return common.New().GetAllSetting(ctx)
}

// GetAppSidebarSetting 获取侧边栏显示的币种
func (c *Controller) GetAppSidebarSetting(ctx context.Context, req *v1.GetAppSidebarSettingReq) (res *v1.GetAppSidebarSettingRes, err error) {
	return common.New().GetAppSidebarSetting(ctx)
}

// GetHomeCoinSetting 获取首页主推显示的币种
func (c *Controller) GetHomeCoinSetting(ctx context.Context, req *v1.GetHomeCoinSettingReq) (res *v1.GetHomeCoinSettingRes, err error) {
	return common.New().GetHomeCoinSetting(ctx)
}

// GetAppCurrencyList 获取充值的通道与开关列表
func (c *Controller) GetAppCurrencyList(ctx context.Context, req *v1.GetAppCurrencyListReq) (res *v1.GetAppCurrencyListRes, err error) {
	return common.New().GetAppCurrencyList(ctx)
}

// GetWithDrawCoinList 获取提现的通道与手续费列表
func (c *Controller) GetWithDrawCoinList(ctx context.Context, req *v1.GetWithDrawCoinListReq) (res *v1.GetWithDrawCoinListRes, err error) {
	return common.New().GetWithDrawCoinList(ctx)
}

// MarketTickerWs 行情推送 WebSocket 预留口
func (c *Controller) MarketTickerWs(ctx context.Context, req *v1.MarketTickerWsReq) (res *v1.MarketTickerWsRes, err error) {
	r := g.RequestFromCtx(ctx)
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 启动一个独立协程，监听 Postgres 的行情广播频道 (替代原本的 Redis)
	go func() {
		defer ws.Close()

		// 获取 PG 连接配置用于 Listener
		dbConfig := g.DB().GetConfig()
		conninfo := "postgres://" + dbConfig.User + ":" + dbConfig.Pass + "@" + dbConfig.Host + ":" + dbConfig.Port + "/" + dbConfig.Name + "?sslmode=disable"

		reportProblem := func(ev pq.ListenerEventType, err error) {
			if err != nil {
				g.Log().Error(ctx, "PG Listener 发生异常:", err)
			}
		}

		listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
		err = listener.Listen("market_ticker")
		if err != nil {
			g.Log().Error(ctx, "PG Listener 订阅失败:", err)
			return
		}
		defer listener.Close()

		g.Log().Info(ctx, "✅ 成功开启 PG LISTEN market_ticker 侦听")

		// 死循环阻塞读取 PG 触发的 NOTIFY 广播消息
		for {
			select {
			case n := <-listener.Notify:
				if n == nil {
					continue
				}
				// 将拿到的真实行情 JSON 字符串原封不动推送给连进来的前端 WebSocket
				g.Log().Info(ctx, "📥 WS收到PG广播:", n.Extra)
				if err := ws.WriteMessage(1, []byte(n.Extra)); err != nil {
					g.Log().Error(ctx, "📤 WS下发前端失败, 客户端已断开:", err)
					return // 前端断开连接，退出推送协程
				}
			case <-time.After(90 * time.Second):
				// 闲置超时或者心跳检活，如果需要ping PG
				go listener.Ping()
			}
		}
	}()

	// 保持主线程阻塞读取客户端发来的消息，例如 ping 包实现心跳保活
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return nil, err // 客户端断网或主动 Close，跳出循环销毁这组连接
		}

		// Echo 测试 / 返回心跳
		if string(msg) == "ping" {
			_ = ws.WriteMessage(msgType, []byte("pong"))
		}
	}
}
