package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Response 统一响应数据结构体
type Response struct {
	Code    int         `json:"code"    summary:"状态码"`
	Message string      `json:"msg"     summary:"提示信息"`
	Data    interface{} `json:"data"    summary:"返回数据"`
}

// HandlerResponse 拦截器，统一处理返回的 JSON 结构
func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，则不再处理
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()

		// 安全防护：屏蔽内部框架级别的错误（例如：50 内部错误, 52 数据库操作错误等）
		c := code.Code()
		if (c >= 50 && c <= 59) || c == -1 {
			// 在服务端控制台真实打印包含堆栈的详细原始错
			g.Log().Errorf(r.Context(), "[全局捕获] 内部严重异常: %+v", err)
			msg = "系统繁忙，请稍后再试"
			// 统一对外输出 500 代表失败，防止客户端不认识 52 等内部生僻码
			code = gcode.New(500, msg, nil)
		}
	} else {
		if r.Response.Status > 0 && r.Response.Status != 200 {
			switch r.Response.Status {
			case 404:
				code = gcode.CodeNotFound
				msg = "Not Found"
			case 403:
				code = gcode.CodeNotAuthorized
				msg = "Forbidden"
			case 401:
				code = gcode.CodeNotAuthorized
				msg = "Unauthorized"
			default:
				code = gcode.CodeUnknown
				msg = "Unknown Error"
			}
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	// 将内置码转换输出格式：如果是 GoFrame 内置正确的 CodeOK(0)，转换给前端为 200
	outputCode := code.Code()
	if outputCode == 0 {
		outputCode = 200
		if msg == "" {
			msg = "success"
		}
	} else if outputCode == -1 {
		// 内部异常 默认 500
		outputCode = 500
	}

	r.Response.WriteJson(Response{
		Code:    outputCode,
		Message: msg,
		Data:    res,
	})
}
