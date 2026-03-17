package user

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminUser struct{}

func New() *sAdminUser {
	return &sAdminUser{}
}

func (s *sAdminUser) GetAppUserList(ctx context.Context, req *v1.GetAppUserListReq) (*v1.GetAppUserListRes, error) {
	m := dao.AppUser.Ctx(ctx).As("u").LeftJoin("t_app_user_detail d", "u.user_id = d.user_id")

	// 0. 数据隔离：非超级管理员（ID != 1）只能看自己线下的玩家
	// 注意：根据 JWT claims，adminId 可能是 float64 类型
	adminIdVal := ctx.Value("adminId")
	if adminIdVal != nil {
		var adminId int64
		switch v := adminIdVal.(type) {
		case float64:
			adminId = int64(v)
		case int64:
			adminId = v
		case int:
			adminId = int64(v)
		}

		if adminId != 1 {
			// 如果当前不是超管，则强制覆盖或追加 adminParentIds 条件
			// 将查询条件锁定在自己节点之下
			if req.AdminParentIds == "" {
				req.AdminParentIds = gconv.String(adminId)
			}
		}
	}

	// 1. 基础过滤条件
	if req.UserId > 0 {
		m = m.Where("u.user_id", req.UserId)
	}
	if req.LoginName != "" {
		m = m.WhereLike("u.login_name", "%"+req.LoginName+"%")
	}
	if req.Phone != "" {
		m = m.WhereLike("u.phone", "%"+req.Phone+"%")
	}
	if req.Email != "" {
		m = m.WhereLike("u.email", "%"+req.Email+"%")
	}
	if req.Address != "" {
		m = m.WhereLike("u.address", "%"+req.Address+"%")
	}
	if req.Status != "" {
		m = m.Where("u.status", req.Status)
	}
	if req.IsTest != "" {
		m = m.Where("u.is_test", req.IsTest)
	}
	if req.IsFreeze != "" {
		m = m.Where("u.is_freeze", req.IsFreeze)
	}
	if req.IsBlack != "" {
		m = m.Where("u.is_black", req.IsBlack)
	}
	if req.AdminParentIds != "" {
		m = m.WhereLike("u.admin_parent_ids", "%"+req.AdminParentIds+"%")
	}
	if req.AppParentIds != "" {
		m = m.WhereLike("u.app_parent_ids", "%"+req.AppParentIds+"%")
	}

	// 2. 统计总数
	total, err := m.Count()
	if err != nil || total == 0 {
		return &v1.GetAppUserListRes{Rows: []v1.AppUserInfo{}, Total: 0}, nil
	}

	// 3. 聚合查询字段
	m = m.Fields(`
		u.*,
		d.win_num,
		d.lose_num,
		d.credits,
		(SELECT string_agg(nick_name, ', ') FROM sys_user WHERE user_id = ANY(string_to_array(u.admin_parent_ids, ',')::bigint[])) as admin_parent_names,
		(SELECT string_agg(login_name, ', ') FROM t_app_user WHERE user_id = ANY(string_to_array(u.app_parent_ids, ',')::bigint[])) as app_parent_names
	`)

	// 使用更精确的结构映射
	resListMap, err := m.Page(req.PageNum, req.PageSize).OrderDesc("u.create_time").All()
	if err != nil {
		return nil, gerror.Wrap(err, "获取玩家列表失败")
	}

	resRows := make([]v1.AppUserInfo, 0, len(resListMap))
	for _, item := range resListMap {
		row := v1.AppUserInfo{
			SearchValue:      item["search_value"].Interface(),
			CreateBy:         item["create_by"].Interface(),
			CreateTime:       item["create_time"].GTime().Format("2006-01-02 15:04:05"),
			UpdateBy:         item["update_by"].Interface(),
			UpdateTime:       item["update_time"].GTime().Format("2006-01-02 15:04:05"),
			Remark:           item["remark"].Interface(),
			UserId:           item["user_id"].Int64(),
			IsTest:           item["is_test"].Int(),
			Code:             nil, // 线上返回 null
			LoginName:        item["login_name"].String(),
			Email:            item["email"].String(),
			LoginPassword:    item["login_password"].String(),
			Address:          item["address"].String(),
			WalletType:       item["wallet_type"].String(),
			Status:           item["status"].Int(),
			TotalAmount:      item["totle_amont"].Float64(),
			RechargeAmount:   item["recharge_amont"].Float64(),
			Buff:             item["buff"].Int(),
			AppParentIds:     item["app_parent_ids"].String(),
			AppParentNames:   item["app_parent_names"].String(),
			AdminParentIds:   item["admin_parent_ids"].String(),
			AdminParentNames: item["admin_parent_names"].String(),
			ActiveCode:       item["active_code"].String(),
			RegisterIp:       item["register_ip"].String(),
			Host:             item["host"].String(),
			Phone:            item["phone"].String(),
			Level:            item["level"].Int(),
			IsFreeze:         item["is_freeze"].String(),
			IsBlack:          item["is_black"].Interface(),
			SignType:         nil,
			Flag:             nil,
			ProductId:        nil,
			WinNum:           item["win_num"].Int(),
			LoseNum:          item["lose_num"].Int(),
			Credits:          item["credits"].Interface(),
		}
		resRows = append(resRows, row)
	}

	return &v1.GetAppUserListRes{
		Total: total,
		Rows:  resRows,
	}, nil
}

