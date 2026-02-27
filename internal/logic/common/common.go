package common

import (
	"context"
	"time"

	v1 "GoCEX/api/app/v1"

	"github.com/gogf/gf/v2/os/gcache"
)

type sCommon struct{}

func New() *sCommon {
	return &sCommon{}
}

// GetConfig 获取系统公共配置 (利用 gcache 减轻 DB 扫库压力)
func (s *sCommon) GetConfig(ctx context.Context) (*v1.CommonConfigRes, error) {
	// 5 分钟热更新缓存，防止大推流下击穿 DB
	val, err := gcache.GetOrSetFunc(ctx, "app:common:config", func(ctx context.Context) (interface{}, error) {
		// 在真实业务中，这里去查 dao.SysConfig
		cfg := &v1.CommonConfigRes{
			CustomerServiceUrl: "https://t.me/exchange_support",
			RechargeMinAmount:  "10.00",
			Banners: []string{
				"https://img.cdn.com/banner_1.png",
				"https://img.cdn.com/banner_2.png",
			},
		}
		return cfg, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}

	return val.Val().(*v1.CommonConfigRes), nil
}
