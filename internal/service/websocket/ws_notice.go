package websocket

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lib/pq"
)

// client 代表一个 WebSocket 连接客户端
type client struct {
	conn       *ghttp.WebSocket
	UserId     string
	Uuid       string
	activeTime int64
}

// wsHub 管理所有 WebSocket 连接和消息广播
type wsHub struct {
	sync.RWMutex
	clients map[*client]bool
}

var (
	NoticeHub = &wsHub{
		clients: make(map[*client]bool),
	}
)

func init() {
	// 启动心跳检测，清理死连接
	go NoticeHub.pingRoutine()

	// 启动 PostgreSQL 跨进程订阅
	go NoticeHub.subscribePgsql()
}

// subscribePgsql 从 Postgres 订阅跨进程通知，转发给当前进程连接的所有 WebSocket 客户端
func (h *wsHub) subscribePgsql() {
	ctx := context.Background()

	connStr := "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=cex sslmode=disable"
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			g.Log().Error(ctx, "PostgreSQL Listener 出现异常:", err)
		}
	}

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("cex_admin_notice")
	if err != nil {
		g.Log().Error(ctx, "WebSocket 跨进程 PostgreSQL 订阅失败", err)
		return
	}

	g.Log().Info(ctx, "已启动 PostgreSQL LISTEN 通道: cex_admin_notice")

	for {
		select {
		case notice := <-listener.Notify:
			if notice != nil {
				h.BroadcastAll(ctx, []byte(notice.Extra))
			}
		case <-time.After(90 * time.Second):
			go func() {
				_ = listener.Ping()
			}()
		}
	}
}

// Connect 处理 WebSocket 升级并注册客户端
func (h *wsHub) Connect(r *ghttp.Request) {
	userId := r.Get("userId").String()
	uuid := r.Get("uuId").String()

	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), "WebSocket升级失败:", err)
		return
	}

	c := &client{
		conn:       ws,
		UserId:     userId,
		Uuid:       uuid,
		activeTime: time.Now().Unix(),
	}

	h.Lock()
	h.clients[c] = true
	h.Unlock()

	g.Log().Debugf(r.Context(), "WebSocket 客户端连接: userId=%s, uuid=%s", userId, uuid)

	// 阻塞读取，直到连接断开
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 可以处理前端发来的心跳包（比如前端发 "ping"）
		if string(msg) == "ping" {
			c.activeTime = time.Now().Unix()
			_ = ws.WriteMessage(1, []byte("pong"))
		}
	}

	// 退出循环说明连接断开，清理客户端
	h.Lock()
	delete(h.clients, c)
	h.Unlock()
	_ = ws.Close()
	g.Log().Debugf(r.Context(), "WebSocket 客户端断开: userId=%s, uuid=%s", userId, uuid)
}

// BroadcastAll 广播消息给所有连接的客户端 (兼容老版 WebSocketNotice.sendInfoAll)
func (h *wsHub) BroadcastAll(ctx context.Context, message interface{}) {
	var msgBytes []byte
	switch v := message.(type) {
	case string:
		msgBytes = []byte(v)
	case []byte:
		msgBytes = v
	default:
		msgBytes, _ = gjson.Encode(v)
	}

	h.RLock()
	defer h.RUnlock()

	for c := range h.clients {
		err := c.conn.WriteMessage(1, msgBytes) // 1 = TextMessage
		if err != nil {
			g.Log().Warningf(ctx, "WebSocket 消息发送失败给 userId=%s: %v", c.UserId, err)
		}
	}
}

// SendToUser 发送消息给特定 UserId 的所有会话 (如需定向发送)
func (h *wsHub) SendToUser(ctx context.Context, userId string, message interface{}) {
	var msgBytes []byte
	switch v := message.(type) {
	case string:
		msgBytes = []byte(v)
	case []byte:
		msgBytes = v
	default:
		msgBytes, _ = gjson.Encode(v)
	}

	h.RLock()
	defer h.RUnlock()

	for c := range h.clients {
		if c.UserId == userId {
			_ = c.conn.WriteMessage(1, msgBytes)
		}
	}
}

// PingRoutine 定期清理超时或者异常断开的死连接
func (h *wsHub) pingRoutine() {
	for {
		time.Sleep(30 * time.Second)
		h.Lock()
		now := time.Now().Unix()
		for c := range h.clients {
			// 一直ping不通前端的话，前端连接大概率死亡 (这里设60秒心跳容忍)
			if now-c.activeTime > 60 {
				_ = c.conn.Close()
				delete(h.clients, c)
			}
		}
		h.Unlock()
	}
}

// SendInfoAll 供当前进程内的业务调用，直接广播给所有的 Admin 客户端
func SendInfoAll(ctx context.Context, eventType int) {
	msgStr := gconv.String(eventType)
	NoticeHub.BroadcastAll(ctx, []byte(msgStr))
}

// PublishToAdmin 由 App/Task 等子系统调用，将通知通过 Postgres 分发给真正的 Admin WebSocket 集群
func PublishToAdmin(ctx context.Context, eventType int) {
	msgStr := gconv.String(eventType)

	_, err := g.DB().Exec(ctx, "SELECT pg_notify($1, $2)", "cex_admin_notice", msgStr)
	if err != nil {
		g.Log().Error(ctx, "推送跨端 WebSocket 消息至 PostgreSQL 失败", err)
	}
}