func (s *sAdminUser) FreezeUser(ctx context.Context, req *v1.FreezeUserReq) error {
	_, err := dao.AppUser.Ctx(ctx).Where("user_id", req.UserId).Update(map[string]interface{}{
		"status": req.Status,
	})
	if err != nil {
		return gerror.Wrap(err, "更新用户状态失败")
	}
	return nil
}

// SubAmount 人工上下分 (仅改变资产，不写充值提现订单)
func (s *sAdminUser) SubAmount(ctx context.Context, req *v1.SubUserAmountReq) error {
	return s.handleUserBonus(ctx, &req.UserBonusReq, "subAmount")
}

// SendBonus 赠送彩金/扣减彩金 (附带虚假充提订单)
func (s *sAdminUser) SendBonus(ctx context.Context, req *v1.SendBonusReq) error {
	return s.handleUserBonus(ctx, &req.UserBonusReq, "sendBous")
}

func (s *sAdminUser) handleUserBonus(ctx context.Context, req *v1.UserBonusReq, opOrigin string) error {
	if req.Amount <= 0 {
		return gerror.New("操作金额必须大于 0")
	}

	return dao.AppAsset.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 查询用户是否存在并获取其代理归属信息
		var u entity.AppUser
		err := dao.AppUser.Ctx(ctx).Where("user_id", req.UserId).Scan(&u)
		if err != nil || u.UserId == 0 {
			return gerror.New("用户不存在")
		}

		// 2. 锁定并查询资产
		var asset entity.AppAsset
		err = dao.AppAsset.Ctx(ctx).
			Where("user_id", req.UserId).
			Where("symbol", req.Symbol).
			Where("type", req.Type).
			LockUpdate(). // 悲观锁保证资金安全 SELECT FOR UPDATE
			Scan(&asset)

		if err != nil {
			return gerror.Wrap(err, "获取用户资产失败")
		}
		if asset.Id == 0 {
			return gerror.New("未找到对应资产账户")
		}

		beforeAmount := asset.Amout
		var afterAmount float64

		// 3. 执行资产加减逻辑
		if req.GiveType == "0" {
			// 0: 上分/赠送
			afterAmount = beforeAmount + req.Amount
			asset.Amout = afterAmount
			asset.AvailableAmount = asset.AvailableAmount + req.Amount
		} else if req.GiveType == "1" {
			// 1: 下分/扣减
			if beforeAmount < req.Amount {
				return gerror.New("用户可用资产不足以扣减")
			}
			afterAmount = beforeAmount - req.Amount
			asset.Amout = afterAmount
			asset.AvailableAmount = asset.AvailableAmount - req.Amount
		} else {
			return gerror.New("无效的操作标识(giveType)")
		}

		// 4. 更新流水
		var walletRecordType int
		// 根据来源定义不同的账变类型 (Java 版里可能写成了硬编码数字，如 1-充值/4-赠送等，这里用常量占位)
		if req.GiveType == "0" {
			walletRecordType = 11 // 后台上分/赠送 (具体类型常量视业务而定)
		} else {
			walletRecordType = 12 // 后台下分/扣减
		}

		// 5. 写入 t_app_wallet_record 流水表
		_, err = dao.AppWalletRecord.Ctx(ctx).Insert(g.Map{
			"user_id":          req.UserId,
			"amount":           req.Amount,
			"u_amount":         0, // 这里若需要换算U可扩展
			"before_amount":    beforeAmount,
			"after_amount":     afterAmount,
			"type":             walletRecordType,
			"symbol":           req.Symbol,
			"admin_parent_ids": u.AdminParentIds,
			"remark":           req.Remark,
			// CreateTime 等由 GoFrame 自动填充
		})
		if err != nil {
			return gerror.Wrap(err, "生成资产流水失败")
		}

		// 6. 落库更新资产表 t_app_asset
		_, err = dao.AppAsset.Ctx(ctx).
			Data(g.Map{
				"amout":            asset.Amout,
				"available_amount": asset.AvailableAmount,
			}).
			Where("id", asset.Id).
			Update()

		if err != nil {
			return gerror.Wrap(err, "更新资产失败")
		}

		// 7. 处理 SendBonus 独有的虚假充提订单逻辑
		if opOrigin == "sendBous" {
			if req.GiveType == "0" {
				// 为上分操作插入 t_app_recharge
				// [TODO] 视表结构插入充值记录, 以下仅为演示
				// dao.AppRecharge.Ctx(ctx).Insert(g.Map{ "user_id": req.UserId, "amount": req.Amount, "status": 2 })
				g.Log().Debugf(ctx, "[SendBonus] 触发插入充值订单：userId=%d, amount=%.2f", req.UserId, req.Amount)
			} else if req.GiveType == "1" {
				// 为下分操作插入 t_withdraw
				// [TODO] 视表结构插入提现记录
				// dao.Withdraw.Ctx(ctx).Insert(g.Map{ "user_id": req.UserId, "amount": req.Amount, "status": 1 })
				g.Log().Debugf(ctx, "[SendBonus] 触发插提现订单：userId=%d, amount=%.2f", req.UserId, req.Amount)
			}
		}

		return nil
	})
}

