package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// HandlerResponse todo 实现默认json返回
func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 有buffer且http状态码不为500时(考虑 panic)直接返回，默认有自定义返回数据
	if r.Response.BufferLength() > 0 && r.Response.Status != http.StatusInternalServerError {
		return
	}

	var (
		msg      string
		ctx      = r.Context()
		err      = r.GetError()
		res      = r.GetHandlerResponse()
		code     = gerror.Code(err)
		httpCode int
	)

	if err == nil && r.Response.Status == http.StatusOK {
		code = gcode.CodeOK
		httpCode = 200
		responseJson(r, httpCode, code, msg, res)
		return
	}
	if err != nil && r.Response.Status == http.StatusOK {
		code = gcode.CodeOK
		httpCode = 200
		responseJson(r, httpCode, code, msg, res)
		return
	}
	if err != nil {
		if code == gcode.CodeNil { // 未定义code的错误，默认业务级一般错误
			code = gcode.CodeInternalError
			httpCode = 400
		} else {
			if code == gcode.CodeInternalError {
				httpCode = 500
			} else {
				httpCode = code.Code()
			}
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
			httpCode = 404
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
			httpCode = 401
		default:
			code = gcode.CodeUnknown
			httpCode = 500
		}
	} else {
		code = gcode.CodeOK
		httpCode = 200
	}
	r.Response.WriteStatus(httpCode)
	r.Response.ClearBuffer()
	internalErr := r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
	if internalErr != nil {
		glog.Errorf(ctx, `%+v`, internalErr)
	}
}

func responseJson(r *ghttp.Request, httpCode int, code gcode.Code, msg string, res interface{}) {
	r.Response.WriteStatus(httpCode)
	r.Response.ClearBuffer()
	internalErr := r.Response.WriteJsonExit(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
	if internalErr != nil {
		glog.Errorf(r.Context(), `%+v`, internalErr)
	}
}
