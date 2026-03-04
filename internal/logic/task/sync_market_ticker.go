package task

import (
	"context"
)

// SyncMarketTicker 外部大盘K线与Tick行情定时拉取调度
// [废弃声明] - 此定时器已被废弃，大盘行情拉取统一交由更实时的 WebSocket 守护进程处理
// 请参见: `internal/service/market/binance_ws.go` (StartBinanceWSDaemon)
func (s *sTask) SyncMarketTicker(ctx context.Context) error {
	// g.Log().Info(ctx, "[废弃] 定时大盘行情轮询已下线，统一由 WebSocket Daemon 实时拉取")
	return nil
}
