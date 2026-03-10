package monitor

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/monitor"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// --------- 服务监控 (Server Monitor) ---------

func (c *Controller) GetServerInfo(ctx context.Context, req *v1.GetAdminServerInfoReq) (res *v1.GetAdminServerInfoRes, err error) {
	return monitor.New().GetServerInfo(ctx, req)
}

// --------- 缓存监控 (Cache Monitor) ---------

func (c *Controller) GetCacheInfo(ctx context.Context, req *v1.GetAdminCacheInfoReq) (res *v1.GetAdminCacheInfoRes, err error) {
	return monitor.New().GetCacheInfo(ctx, req)
}

func (c *Controller) GetCacheNames(ctx context.Context, req *v1.GetAdminCacheNamesReq) (res *v1.GetAdminCacheNamesRes, err error) {
	return monitor.New().GetCacheNames(ctx, req)
}

func (c *Controller) GetCacheKeys(ctx context.Context, req *v1.GetAdminCacheKeysReq) (res *v1.GetAdminCacheKeysRes, err error) {
	return monitor.New().GetCacheKeys(ctx, req)
}

func (c *Controller) GetCacheValue(ctx context.Context, req *v1.GetAdminCacheValueReq) (res *v1.GetAdminCacheValueRes, err error) {
	return monitor.New().GetCacheValue(ctx, req)
}

func (c *Controller) ClearCacheName(ctx context.Context, req *v1.ClearAdminCacheNameReq) (res *v1.ClearAdminCacheNameRes, err error) {
	return monitor.New().ClearCacheName(ctx, req)
}

func (c *Controller) ClearCacheKey(ctx context.Context, req *v1.ClearAdminCacheKeyReq) (res *v1.ClearAdminCacheKeyRes, err error) {
	return monitor.New().ClearCacheKey(ctx, req)
}

func (c *Controller) ClearCacheAll(ctx context.Context, req *v1.ClearAdminCacheAllReq) (res *v1.ClearAdminCacheAllRes, err error) {
	return monitor.New().ClearCacheAll(ctx, req)
}
