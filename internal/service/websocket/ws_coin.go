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

// coinClient 承载具体外场币种行情的客户端
type coinClient struct {
	conn       *ghttp.WebSocket
	UserId     string
	activeTime int64
	// symbolSubs 存储客户端主动订阅的币种
	symbolSubs map[string]bool
	sync.RWMutex
}

// wsCoinHub 管理行情推送专用 WebSocket
type wsCoinHub struct {
	sync.RWMutex
	clients map[*coinClient]bool
}

var (
	CoinHub = &wsCoinHub{
		clients: make(map[*coinClient]bool),
	}
)

func init() {
	go CoinHub.pingRoutine()
	go CoinHub.subscribeTickerPgsql()
}

// CoinConnect 处理 /ws/coin/{userId} 端点连接
func (h *wsCoinHub) CoinConnect(r *ghttp.Request) {
	userId := r.Get("userId").String()
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), "Coin WebSocket 升级失败:", err)
		return
	}

	c := &coinClient{
		conn:       ws,
		UserId:     userId,
		activeTime: time.Now().Unix(),
		symbolSubs: make(map[string]bool),
	}

	h.Lock()
	h.clients[c] = true
	h.Unlock()

	g.Log().Debugf(r.Context(), "Coin WebSocket 连接建立: userId=%s", userId)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		msgStr := strings.TrimSpace(string(msg))
		c.activeTime = time.Now().Unix()

		g.Log().Infof(r.Context(), "CoinHub 收到前端原始消息: [%s]", msgStr)

		// 简单指令协议
		lowerMsg := strings.ToLower(msgStr)
		if strings.Contains(lowerMsg, "unsub:") {
			// 先处理 unsub，防止混淆
			idx := strings.Index(lowerMsg, "unsub:")
			symbol := msgStr[idx+6:]
			symbol = strings.TrimSpace(strings.ReplaceAll(symbol, "\"", ""))
			symbol = strings.ToUpper(strings.TrimSpace(symbol))
			c.Lock()
			delete(c.symbolSubs, symbol)
			c.Unlock()
			_ = ws.WriteMessage(1, []byte("unsub-success: "+symbol))
			g.Log().Infof(r.Context(), "CoinHub 客户端 %s 取消关注: [%s]", c.UserId, symbol)
		} else if strings.Contains(lowerMsg, "sub:") {
			idx := strings.Index(lowerMsg, "sub:")
			symbol := msgStr[idx+4:]
			symbol = strings.TrimSpace(strings.ReplaceAll(symbol, "\"", ""))
			symbol = strings.ToUpper(strings.TrimSpace(symbol)) // 强转大写，防止前端传 btc_usdt
			c.Lock()
			c.symbolSubs[symbol] = true
			c.Unlock()
			_ = ws.WriteMessage(1, []byte("sub-success: "+symbol))
			g.Log().Infof(r.Context(), "CoinHub 客户端 %s 成功关注: [%s]", c.UserId, symbol)
		} else if lowerMsg == "ping" {
			_ = ws.WriteMessage(1, []byte("pong"))
		}
	}

	h.Lock()
	delete(h.clients, c)
	h.Unlock()
	_ = ws.Close()
	g.Log().Debugf(r.Context(), "Coin WebSocket 断开连接: userId=%s", c.UserId)
}

func (h *wsCoinHub) subscribeTickerPgsql() {
	ctx := context.Background()

	connStr := "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=cex sslmode=disable"
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			g.Log().Error(ctx, "PostgreSQL Ticker Listener 异常:", err)
		}
	}

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("market_ticker")
	if err != nil {
		g.Log().Error(ctx, "CoinHub 订阅行情 market_ticker 失败", err)
		return
	}

	g.Log().Info(ctx, "已启动 PostgreSQL 行情通道: market_ticker")

	for {
		select {
		case notice := <-listener.Notify:
			if notice != nil {
				// notice.Extra 是形如: {"BTC_USDT": {"price": "65000", ...}}
				h.broadcastMarketTicker(ctx, []byte(notice.Extra))
			}
		case <-time.After(90 * time.Second):
			go func() {
				_ = listener.Ping()
			}()
		}
	}
}

func (h *wsCoinHub) broadcastMarketTicker(ctx context.Context, payload []byte) {
	// 尝试解析出行情中的 symbol，进行精准分发
	j, err := gjson.DecodeToJson(payload)
	if err != nil {
		return
	}

	// j.Map() 包含了诸如 map["BTC_USDT"]interface{}...
	symbols := make([]string, 0)
	for k := range j.Map() {
		symbols = append(symbols, k)
	}
	if len(symbols) == 0 {
		return
	}

	h.RLock()
	defer h.RUnlock()

	for c := range h.clients {
		// 检查这个客户端是否订阅了里面的任意一个 symbol
		shouldSend := false
		c.RLock()
		if len(c.symbolSubs) == 0 {
			// 如果没发任何 sub 指令，则什么都不推，防止流量爆炸卡死前端
			shouldSend = false
		} else {
			for _, sym := range symbols {
				if c.symbolSubs[sym] {
					shouldSend = true
					break
				} else {
					// 仅调试打印，对比实际内存里的 subs 和行情 sym
					g.Log().Debugf(ctx, "userId=%s mismatch: sym=[%s], subs=%v", c.UserId, sym, c.symbolSubs)
				}
			}
		}
		c.RUnlock()

		if shouldSend {
			err := c.conn.WriteMessage(1, payload)
			if err != nil {
				g.Log().Warningf(ctx, "CoinHub 推送失败给 userId=%s: %v", c.UserId, err)
			}
		}
	}
}

func (h *wsCoinHub) pingRoutine() {
	for {
		time.Sleep(30 * time.Second)
		h.Lock()
		now := time.Now().Unix()
		for c := range h.clients {
			if now-c.activeTime > 60 {
				_ = c.conn.Close()
				delete(h.clients, c)
			}
		}
		h.Unlock()
	}
}
