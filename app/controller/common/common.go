package common

import (
	"context"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/common"

	"github.com/gogf/gf/v2/frame/g"
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

	// 启动一个独立协程，监听 Redis 的行情广播频道
	go func() {
		defer ws.Close()

		// 假设原系统向 Redis 发布行情的频道是 MARKET:TICKER
		conn, err := g.Redis().Conn(ctx)
		if err != nil {
			g.Log().Error(ctx, "无法连接 Redis 行情频道:", err)
			return
		}
		defer conn.Close(ctx)

		// 订阅行情频道
		_, err = conn.Subscribe(ctx, "MARKET:TICKER")
		if err != nil {
			g.Log().Error(ctx, "订阅行情频道 MARKET:TICKER 失败:", err)
			return
		}

		// 死循环阻塞读取广播消息
		for {
			msg, err := conn.ReceiveMessage(ctx)
			if err != nil {
				// Redis 故障或断开，跳出循环
				g.Log().Error(ctx, "Redis ReceiveMessage 错误:", err)
				break
			}

			// 将拿到的真实行情 JSON 字符串原封不动推送给连进来的前端 WebSocket
			if err := ws.WriteMessage(1, []byte(msg.Payload)); err != nil {
				// 前端断开连接，发送失败直接退出推送协程
				break
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