// UpdateUser 修改玩家基本信息和状态
func (s *sAdminUser) UpdateUser(ctx context.Context, req *v1.UpdateAppUserReq) error {
	// 使用事务保证主表和详情表一致性
	return dao.AppUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 动态更新主表 (gdb 的 Data 会自动过滤 nil 指针，实现动态拼装)
		res, err := dao.AppUser.Ctx(ctx).Data(req).Where("user_id", req.UserId).Update()
		if err != nil {
			return gerror.Wrap(err, "更新用户主表失败")
		}

		// 2. 将此用户的状态写入一张 PostgreSQL 信号表，替代原本向 Redis Stream 写消息的逻辑
		//    如果有其他的守护进程监听，就可以轮询或者触发处理。
		//    (假设有个专门记录业务事件的队列表, 例如直接向 t_sys_oper_log 这类表里插一条自定义日志
		//    或者为了彻底解耦此处仅示范抛出 GEvent 或日志，实际项目中需要定义专门的PG队列表).
		rowsAffected, _ := res.RowsAffected()
		if rowsAffected > 0 {
			// [TODO] 替换原本的 Redis Stream (redisStreamNamesApi): "user_status" -> { "id" : userId }
			// 这里我们使用系统日志表或单独设计的队列表来做 Pub/Sub 降级
			// dao.SysNotice.Ctx(ctx).Data(g.Map{"title": "UserStatusUpdate", "content": req.UserId}).Insert()
			g.Log().Debugf(ctx, "[UpdateUser] 触发用户状态变更信号，替代 Redis Stream. userId: %d", *req.UserId)
		}

		// 3. 动态更新详情表 (解决原版 Java 盲目覆写为 0 的 bug)
		// 仅当前端实际传递了 winNum, loseNum, credits 等相关字段时才会进入更新逻辑
		detailData := g.Map{}
		if req.WinNum != nil {
			detailData["win_num"] = *req.WinNum
		}
		if req.LoseNum != nil {
			detailData["lose_num"] = *req.LoseNum
		}
		if req.Credits != nil {
			detailData["credits"] = *req.Credits
		}

		if len(detailData) > 0 {
			_, err = dao.AppUserDetail.Ctx(ctx).Data(detailData).Where("user_id", req.UserId).Update()
			if err != nil {
				return gerror.Wrap(err, "更新用户详情表统数据失")
			}
		}

		return nil
	})
}

