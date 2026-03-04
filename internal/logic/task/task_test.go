package task

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

// Test_UpdateCodingDaily 测试打码量每日归集计算任务
func Test_UpdateCodingDaily(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		task := New() // 获取 task 实例

		err := task.UpdateCodingDaily(ctx)
		t.AssertNil(err)

		// 可选：加个日志看看执行没
		g.Log().Info(ctx, "Test_UpdateCodingDaily 跑通了")
	})
}

// Test_SettleSecondContract 测试秒合约/期权杀客结算任务
func Test_SettleSecondContract(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		task := New()

		err := task.SettleSecondContract(ctx)
		t.AssertNil(err)

		g.Log().Info(ctx, "Test_SettleSecondContract 跑通了")
	})
}

// Test_SettleCurrencyOrder 测试现货委托单撮合
func Test_SettleCurrencyOrder(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		task := New()

		err := task.SettleCurrencyOrder(ctx)
		t.AssertNil(err)

		g.Log().Info(ctx, "Test_SettleCurrencyOrder 跑通了")
	})
}

// Test_HandleOwnCoinStart 测试盘点发行开始
func Test_HandleOwnCoinStart(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().HandleOwnCoinStart(ctx)
		t.AssertNil(err)
	})
}

// Test_HandleOwnCoinEnd 测试盘点发行结束
func Test_HandleOwnCoinEnd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().HandleOwnCoinEnd(ctx)
		t.AssertNil(err)
	})
}

// Test_QueryCollectionStatus 测试资金归集状态查询
func Test_QueryCollectionStatus(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().QueryCollectionStatus(ctx)
		t.AssertNil(err)
	})
}

// Test_MonitorUsdtAllowed 测试监控链上授权
func Test_MonitorUsdtAllowed(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().MonitorUsdtAllowed(ctx)
		t.AssertNil(err)
	})
}

// Test_CheckContractPosition 测试合约强平监控
func Test_CheckContractPosition(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().CheckContractPosition(ctx)
		t.AssertNil(err)
	})
}

// Test_SettleContractOrder 测试永续合约入场结算
func Test_SettleContractOrder(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().SettleContractOrder(ctx)
		t.AssertNil(err)
	})
}

// Test_SettleFinancial 测试理财派息
func Test_SettleFinancial(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().SettleFinancial(ctx)
		t.AssertNil(err)
	})
}

// Test_SyncMarketTicker 测试大盘K线拉取
func Test_SyncMarketTicker(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()
		err := New().SyncMarketTicker(ctx)
		t.AssertNil(err)
	})
}

// Test_DistributeDefiRate 测试DeFi挖矿派息
func Test_DistributeDefiRate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.TODO()

		// 先检查是否存在
		count, errCount := g.DB().Model("sys_config").Where("config_key", "CURRENCY_PRICE:ETH").Count()
		if errCount != nil {
			g.Log().Error(ctx, "COUNT 检查失败:", errCount)
		}

		if count > 0 {
			_, errUp := g.DB().Model("sys_config").Data(g.Map{"config_value": "3500.00"}).Where("config_key", "CURRENCY_PRICE:ETH").Update()
			if errUp != nil {
				g.Log().Error(ctx, "UPDATE 失败:", errUp)
			}
		} else {
			// sys_config 的 config_id 可能缺少序列生成器自增，手工造一个最大的 ID 防止报错
			maxIdVal, _ := g.DB().GetValue(ctx, "SELECT COALESCE(MAX(config_id), 0) + 1 FROM sys_config")
			maxId := maxIdVal.Int()

			_, errIns := g.DB().Model("sys_config").Data(g.Map{
				"config_id":    maxId,
				"config_name":  "ETH测试价",
				"config_key":   "CURRENCY_PRICE:ETH",
				"config_value": "3500.00",
				"config_type":  "Y",
			}).Insert()
			if errIns != nil {
				g.Log().Error(ctx, "INSERT 失败:", errIns)
			}
		}

		// 校验写入结果
		val, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", "CURRENCY_PRICE:ETH")
		g.Log().Info(ctx, "测试准备，当前 DB 里的 ETH 价格为:", val.String())

		err := New().DistributeDefiRate(ctx)
		t.AssertNil(err)
	})
}
