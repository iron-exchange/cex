package websocket

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AppHub 用于承载 App 端的 WebSocket 客户端并提供行情或通知推送功能
var (
	AppHub = &wsHub{
		clients: make(map[*client]bool),
	}
)

func init() {
	// 复用 wsHub 结构体，启动专属的 App 前端死连接监测
	go AppHub.pingRoutine()
}

// AppConnect 承接 App 前端基于 /api/v1/ws/{uuid} 定义的接入
func (h *wsHub) AppConnect(r *ghttp.Request) {
	uuid := r.Get("uuid").String()
	if uuid == "" {
		g.Log().Warning(r.Context(), "App WebSocket 引流缺少 uuid 参数")
	}

	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), "App WebSocket 升级失败:", err)
		return
	}

	c := &client{
		conn:       ws,
		UserId:     "app_user", // 如有鉴权可从 Header/Token 取，目前以公共行情为主
		Uuid:       uuid,
		activeTime: time.Now().Unix(),
	}

	h.Lock()
	h.clients[c] = true
	h.Unlock()

	g.Log().Debugf(r.Context(), "App WebSocket 连接建立: uuid=%s", uuid)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		msgStr := string(msg)
		// 自动应答 Ping 保活心跳协议
		if msgStr == "ping" || msgStr == "PING" {
			c.activeTime = time.Now().Unix()
			_ = ws.WriteMessage(1, []byte("pong"))
		}
	}

	h.Lock()
	delete(h.clients, c)
	h.Unlock()
	_ = ws.Close()
	g.Log().Debugf(r.Context(), "App WebSocket 断开连接: uuid=%s", uuid)
}

// BroadcastApp 向所有在连的 App 前端广播消息 (例如：K线推送，最新成交价推送等)
func BroadcastApp(ctx context.Context, data interface{}) {
	var msgBytes []byte
	switch v := data.(type) {
	case string:
		msgBytes = []byte(v)
	case []byte:
		msgBytes = v
	default:
		msgBytes, _ = gjson.Encode(v)
	}

	AppHub.RLock()
	defer AppHub.RUnlock()

	for c := range AppHub.clients {
		err := c.conn.WriteMessage(1, msgBytes)
		if err != nil {
			g.Log().Warningf(ctx, "App WebSocket 消息发送失败给 uuid=%s: %v", c.Uuid, err)
		}
	}
}
