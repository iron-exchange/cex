package common

import (
	"context"
	"time"

	v1 "GoCEX/app/api"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

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

// GetAllSetting 获取全部全站配置参数 (字典表下放)
func (s *sCommon) GetAllSetting(ctx context.Context) (*v1.GetAllSettingRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:all_settings", func(ctx context.Context) (interface{}, error) {
		var configs []entity.SysConfig
		err := dao.SysConfig.Ctx(ctx).Scan(&configs)
		if err != nil {
			return nil, err
		}

		res := &v1.GetAllSettingRes{
			Settings: make(map[string]string),
			List:     make([]v1.SettingInfo, 0, len(configs)),
		}

		for _, v := range configs {
			res.Settings[v.ConfigKey] = v.ConfigValue
			res.List = append(res.List, v1.SettingInfo{
				Key:   v.ConfigKey,
				Value: v.ConfigValue,
				Desc:  v.ConfigName,
			})
		}
		return res, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetAllSettingRes), nil
}

// GetAppSidebarSetting 获取侧边栏显示的币种
func (s *sCommon) GetAppSidebarSetting(ctx context.Context) (*v1.GetAppSidebarSettingRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:sidebar_coins", func(ctx context.Context) (interface{}, error) {
		// 为了简化，假设所有的币种主档配置都在 currency_symbol 或者 sys_dict_data，
		// 根据原版系统特性，通常放在类似 t_app_currency_symbol 进行显示控制
		var symbols []entity.CurrencySymbol
		err := dao.CurrencySymbol.Ctx(ctx).Where(dao.CurrencySymbol.Columns().IsShow, "1").Scan(&symbols)
		if err != nil {
			return nil, err
		}

		list := make([]string, 0, len(symbols))
		for _, sym := range symbols {
			list = append(list, sym.Symbol)
		}
		return &v1.GetAppSidebarSettingRes{List: list}, nil
	}, 1*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetAppSidebarSettingRes), nil
}

// GetHomeCoinSetting 获取首页主推币种
func (s *sCommon) GetHomeCoinSetting(ctx context.Context) (*v1.GetHomeCoinSettingRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:home_coins", func(ctx context.Context) (interface{}, error) {
		var symbols []entity.CurrencySymbol
		// 结合原架构规则，筛选特定标识位的热门或推荐币种
		err := dao.CurrencySymbol.Ctx(ctx).Where(dao.CurrencySymbol.Columns().IsShow, "1").Limit(10).Scan(&symbols)
		if err != nil {
			return nil, err
		}

		list := make([]string, 0, len(symbols))
		for _, sym := range symbols {
			list = append(list, sym.Symbol)
		}
		return &v1.GetHomeCoinSettingRes{List: list}, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetHomeCoinSettingRes), nil
}

// GetAppCurrencyList 获取充值的通道与开关列表
func (s *sCommon) GetAppCurrencyList(ctx context.Context) (*v1.GetAppCurrencyListRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:recharge_channels", func(ctx context.Context) (interface{}, error) {
		var manage []entity.SymbolManage
		// 充值通道通常挂载在 SymbolManage 或者专属的通道表。这里以 Manage 表演示配置加载
		err := dao.SymbolManage.Ctx(ctx).Where(dao.SymbolManage.Columns().Enable, "1").Scan(&manage)
		if err != nil {
			return nil, err
		}

		list := make([]v1.CurrencyChannelInfo, 0, len(manage))
		for _, m := range manage {
			// 将 DB 模型字段脱敏组装
			list = append(list, v1.CurrencyChannelInfo{
				CoinName:       m.Symbol,
				Type:           "TRC20", // 伪数据映射 (基于具体业务表拓展)
				MinLimit:       "10",
				MaxLimit:       "1000000",
				DepositAddress: "T...Address.Placeholder",
				IsOpen:         1,
			})
		}
		return &v1.GetAppCurrencyListRes{List: list}, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetAppCurrencyListRes), nil
}

// GetWithDrawCoinList 获取提现的通道与手续费列表
func (s *sCommon) GetWithDrawCoinList(ctx context.Context) (*v1.GetWithDrawCoinListRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:withdraw_channels", func(ctx context.Context) (interface{}, error) {
		var symbols []entity.CurrencySymbol
		// 通用提现依赖于 CurrencySymbol 表里的设置
		err := dao.CurrencySymbol.Ctx(ctx).Where(dao.CurrencySymbol.Columns().Enable, "1").Scan(&symbols)
		if err != nil {
			return nil, err
		}

		list := make([]v1.WithdrawCoinInfo, 0, len(symbols))
		for _, sym := range symbols {
			list = append(list, v1.WithdrawCoinInfo{
				CoinName: sym.Coin,
				Type:     "ERC20/TRC20", // 通道说明
				MinLimit: "0",
				MaxLimit: "999999",
				FeeRate:  "0.01",
				IsOpen:   1,
			})
		}
		return &v1.GetWithDrawCoinListRes{List: list}, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetWithDrawCoinListRes), nil
}
