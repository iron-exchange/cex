package codes

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	// 定义业务错误码 (与原 Java 系统 200 / 500 / 401 约定对齐)
	Success     = gcode.New(200, "Success", nil)
	Failed      = gcode.New(500, "Failed", nil)
	ClientError = gcode.New(400, "Bad Request", nil)

	// 认证相关错误码
	Unauthorized = gcode.New(401, "Unauthorized", nil)
	Forbidden    = gcode.New(403, "Forbidden", nil)

	// User 系统错误码 (以 10XX 开头)
	UserNotFound = gcode.New(1001, "User Not Found", nil)
	UserDisabled = gcode.New(1002, "User account is disabled", nil)

	// Asset/Funding 系统错误码 (以 20XX 开头)
	BalanceNotEnough = gcode.New(2001, "Insufficient Balance", nil)
	AssetLocked      = gcode.New(2002, "Asset is locked", nil)
)
