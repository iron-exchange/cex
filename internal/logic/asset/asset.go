package asset

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
)

type sAsset struct{}

// New 创建独立锁的 Asset 服务层
func New() *sAsset {
	return &sAsset{}
}

// GetExchangeRate 获取实时汇率 (对 USDT 折合)
func (s *sAsset) GetExchangeRate(ctx context.Context, symbol string) decimal.Decimal {
	upperSymbol := strings.ToUpper(symbol)
	if upperSymbol == "USDT" {
		return decimal.NewFromInt(1)
	}
	// 从 Redis 取出实时行情。统一使用大写键名
	priceStr, _ := g.Redis().Do(ctx, "GET", "CEX:PRICE:"+upperSymbol)
	if !priceStr.IsEmpty() {
		p, errParse := decimal.NewFromString(priceStr.String())
		if errParse == nil && !p.IsZero() {
			return p
		}
	}
	return decimal.NewFromInt(0)
}

// SubAmount 通用资产增减扣款锁保护入口 ("加减款/人工上下分")
func (s *sAsset) SubAmount(ctx context.Context, in *v1.SubAmountReq, callbacks ...func(ctx context.Context, tx gdb.TX) error) (*v1.SubAmountRes, error) {
	var recordId int64
	var finalAmount float64

	symbol := strings.ToLower(in.Symbol)

	// 2. 数据库事务开启 (PG Advisory Lock 必须在事务内执行以实现自动释放)
	err := dao.AppAsset.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error
		// 1. 获取 PostgreSQL 事务级咨询锁 (锁住用户+币种维度)
		// 生成 64位 Hash Key: hash(userId:symbol)
		h := fnv.New64a()
		h.Write([]byte(fmt.Sprintf("%d:%s", in.UserId, symbol)))
		lockKey := int64(h.Sum64())

		_, err = tx.Exec("SELECT pg_advisory_xact_lock(?)", lockKey)
		if err != nil {
			return gerror.NewCode(codes.Failed, "获取资金安全处理锁失败，系统异常")
		}

		var asset entity.AppAsset
		g.Log().Debugf(ctx, "SubAmount Intercept: UserId=%d, Symbol=%s, Amount=%f", in.UserId, symbol, in.Amount)

		// 3. PostgreSQL 行级悲观锁护航 (SELECT ... FOR UPDATE)
		err = dao.AppAsset.Ctx(ctx).TX(tx).
			Where(dao.AppAsset.Columns().UserId, in.UserId).
			Where(dao.AppAsset.Columns().Symbol, symbol).
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
			newAsset := g.Map{
				dao.AppAsset.Columns().UserId:               in.UserId,
				dao.AppAsset.Columns().Symbol:               symbol,
				dao.AppAsset.Columns().Type:                 "1", // 默认资金大类
				dao.AppAsset.Columns().AvailableAmount:      0,
				dao.AppAsset.Columns().Amout:                0,
				dao.AppAsset.Columns().OccupiedAmount:       0,
				dao.AppAsset.Columns().AvailableAmountDaily: 0,
				dao.AppAsset.Columns().CodingVolumeDaily:    0,
			}
			newId, err := dao.AppAsset.Ctx(ctx).TX(tx).Data(newAsset).InsertAndGetId()
			if err != nil {
				return err
			}
			asset.Id = int(newId)
			asset.AvailableAmount = 0
			asset.Amout = 0
			asset.Symbol = symbol
		}

		// 4. 计算余额，全程使用 Decimal 防止产生弱浮点计算溢出（例如 1.000000001 的鬼影子）
		currentAvailable := decimal.NewFromFloat(asset.AvailableAmount)

		var changeAmount decimal.Decimal
		if in.AmountStr != "" {
			var errParse error
			changeAmount, errParse = decimal.NewFromString(in.AmountStr)
			if errParse != nil {
				return gerror.NewCode(codes.ClientError, "高精度金额解析失败")
			}
		} else {
			changeAmount = decimal.NewFromFloat(in.Amount)
		}

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

		// 6. 严谨落地资金账目 (WalletRecord) 记录前值后值
		var uAmount float64
		if in.UAmount != 0 {
			uAmount = in.UAmount
		} else {
			rate := s.GetExchangeRate(ctx, in.Symbol)
			uAmount, _ = changeAmount.Abs().Mul(rate).Float64()
		}

		record := g.Map{
			dao.AppWalletRecord.Columns().UserId:         in.UserId,
			dao.AppWalletRecord.Columns().BeforeAmount:   asset.AvailableAmount,
			dao.AppWalletRecord.Columns().Amount:         changeAmount.Abs().String(), // 记录绝对值，对齐 Java 版
			dao.AppWalletRecord.Columns().AfterAmount:    finalAmount,
			dao.AppWalletRecord.Columns().Symbol:         symbol,
			dao.AppWalletRecord.Columns().Type:           in.RecordType,
			dao.AppWalletRecord.Columns().Remark:         in.Remark,
			dao.AppWalletRecord.Columns().UAmount:        uAmount, // 根据汇率折算的 U 金额
			dao.AppWalletRecord.Columns().CreateBy:       "",
			dao.AppWalletRecord.Columns().UpdateBy:       "",
			dao.AppWalletRecord.Columns().SearchValue:    "",
			dao.AppWalletRecord.Columns().SerialId:       "",
			dao.AppWalletRecord.Columns().AdminParentIds: "",
		}

		recordRes, err := dao.AppWalletRecord.Ctx(ctx).TX(tx).Data(record).Insert()
		if err != nil {
			return err
		}
		recordId, _ = recordRes.LastInsertId()

		// 7. 处理同属一个事务内的回调钩子 (例如更新订单状态，防超卖防并发幂等强检验)
		for _, cb := range callbacks {
			if err := cb(ctx, tx); err != nil {
				return err
			}
		}

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

