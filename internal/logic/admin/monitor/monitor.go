package monitor

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	v1 "GoCEX/api/admin/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var startTime = time.Now()

type sAdminMonitor struct{}

func New() *sAdminMonitor {
	return &sAdminMonitor{}
}

func (s *sAdminMonitor) GetServerInfo(ctx context.Context, req *v1.GetAdminServerInfoReq) (*v1.GetAdminServerInfoRes, error) {
	// 1. CPU
	cpuStats, _ := cpu.Info()
	cpuPercents, _ := cpu.Percent(0, false)
	cpuNum := 0
	if len(cpuStats) > 0 {
		cpuNum = len(cpuStats) // logical cores
	}
	cpuTotal := 100.0
	cpuSys := 0.0
	cpuUsed := 0.0
	cpuFree := 100.0
	if len(cpuPercents) > 0 {
		cpuSys = cpuPercents[0] / 2
		cpuUsed = cpuPercents[0] / 2
		cpuFree = 100 - cpuPercents[0]
	}

	// 2. Mem
	vMem, _ := mem.VirtualMemory()

	// 3. Sys (Host)
	hostStat, _ := host.Info()

	// 4. SysFiles (Disk)
	partitions, _ := disk.Partitions(false)
	var sysFiles []g.Map
	for _, p := range partitions {
		usage, _ := disk.Usage(p.Mountpoint)
		if usage != nil {
			sysFiles = append(sysFiles, g.Map{
				"dirName":     p.Mountpoint,
				"sysTypeName": p.Fstype,
				"typeName":    p.Device,
				"total":       fmt.Sprintf("%.2f GB", float64(usage.Total)/1024/1024/1024),
				"free":        fmt.Sprintf("%.2f GB", float64(usage.Free)/1024/1024/1024),
				"used":        fmt.Sprintf("%.2f GB", float64(usage.Used)/1024/1024/1024),
				"usage":       usage.UsedPercent,
			})
		}
	}

	// Calculate uptime
	uptimeRaw := time.Since(startTime)
	runTime := fmt.Sprintf("%d天%d小时%d分钟", int(uptimeRaw.Hours()/24), int(uptimeRaw.Hours())%24, int(uptimeRaw.Minutes())%60)

	return &v1.GetAdminServerInfoRes{
		Cpu: g.Map{
			"cpuNum": cpuNum,
			"total":  cpuTotal,
			"sys":    cpuSys,
			"used":   cpuUsed,
			"wait":   0.0,
			"free":   cpuFree,
		},
		Mem: g.Map{
			"total": float64(vMem.Total) / 1024 / 1024 / 1024,
			"used":  float64(vMem.Used) / 1024 / 1024 / 1024,
			"free":  float64(vMem.Free) / 1024 / 1024 / 1024,
			"usage": vMem.UsedPercent,
		},
		Sys: g.Map{
			"computerName": hostStat.Hostname,
			"computerIp":   "127.0.0.1",
			"osName":       hostStat.OS,
			"osArch":       runtime.GOARCH,
		},
		Jvm: g.Map{
			"name":      "Go Runtime",
			"version":   runtime.Version(),
			"home":      runtime.GOROOT(),
			"startTime": startTime.Format("2006-01-02 15:04:05"),
			"runTime":   runTime,
		},
		SysFiles: sysFiles,
	}, nil
}

