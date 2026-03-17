package v1

import "github.com/gogf/gf/v2/frame/g"

type AppUserInfo struct {
	SearchValue      interface{} `json:"searchValue"`
	CreateBy         interface{} `json:"createBy"`
	CreateTime       string      `json:"createTime"`
	UpdateBy         interface{} `json:"updateBy"`
	UpdateTime       string      `json:"updateTime"`
	Remark           interface{} `json:"remark"`
	UserId           int64       `json:"userId"`
	IsTest           int         `json:"isTest"`
	Code             interface{} `json:"code"`
	LoginName        string      `json:"loginName"`
	Email            string      `json:"email"`
	LoginPassword    string      `json:"loginPassword"`
	Address          string      `json:"address"`
	WalletType       string      `json:"walletType"`
	Status           int         `json:"status"`
	TotalAmount      float64     `json:"totleAmont"`
	RechargeAmount   float64     `json:"rechargeAmont"`
	Buff             int         `json:"buff"`
	AppParentIds     string      `json:"appParentIds"`
	AppParentNames   string      `json:"appParentNames"`
	AdminParentIds   string      `json:"adminParentIds"`
	AdminParentNames string      `json:"adminParentNames"`
	ActiveCode       string      `json:"activeCode"`
	RegisterIp       string      `json:"registerIp"`
	Host             string      `json:"host"`
	Phone            string      `json:"phone"`
	Level            int         `json:"level"`
	IsFreeze         string      `json:"isFreeze"` // 线上返回是字符串 "1"
	IsBlack          interface{} `json:"isBlack"`
	SignType         interface{} `json:"signType"`
	Flag             interface{} `json:"flag"`
	ProductId        interface{} `json:"productId"`
	WinNum           int         `json:"winNum"`
	LoseNum          int         `json:"loseNum"`
	Credits          interface{} `json:"credits"`
}

// GetAppUserListReq 获取玩家用户列表
type GetAppUserListReq struct {
	g.Meta         `path:"/user/list" tags:"AdminUser" method:"get" summary:"获取玩家用户列表"`
	PageNum        int    `json:"pageNum" d:"1"`
	PageSize       int    `json:"pageSize" d:"10"`
	UserId         int64  `json:"userId"`
	LoginName      string `json:"loginName"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Status         string `json:"status"`
	IsTest         string `json:"isTest"`
	IsFreeze       string `json:"isFreeze"`
	IsBlack        string `json:"isBlack"`
	AdminParentIds string `json:"adminParentIds"`
	AppParentIds   string `json:"appParentIds"`
}

type GetAppUserListRes struct {
	Total int           `json:"total"`
	Rows  []AppUserInfo `json:"data"`
}

// FreezeUserReq 冻结/解冻用户
type FreezeUserReq struct {
	g.Meta `path:"/user/freeze" tags:"AdminUser" method:"post" summary:"冻结或解冻玩家"`
	UserId int64 `json:"userId" v:"required#用户ID不能为空"`
	Status int   `json:"status" v:"required|in:0,1#状态只能为 0(正常) 或 1(冻结)"`
}

type FreezeUserRes struct{}

// UpdateUserParentReq 修改用户代理层级
type UpdateUserParentReq struct {
	g.Meta         `path:"/user/updateParent" tags:"AdminUser" method:"post" summary:"修改用户代理线"`
	UserId         int64  `json:"userId" v:"required#用户ID不能为空"`
	AdminParentIds string `json:"adminParentIds" v:"required#代理上级ID不能为空"`
}

type UpdateUserParentRes struct{}

// UpdateAppUserReq 修改玩家基本信息和状态
type UpdateAppUserReq struct {
	g.Meta         `path:"/user" tags:"AdminUser" method:"put" summary:"修改玩家基本信息和状态"`
	UserId         *int64  `json:"userId" v:"required#用户ID不能为空"`
	LoginName      *string `json:"loginName"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email"`
	Address        *string `json:"address"`
	Level          *int    `json:"level"`
	Remark         *string `json:"remark"`
	Status         *int    `json:"status"`
	IsFreeze       *string `json:"isFreeze"`
	Buff           *int    `json:"buff"`
	AdminParentIds *string `json:"adminParentIds"`
	AppParentIds   *string `json:"appParentIds"`
	WinNum         *int    `json:"winNum"`
	LoseNum        *int    `json:"loseNum"`
	Credits        *int    `json:"credits"`
}

type UpdateAppUserRes struct{}

// UserBonusReq 人工上下分 / 赠送彩金 通用结构体
type UserBonusReq struct {
	UserId   int64   `json:"userId" v:"required#用户ID不能为空"`
	Amount   float64 `json:"amount" v:"required#操作金额不能为空"`
	Symbol   string  `json:"symbol" v:"required#币种不能为空"`
	Type     int     `json:"type" v:"required|in:1,2#资产类型只能为 1(平台资产) 或 2(合约资产)"`
	GiveType string  `json:"giveType" v:"required|in:0,1#上分扣分标识只能为 0 或 1" dc:"0-上分/赠送, 1-下分/扣减"`
	Remark   string  `json:"remark"`
}

// SendBonusReq 赠送彩金/扣减彩金 (附带虚假充提订单)
type SendBonusReq struct {
	g.Meta `path:"/user/sendBous" tags:"AdminUser" method:"post" summary:"赠送/扣减彩金"`
	UserBonusReq
}

type SendBonusRes struct{}

// SubUserAmountReq 人工上下分 (仅改变资产并记流水)
type SubUserAmountReq struct {
	g.Meta `path:"/user/subAmount" tags:"AdminUser" method:"post" summary:"人工上/下分"`
	UserBonusReq
}

type SubUserAmountRes struct{}

// UpdateUserAppIdsReq 修改玩家归属代理
type UpdateUserAppIdsReq struct {
	g.Meta      `path:"/user/updateUserAppIds" tags:"AdminUser" method:"put" summary:"修改玩家所属后台上级代理"`
	AppUserId   int64 `json:"appUserId" v:"required#玩家ID不能为空"`
	AgentUserId int64 `json:"agentUserId" v:"required#目标代理ID不能为空"`
}

type UpdateUserAppIdsRes struct{}

// AuditUserRealNameReq 审核玩家实名认证
type AuditUserRealNameReq struct {
	g.Meta `path:"/user/realName" tags:"AdminUser" method:"post" summary:"审核玩家实名认证(通过/拒绝)"`
	UserId int64  `json:"userId" v:"required#用户ID不能为空"`
	Flag   string `json:"flag" v:"required|in:1,2,3,4#审核标识只能是 1,2,3,4"`
}

type AuditUserRealNameRes struct{}

// ResetUserRealNameReq 重置实名认证
type ResetUserRealNameReq struct {
	g.Meta    `path:"/user/reSetRealName" tags:"AdminUser" method:"post" summary:"重置实名认证(打回原形)"`
	UserId    int64  `json:"userId" v:"required#用户ID不能为空"`
	ReSetFlag string `json:"reSetFlag" v:"required|in:1,2#重置标识只能是 1 或 2"`
}

type ResetUserRealNameRes struct{}

// UpdateUserRealNameReq 清理/重置玩家实名关联主表数据
type UpdateUserRealNameReq struct {
	g.Meta `path:"/user/updateRealName" tags:"AdminUser" method:"post" summary:"擦除实名认证数据"`
	UserId int64 `json:"userId" v:"required#用户ID不能为空"`
}

type UpdateUserRealNameRes struct{}
