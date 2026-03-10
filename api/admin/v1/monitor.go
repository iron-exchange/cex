package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 服务监控 (Server Monitor) ---------
type GetAdminServerInfoReq struct {
	g.Meta `path:"/monitor/server" tags:"AdminMonitor" method:"get" summary:"获取服务监控信息"`
}

type GetAdminServerInfoRes struct {
	Cpu      g.Map   `json:"cpu"`
	Mem      g.Map   `json:"mem"`
	Sys      g.Map   `json:"sys"`
	Jvm      g.Map   `json:"jvm"` // Map correctly to UI expecting 'jvm' for compatibility, populating golang info
	SysFiles []g.Map `json:"sysFiles"`
}

// --------- 缓存监控 (Cache Monitor) ---------
type GetAdminCacheInfoReq struct {
	g.Meta `path:"/monitor/cache" tags:"AdminMonitor" method:"get" summary:"获取缓存监控"`
}

type GetAdminCacheInfoRes struct {
	Info         g.Map   `json:"info"`
	DbSize       int64   `json:"dbSize"`
	CommandStats []g.Map `json:"commandStats"`
}

// --------- 缓存列表 (Cache List) ---------
type GetAdminCacheNamesReq struct {
	g.Meta `path:"/monitor/cache/getNames" tags:"AdminMonitor" method:"get" summary:"获取缓存名称列表"`
}

type SysCache struct {
	CacheName string `json:"cacheName"`
	Remark    string `json:"remark"`
}

type GetAdminCacheNamesRes struct {
	Data []SysCache `json:"data"` // Default ruoyi uses "data" object unwrap or array root
}

type GetAdminCacheKeysReq struct {
	g.Meta    `path:"/monitor/cache/getKeys/{cacheName}" tags:"AdminMonitor" method:"get" summary:"获取缓存Key列表"`
	CacheName string `json:"cacheName" in:"path"`
}

type GetAdminCacheKeysRes struct {
	Data []string `json:"data"`
}

type GetAdminCacheValueReq struct {
	g.Meta    `path:"/monitor/cache/getValue/{cacheName}/{cacheKey}" tags:"AdminMonitor" method:"get" summary:"获取缓存Key内容"`
	CacheName string `json:"cacheName" in:"path"`
	CacheKey  string `json:"cacheKey" in:"path"`
}

type SysCacheValue struct {
	CacheName  string `json:"cacheName"`
	CacheKey   string `json:"cacheKey"`
	CacheValue string `json:"cacheValue"`
}

type GetAdminCacheValueRes struct {
	Data SysCacheValue `json:"data"`
}

type ClearAdminCacheNameReq struct {
	g.Meta    `path:"/monitor/cache/clearCacheName/{cacheName}" tags:"AdminMonitor" method:"delete" summary:"清除指定缓存名"`
	CacheName string `json:"cacheName" in:"path"`
}
type ClearAdminCacheNameRes struct{}

type ClearAdminCacheKeyReq struct {
	g.Meta   `path:"/monitor/cache/clearCacheKey/{cacheKey}" tags:"AdminMonitor" method:"delete" summary:"清除缓存Key"`
	CacheKey string `json:"cacheKey" in:"path"`
}
type ClearAdminCacheKeyRes struct{}

type ClearAdminCacheAllReq struct {
	g.Meta `path:"/monitor/cache/clearCacheAll" tags:"AdminMonitor" method:"delete" summary:"清理全部缓存"`
}
type ClearAdminCacheAllRes struct{}
