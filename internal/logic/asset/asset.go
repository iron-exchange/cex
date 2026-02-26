package asset

import (
	"context"
	"errors"
	"time"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/consts"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	goredislib "github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type sAsset struct {
	rs *redsync.Redsync
}

// New 创建独立锁的 Asset 服务层
func New() *sAsset {
	// GF V2 推荐直接从全局配置取参数建立 Redis 独立外锁连接
	addr, _ := g.Cfg().Get(context.Background(), "redis.default.address")
	if addr.IsEmpty() {
		addr = g.NewVar("127.0.0.1:6379")
	}

	client := goredislib.NewClient(&goredislib.Options{
		Addr: addr.String(),
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	return &sAsset{
		rs: rs,
	}
}

// SubAmount 通用资产增减扣款锁保护入口 ("加减款/人工上下分")
func (s *sAsset) SubAmount(ctx context.Context, in *v1.SubAmountReq) (*v1.SubAmountRes, error) {
	if in.Amount == 0 {
		return nil, gerror.NewCode(codes.ClientError, "变动金额不能为0")
	}

	// 1. 获取 Redis 分布式排他锁 (锁住用户+币种维度，防御高频发重放防止表层穿透导致幻读)
	lockKey := consts.RedisAssetLockPrefix + gconv.String(in.UserId) + ":" + in.Symbol
	mutex := s.rs.NewMutex(lockKey, redsync.WithExpiry(time.Duration(consts.LockWatchDogTimeout)*time.Millisecond))

	if err := mutex.Lock(); err != nil {
		return nil, gerror.NewCode(codes.Failed, "获取资金安全处理锁失败，系统繁忙，请重试")
	}
	defer mutex.Unlock()

	var recordId int64
	var finalAmount float64

	// 2. 数据库事务开启 (保证资产更新与流水插入同时成功或失败)
	err := dao.AppAsset.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var asset entity.AppAsset

		// 3. PostgreSQL 行级悲观锁护航 (SELECT ... FOR UPDATE)
		err := dao.AppAsset.Ctx(ctx).TX(tx).
			Where(dao.AppAsset.Columns().UserId, in.UserId).
			Where(dao.AppAsset.Columns().Symbol, in.Symbol).
			LockUpdate().
			Scan(&asset)

		if err != nil {
			// 允许第一次充值时资金账户为空 (触发新建)
			errStr := err.Error()
			if errStr != "sql: no rows in result set" && errStr != "not found" {
				return err
			}
		}

		if asset.Id == 0 {
			if in.Amount < 0 {
				return gerror.NewCode(codes.Failed, "资产账户不存在，无法进行扣款")
			}
			// 新建初始资金为 0 的空资产账户
			newAsset := entity.AppAsset{
				UserId:          in.UserId,
				Symbol:          in.Symbol,
				Type:            "1", // 默认资金大类
				AvailableAmount: 0,
				Amout:           0,
			}
			newId, err := dao.AppAsset.Ctx(ctx).TX(tx).Data(newAsset).InsertAndGetId()
			if err != nil {
				return err
			}
			asset.Id = int(newId)
			asset.AvailableAmount = 0
			asset.Amout = 0
		}

		// 4. 计算余额，全程使用 Decimal 防止产生弱浮点计算溢出（例如 1.000000001 的鬼影子）
		currentAvailable := decimal.NewFromFloat(asset.AvailableAmount)
		changeAmount := decimal.NewFromFloat(in.Amount)
		newAvailable := currentAvailable.Add(changeAmount)

		if newAvailable.IsNegative() {
			return gerror.NewCode(codes.BalanceNotEnough, "余额不足以执行此扣除操作")
		}

		finalAmount, _ = newAvailable.Float64()
		newAmout, _ := decimal.NewFromFloat(asset.Amout).Add(changeAmount).Float64()

		// 5. 更新资产数据
		updateRes, err := dao.AppAsset.Ctx(ctx).TX(tx).
			Where(dao.AppAsset.Columns().Id, asset.Id).
			// 第三个防线 (DB级防超卖): 要求原先加上当前的剩余必定 >= 0，如果没有这重限制，前置并发仍存在 0.0001% 的破防几率
			Where("available_amount + ? >= 0", in.Amount).
			Data(g.Map{
				dao.AppAsset.Columns().AvailableAmount: finalAmount,
				dao.AppAsset.Columns().Amout:           newAmout,
			}).Update()

		if err != nil {
			return err
		}

		affected, _ := updateRes.RowsAffected()
		if affected == 0 {
			return errors.New("并发余额不足扣款失败，数据可能已过期")
		}

		// 6. 严谨落地资金账单 (WalletRecord) 记录前值后值
		record := entity.AppWalletRecord{
			UserId:       in.UserId,
			BeforeAmount: asset.AvailableAmount,
			Amount:       in.Amount, // 正数充，负数扣
			AfterAmount:  finalAmount,
			Symbol:       in.Symbol,
			Type:         in.RecordType,
			Remark:       in.Remark,
		}

		recordRes, err := dao.AppWalletRecord.Ctx(ctx).TX(tx).Data(record).Insert()
		if err != nil {
			return err
		}
		recordId, _ = recordRes.LastInsertId()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.SubAmountRes{
		RecordId:      recordId,
		CurrentAmount: finalAmount,
	}, nil
}