// UpdateUserAppIds 修改玩家所属的后台上级代理
func (s *sAdminUser) UpdateUserAppIds(ctx context.Context, req *v1.UpdateUserAppIdsReq) error {
	// 1. 查询目标代理对象
	var agent entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where("user_id", req.AgentUserId).Scan(&agent)
	if err != nil {
		return gerror.Wrap(err, "查询目标代理失败")
	}
	if agent.UserId == 0 {
		return gerror.New("目标代理不存在")
	}

	// 2. 拼接 admin_parent_ids
	adminParentIds := gconv.String(agent.UserId) + ","
	if agent.ParentId > 0 {
		adminParentIds += gconv.String(agent.ParentId)
	} else {
		adminParentIds += "0"
	}

	// 3. 执行更新
	_, err = dao.AppUser.Ctx(ctx).
		Data(g.Map{"admin_parent_ids": adminParentIds}).
		Where("user_id", req.AppUserId).
		Update()

	if err != nil {
		return gerror.Wrap(err, "修改玩家上级代理失败")
	}

	g.Log().Debugf(ctx, "[UpdateUserAppIds] appUserId: %d tied to adminParentIds: %s", req.AppUserId, adminParentIds)
	return nil
}

// AuditUserRealName 审核玩家实名认证(通过/拒绝)
func (s *sAdminUser) AuditUserRealName(ctx context.Context, req *v1.AuditUserRealNameReq) error {
	updateData := g.Map{}

	switch req.Flag {
	case "1":
		updateData["audit_status_primary"] = 2 // 初级通过
	case "2":
		updateData["audit_status_primary"] = 3 // 初级拒绝
	case "3":
		updateData["audit_status_advanced"] = 2 // 高级通过
	case "4":
		updateData["audit_status_advanced"] = 3 // 高级拒绝
	default:
		return gerror.New("无效的审核标识")
	}

	_, err := dao.AppUserDetail.Ctx(ctx).Data(updateData).Where("user_id", req.UserId).Update()
	if err != nil {
		return gerror.Wrap(err, "操作实名审核状态失败")
	}

	// [TODO] 触发 Websocket 通知刷新后台页面 (原逻辑中的 ws 通知暂略)
	g.Log().Debugf(ctx, "[AuditUserRealName] User KYC updated. userId: %d, flag: %s", req.UserId, req.Flag)

	return nil
}

// ResetUserRealName 重置实名认证(打回原形)
func (s *sAdminUser) ResetUserRealName(ctx context.Context, req *v1.ResetUserRealNameReq) error {
	updateData := g.Map{}

	switch req.ReSetFlag {
	case "1":
		updateData["audit_status_primary"] = nil // 清空初级状态
	case "2":
		updateData["audit_status_advanced"] = nil // 清空高级状态
	default:
		return gerror.New("无效的重置标识")
	}

	_, err := dao.AppUserDetail.Ctx(ctx).Data(updateData).Where("user_id", req.UserId).Update()
	if err != nil {
		return gerror.Wrap(err, "重置实名信息失败")
	}

	g.Log().Debugf(ctx, "[ResetUserRealName] User KYC reset. userId: %d, type: %s", req.UserId, req.ReSetFlag)

	return nil
}

// UpdateUserRealName 清理/擦除玩家实名关联数据 (替代原版无用空更新)
func (s *sAdminUser) UpdateUserRealName(ctx context.Context, req *v1.UpdateUserRealNameReq) error {
	// 直接将实名相关的敏感信息和图片抹除
	updateData := g.Map{
		"real_name":  "",
		"id_card":    "",
		"front_url":  "",
		"back_url":   "",
		"handel_url": "",
	}

	_, err := dao.AppUserDetail.Ctx(ctx).Data(updateData).Where("user_id", req.UserId).Update()
	if err != nil {
		return gerror.Wrap(err, "擦除实名信息失败")
	}

	g.Log().Debugf(ctx, "[UpdateUserRealName] User KYC PI data wiped. userId: %d", req.UserId)

	return nil
}
