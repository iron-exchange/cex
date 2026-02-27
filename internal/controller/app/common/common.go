package common

import (
	"context"

	v1 "GoCEX/api/app/v1"
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
}

// MarketTickerWs 行情推送 WebSocket 预留口
func (c *Controller) MarketTickerWs(ctx context.Context, req *v1.MarketTickerWsReq) (res *v1.MarketTickerWsRes, err error) {
	r := g.RequestFromCtx(ctx)
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 开始死循环监听与推送 (这里仅做架构预留，实际应起 Goroutine 订阅 Redis 的 Price Channel)
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return nil, err // 客户端断网跳出
		}
		// Echo 测试
		if string(msg) == "ping" {
			_ = ws.WriteMessage(msgType, []byte("pong"))
		}
	}
}