// AddAmount 单纯的加款代理 (本质复用 SubAmount 底层的双锁)
func (s *sAsset) AddAmount(ctx context.Context, in *v1.SubAmountReq, callbacks ...func(ctx context.Context, tx gdb.TX) error) (*v1.SubAmountRes, error) {
	if in.Amount <= 0 && in.AmountStr == "" {
		return nil, gerror.NewCode(codes.ClientError, "加款金额必须为正数")
	}
	return s.SubAmount(ctx, in, callbacks...)
}

// FreezeAmount 冻结或解冻资金 (Amount 为正数表示冻结，为负数表示解冻)
func (s *sAsset) FreezeAmount(ctx context.Context, in *v1.SubAmountReq) (*v1.SubAmountRes, error) {
	var recordId int64
	var finalAvailable float64

	symbol := strings.ToLower(in.Symbol)

	// 2. 数据库事务开启
	err := dao.AppAsset.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error
		// 1. 获取 PostgreSQL 事务级咨询锁
		h := fnv.New64a()
		h.Write([]byte(fmt.Sprintf("%d:%s", in.UserId, symbol)))
		lockKey := int64(h.Sum64())

		_, err = tx.Exec("SELECT pg_advisory_xact_lock(?)", lockKey)
		if err != nil {
			return gerror.NewCode(codes.Failed, "获取资金安全处理锁失败，系统异常")
		}

		var asset entity.AppAsset
		err = dao.AppAsset.Ctx(ctx).TX(tx).
			Where(dao.AppAsset.Columns().UserId, in.UserId).
			Where(dao.AppAsset.Columns().Symbol, symbol).
			Where(dao.AppAsset.Columns().Type, 1). // 必须是平台普通资产类型
			LockUpdate().
			Scan(&asset)

		if err != nil {
			errStr := err.Error()
			if errStr != "sql: no rows in result set" && errStr != "not found" {
				g.Log().Errorf(ctx, "[资产系统] 查询账户异常 UID:%v Symbol:%v Err:%v", in.UserId, symbol, err)
				return err
			}
		}

		if asset.Id == 0 {
			// 如果账户不存在，说明从未有过该币种，余额自然为 0，冻结操作（非加款）必定失败
			return gerror.NewCode(codes.BalanceNotEnough, "可用余额不足，请先充值")
		}

		// 3. 精度运算
		changeAmount := decimal.NewFromFloat(in.Amount)
		newAvailable := decimal.NewFromFloat(asset.AvailableAmount).Sub(changeAmount)
		newOccupied := decimal.NewFromFloat(asset.OccupiedAmount).Add(changeAmount)

		if newAvailable.IsNegative() {
			return gerror.NewCode(codes.BalanceNotEnough, "可用余额不足以执行此冻结操作")
		}
		if newOccupied.IsNegative() {
			return gerror.NewCode(codes.BalanceNotEnough, "解冻金额超过了当前已冻结的资产")
		}

		finalAvailable, _ = newAvailable.Float64()
		finalOccupied, _ := newOccupied.Float64()

		// 4. 更新资产数据 (使用底层限制保护超卖)
		updateRes, err := dao.AppAsset.Ctx(ctx).TX(tx).
			Where(dao.AppAsset.Columns().Id, asset.Id).
			Where("available_amount - ? >= 0", in.Amount).
			Where("occupied_amount + ? >= 0", in.Amount).
			Data(g.Map{
				dao.AppAsset.Columns().AvailableAmount: finalAvailable,
				dao.AppAsset.Columns().OccupiedAmount:  finalOccupied,
			}).Update()

		if err != nil {
			return err
		}

		affected, _ := updateRes.RowsAffected()
		if affected == 0 {
			return errors.New("并发余额不足冻结失败，数据可能已过期")
		}

		// 5. 落地资金流水 (WalletRecord) 记录账变
		var uAmount float64
		if in.UAmount != 0 {
			uAmount = in.UAmount
		} else {
			rate := s.GetExchangeRate(ctx, in.Symbol)
			uAmount, _ = changeAmount.Abs().Mul(rate).Float64()
		}

		record := g.Map{
			dao.AppWalletRecord.Columns().UserId:       in.UserId,
			dao.AppWalletRecord.Columns().BeforeAmount: asset.AvailableAmount,
			// 这里记录绝对值，对齐 Java 版
			dao.AppWalletRecord.Columns().Amount:         changeAmount.Abs().InexactFloat64(),
			dao.AppWalletRecord.Columns().AfterAmount:    finalAvailable,
			dao.AppWalletRecord.Columns().Symbol:         symbol,
			dao.AppWalletRecord.Columns().Type:           in.RecordType, // 例如 21 代表提现冻结
			dao.AppWalletRecord.Columns().Remark:         in.Remark,
			dao.AppWalletRecord.Columns().UAmount:        uAmount,
			dao.AppWalletRecord.Columns().CreateBy:       "",
			dao.AppWalletRecord.Columns().UpdateBy:       "",
			dao.AppWalletRecord.Columns().SearchValue:    "",
			dao.AppWalletRecord.Columns().SerialId:       "",
			dao.AppWalletRecord.Columns().AdminParentIds: "",
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
		CurrentAmount: finalAvailable,
	}, nil
}

// UnfreezeAmount 解提资金 (本质上是 FreezeAmount 提交负数)
func (s *sAsset) UnfreezeAmount(ctx context.Context, in *v1.SubAmountReq) (*v1.SubAmountRes, error) {
	if in.Amount > 0 {
		in.Amount = -in.Amount
	}
	return s.FreezeAmount(ctx, in)
}