func (s *sAdminMonitor) GetCacheInfo(ctx context.Context, req *v1.GetAdminCacheInfoReq) (*v1.GetAdminCacheInfoRes, error) {
	// INFO
	infoRes, err := g.Redis().Do(ctx, "INFO")
	if err != nil {
		return nil, err
	}
	infoStr := gconv.String(infoRes)
	infoMap := g.Map{}
	for _, line := range strings.Split(infoStr, "\r\n") {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			infoMap[parts[0]] = parts[1]
		}
	}

	// DBSIZE
	dbSizeRes, _ := g.Redis().Do(ctx, "DBSIZE")
	dbSize := gconv.Int64(dbSizeRes)

	// CommandStats
	commandStatsRes, _ := g.Redis().Do(ctx, "INFO", "commandstats")
	commandStatsStr := gconv.String(commandStatsRes)
	var commandStats []g.Map
	for _, line := range strings.Split(commandStatsStr, "\r\n") {
		if strings.HasPrefix(line, "cmdstat_") {
			parts := strings.SplitN(line, ":", 2)
			cmdName := strings.TrimPrefix(parts[0], "cmdstat_")
			// calls=3,usec=15,usec_per_call=5.00
			vals := strings.Split(parts[1], ",")
			callsStr := strings.TrimPrefix(vals[0], "calls=")
			commandStats = append(commandStats, g.Map{
				"name":  cmdName,
				"value": callsStr,
			})
		}
	}

	return &v1.GetAdminCacheInfoRes{
		Info:         infoMap,
		DbSize:       dbSize,
		CommandStats: commandStats,
	}, nil
}

func (s *sAdminMonitor) GetCacheNames(ctx context.Context, req *v1.GetAdminCacheNamesReq) (*v1.GetAdminCacheNamesRes, error) {
	// Mock Ruoyi common cache names directly
	list := []v1.SysCache{
		{CacheName: "CEX:VERIFY_CODE:", Remark: "验证码"},
		{CacheName: "CEX:DICT:", Remark: "系统字典"},
		{CacheName: "CEX:CONFIG:", Remark: "系统参数"},
		{CacheName: "SYS_LOGIN_TOKEN:", Remark: "登录Token"},
		{CacheName: "rate_limit:", Remark: "防刷限流"},
		{CacheName: "CURRENCY_PRICE:", Remark: "行情缓存"},
		{CacheName: "trade_lock:", Remark: "交易分布式锁"},
	}
	return &v1.GetAdminCacheNamesRes{Data: list}, nil
}

func (s *sAdminMonitor) GetCacheKeys(ctx context.Context, req *v1.GetAdminCacheKeysReq) (*v1.GetAdminCacheKeysRes, error) {
	res, _ := g.Redis().Do(ctx, "KEYS", req.CacheName+"*")
	return &v1.GetAdminCacheKeysRes{
		Data: gconv.Strings(res),
	}, nil
}

func (s *sAdminMonitor) GetCacheValue(ctx context.Context, req *v1.GetAdminCacheValueReq) (*v1.GetAdminCacheValueRes, error) {
	val, _ := g.Redis().Do(ctx, "GET", req.CacheKey)
	return &v1.GetAdminCacheValueRes{
		Data: v1.SysCacheValue{
			CacheName:  req.CacheName,
			CacheKey:   req.CacheKey,
			CacheValue: gconv.String(val),
		},
	}, nil
}

func (s *sAdminMonitor) ClearCacheName(ctx context.Context, req *v1.ClearAdminCacheNameReq) (*v1.ClearAdminCacheNameRes, error) {
	keys, _ := g.Redis().Do(ctx, "KEYS", req.CacheName+"*")
	for _, k := range gconv.Strings(keys) {
		g.Redis().Do(ctx, "DEL", k)
	}
	return &v1.ClearAdminCacheNameRes{}, nil
}

func (s *sAdminMonitor) ClearCacheKey(ctx context.Context, req *v1.ClearAdminCacheKeyReq) (*v1.ClearAdminCacheKeyRes, error) {
	g.Redis().Do(ctx, "DEL", req.CacheKey)
	return &v1.ClearAdminCacheKeyRes{}, nil
}

func (s *sAdminMonitor) ClearCacheAll(ctx context.Context, req *v1.ClearAdminCacheAllReq) (*v1.ClearAdminCacheAllRes, error) {
	g.Redis().Do(ctx, "FLUSHDB")
	return &v1.ClearAdminCacheAllRes{}, nil
}
