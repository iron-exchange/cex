package websocket

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lib/pq"
)

// wsCoinOverHub 管理秒合约订单结算等相关推送
type wsCoinOverHub struct {
	sync.RWMutex
	clients map[*client]bool
}

var (
	CoinOverHub = &wsCoinOverHub{
		clients: make(map[*client]bool),
	}
)

func init() {
	go CoinOverHub.pingRoutine()
	go CoinOverHub.subscribeCoinOverPgsql()
}

// CoinOverConnect 处理 /webSocket/coinOver/{userId} 接入
func (h *wsCoinOverHub) CoinOverConnect(r *ghttp.Request) {
	userId := r.Get("userId").String()
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), "CoinOver WebSocket升级失败:", err)
		return
	}

	c := &client{
		conn:       ws,
		UserId:     userId,
		activeTime: time.Now().Unix(),
	}

	h.Lock()
	h.clients[c] = true
	h.Unlock()

	g.Log().Debugf(r.Context(), "CoinOver WebSocket 连接建立: userId=%s", userId)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		msgStr := string(msg)
		if strings.ToLower(msgStr) == "ping" {
			c.activeTime = time.Now().Unix()
			_ = ws.WriteMessage(1, []byte("pong"))
		}
	}

	h.Lock()
	delete(h.clients, c)
	h.Unlock()
	_ = ws.Close()
	g.Log().Debugf(r.Context(), "CoinOver WebSocket 断开连接: userId=%s", userId)
}

// subscribeCoinOverPgsql 从 Postgres 订阅跨进程通知
func (h *wsCoinOverHub) subscribeCoinOverPgsql() {
	ctx := context.Background()

	connStr := "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=cex sslmode=disable"
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			g.Log().Error(ctx, "PostgreSQL CoinOver Listener 出现异常:", err)
		}
	}

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("cex_coin_over")
	if err != nil {
		g.Log().Error(ctx, "WebSocket 跨进程 PostgreSQL 订阅失败 (coinOver)", err)
		return
	}

	g.Log().Info(ctx, "已启动 PostgreSQL LISTEN 通道: cex_coin_over")

	for {
		select {
		case notice := <-listener.Notify:
			if notice != nil {
				// 获取传过来的 JSON 并解析 userId 进行点对点发送
				j, err := gjson.DecodeToJson([]byte(notice.Extra))
				if err == nil {
					targetUid := j.Get("userId").String()
					h.SendToUser(ctx, targetUid, []byte(notice.Extra))
				}
			}
		case <-time.After(90 * time.Second):
			go func() {
				_ = listener.Ping()
			}()
		}
	}
}

// SendToUser 点对点推送到具体的 userId 网页
func (h *wsCoinOverHub) SendToUser(ctx context.Context, userId string, message []byte) {
	h.RLock()
	defer h.RUnlock()

	for c := range h.clients {
		if c.UserId == userId {
			err := c.conn.WriteMessage(1, message)
			if err != nil {
				g.Log().Warningf(ctx, "CoinOver 消息发送失败给 userId=%s: %v", c.UserId, err)
			}
		}
	}
}

func (h *wsCoinOverHub) pingRoutine() {
	for {
		time.Sleep(30 * time.Second)
		h.Lock()
		now := time.Now().Unix()
		for c := range h.clients {
			if now-c.activeTime > 180 {
				// 超过 3 分钟没有任何活跃迹象才真正踢掉
				_ = c.conn.Close()
				delete(h.clients, c)
			} else {
				// 主动向客户端发心跳 ping，维持连接续命
				c.activeTime = now
				_ = c.conn.WriteMessage(1, []byte("ping"))
			}
		}
		h.Unlock()
	}
}

// PublishCoinOver 由定时结算 Task 进程调用，推送该用户的开奖结果给 App 前端
func PublishCoinOver(ctx context.Context, userId int, orderNo string, reward float64, result string) {
	msg := g.Map{
		"userId":  userId,
		"orderNo": orderNo,
		"reward":  reward,
		"result":  result, // 例如 "WIN", "LOSE", "DRAW"
	}
	msgBytes, _ := gjson.Encode(msg)

	_, err := g.DB().Exec(ctx, "SELECT pg_notify($1, $2)", "cex_coin_over", string(msgBytes))
	if err != nil {
		g.Log().Error(ctx, "推送跨端 CoinOver 消息至 PostgreSQL 失败", err)
	}
}
