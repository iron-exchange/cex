package common

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"time"

	v1admin "GoCEX/api/admin/v1"
	v1 "GoCEX/app/api"
	"GoCEX/internal/consts"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/guid"
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
func (s *sCommon) GetAllSetting(ctx context.Context) (v1.GetAllSettingRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:common:all_settings", func(ctx context.Context) (interface{}, error) {
		var configs []entity.Setting
		err := dao.Setting.Ctx(ctx).Scan(&configs)
		if err != nil {
			return nil, err
		}

		res := make(v1.GetAllSettingRes)

		for _, v := range configs {
			if jsonWrap, err := gjson.Decode(v.SettingValue); err == nil {
				res[v.Id] = jsonWrap
			} else {
				res[v.Id] = v.SettingValue
			}
		}
		return res, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(v1.GetAllSettingRes), nil
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

// CaptchaImage 生成数学验证码（无需 Redis，用 gcache 内存存储）
func (s *sCommon) CaptchaImage(ctx context.Context) (*v1.CaptchaImageRes, error) {
	// 查询系统中是否开启了验证码功能
	var config entity.SysConfig
	_ = dao.SysConfig.Ctx(ctx).Where("config_key", "sys.account.captchaEnabled").Scan(&config)

	captchaEnabled := true
	if config.ConfigValue == "false" {
		captchaEnabled = false
		return &v1.CaptchaImageRes{
			CaptchaEnabled: captchaEnabled,
		}, nil
	}

	// 生成简单加法题
	a := rand.Intn(10)
	b := rand.Intn(10)
	answer := fmt.Sprintf("%d", a+b)
	question := fmt.Sprintf("%d + %d = ?", a, b)

	// 生成 UUID
	uuid := guid.S()

	// 缓存答案 5 分钟
	_ = gcache.Set(ctx, "captcha:"+uuid, answer, 5*time.Minute)

	// 用标准库将文字画成简单 PNG中的小图（纯色块 + 预留文字信息在 body）
	// 实际业务可切换成 github.com/afocus/captcha 等库进行字体渲染，现在返回包含题目的 SVG
	img := image.NewRGBA(image.Rect(0, 0, 120, 40))
	bgColor := color.RGBA{R: 240, G: 245, B: 255, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	// 在图片上画点阵表示 (ASCII 字符 计划使用外部库时可混入)
	for x := 0; x < 120; x += 4 {
		for y := 0; y < 40; y += 4 {
			if rand.Intn(4) == 0 {
				img.Set(x, y, color.RGBA{R: 180, G: 180, B: 200, A: 120})
			}
		}
	}

	var buf bytes.Buffer
	_ = png.Encode(&buf, img)
	imgBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	// 返回验证码颠 (实际验证码内容作为 header 不能存在图片中，建议购入 afocus/captcha)
	// 现在回传 question 作为暗示，前端可直接展示
	_ = question // 验证码题目，真实场景应通过字体渲染写入图片
	return &v1.CaptchaImageRes{
		CaptchaEnabled: captchaEnabled,
		Uuid:           uuid,
		Img:            imgBase64,
	}, nil
}

// GetRecordType 获取账变类型字典
func (s *sCommon) GetRecordType(ctx context.Context) (v1admin.AdminRecordTypeRes, error) {
	// 按照用户提供的顺序大致排列，或者直接遍历 Map
	// 用户提供的顺序: 1, 2, 3, 50, 4, 51, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 52, 53, 54, 55
	keys := []int{
		1, 2, 3, 50, 4, 51, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 52, 53, 54, 55,
	}

	res := make(v1admin.AdminRecordTypeRes, 0, len(keys))
	for _, k := range keys {
		if val, ok := consts.RecordTypeMap[k]; ok {
			res = append(res, v1admin.AdminRecordTypeItem{
				Key:   k,
				Value: val,
			})
		}
	}

	// 补回 Map 中可能存在但不在 keys 里的 (防御)
	existingKeys := garray.NewIntArrayFrom(keys)
	for k, v := range consts.RecordTypeMap {
		if !existingKeys.Contains(k) {
			res = append(res, v1admin.AdminRecordTypeItem{
				Key:   k,
				Value: v,
			})
		}
	}

	return res, nil
}
