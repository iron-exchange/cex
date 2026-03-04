package market

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
)

type binanceTickerResponse struct {
	Symbol      string      `json:"s"` // 交易对 (例如 BTCUSDT)
	ClosePrice  interface{} `json:"c"` // 最新成交价
	PriceChange interface{} `json:"P"` // 24小时涨跌幅 (百分比)
	Volume      interface{} `json:"v"` // 24小时成交量
}

// 订阅的币种，如果是生产盘口，建议提炼到配置里
// 这里转化成了 binance stream 格式: btcusdt@ticker
var subscribeSymbols = []string{
	"btcusdt@ticker",
	"ethusdt@ticker",
	"dogeusdt@ticker",
	"trxusdt@ticker",
	"solusdt@ticker",
}

// StartBinanceWSDaemon 启动币安行情守护进程
func StartBinanceWSDaemon(ctx context.Context) {
	// 连接串组合，形如 wss://stream.binance.com:9443/stream?streams=btcusdt@ticker/ethusdt@ticker
	streams := strings.Join(subscribeSymbols, "/")
	wsURL := fmt.Sprintf("wss://stream.binance.com:9443/stream?streams=%s", streams)

	g.Log().Info(ctx, "🚀 启动 Binance WebSocket 行情守护进程...", wsURL)

	go func() {
		for { // 外层死循环：断线无限重连
			connectAndServe(ctx, wsURL)
			g.Log().Warning(ctx, "⚠️ Binance WS 连接断开，3秒后重连...")
			time.Sleep(3 * time.Second)
		}
	}()
}

func connectAndServe(ctx context.Context, wsURL string) {
	// 拨号连接
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		g.Log().Error(ctx, "❌ 连接 Binance WebSocket 失败:", err)
		return
	}
	defer conn.Close()

	g.Log().Info(ctx, "✅ 成功连入 Binance WebSocket")

	// 心跳维持，应对被动断开
	// 币安流如果在15分钟内没有ping pong会被切断，
	pingTicker := time.NewTicker(3 * time.Minute)
	defer pingTicker.Stop()

	// 专门处理 Ping 发送的 Goroutine
	go func() {
		for range pingTicker.C {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

	// 数据收取死循环
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			g.Log().Error(ctx, "读取 WS 消息报错或被服务端掐断:", err)
			break
		}

		// 币安组合流的数据最外层会有个 stream 字段和 data 字段
		var payload struct {
			Stream string                `json:"stream"`
			Data   binanceTickerResponse `json:"data"`
		}
		if err := json.Unmarshal(message, &payload); err != nil {
			g.Log().Error(ctx, "JSON Unmarshal error:", err, "Raw WS Message:", string(message))
			continue
		}

		// 类型抹平：防止上游既发数字又发字符串
		closePrice := fmt.Sprintf("%v", payload.Data.ClosePrice)
		priceChange := fmt.Sprintf("%v", payload.Data.PriceChange)
		volume := fmt.Sprintf("%v", payload.Data.Volume)

		// 解析落库 Redis，抹平标准化的 `CURRENCY_PRICE:BTC` 键名
		symbolUp := strings.ToUpper(strings.ReplaceAll(payload.Data.Symbol, "USDT", ""))
		redisKey := "CURRENCY_PRICE:" + symbolUp

		// 存储最新价，定时器的 `SettleCurrencyOrder` 等强依赖这个缓存 (改为储存到 PG sys_config 以替换 Redis)
		rs, _ := g.DB().Model("sys_config").Data(g.Map{"config_value": closePrice, "update_time": gtime.Now()}).Where("config_key", redisKey).Update()
		if aff, _ := rs.RowsAffected(); aff == 0 {
			_, _ = g.DB().Model("sys_config").Data(g.Map{
				"config_name":  "行情最新价",
				"config_key":   redisKey,
				"config_value": closePrice,
				"config_type":  "Y",
				"create_time":  gtime.Now(),
			}).Insert()
		}

		// 附带一份全量简易数据供前端通过 /api/v1/market/ticker 取用
		tickerKey := "MARKET:TICKER_FULL:" + symbolUp
		_, _ = g.Redis().Do(ctx, "HSET", tickerKey, "price", closePrice, "change", priceChange, "vol", volume)

		// 将组装好的 JSON 发布到 Redis 的 MARKET:TICKER 广播频道
		// 统一组装成前端期待的格式: {"BTC_USDT": {"price": "65000.00", ...}}
		formattedData := g.Map{
			symbolUp + "_USDT": g.Map{
				"price":  closePrice,
				"change": priceChange,
				"vol":    volume,
			},
		}
		jsonMsg, _ := json.Marshal(formattedData)

		// 更改为 Postgres LISTEN/NOTIFY 架构 (替代 Redis Publish)
		_, _ = g.DB().Exec(ctx, "SELECT pg_notify($1, $2)", "market_ticker", string(jsonMsg))

		// 临时增加一条调试日志，看看每秒钟内部到底有没有在疯狂对外 Publish
		g.Log().Info(ctx, "[Binance WS] pg_notify market_ticker:", string(jsonMsg))
	}
}
