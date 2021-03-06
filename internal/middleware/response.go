package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Response = cResponse{}
)

type cResponse struct{}

type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func (c *cResponse) NotAuthorized() error {
	return gerror.NewCode(gcode.CodeNotAuthorized)
}

func (c *cResponse) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 自定义返回数据格式
	format := r.GetCtxVar("response_format").String()
	g.Dump(format)
	if r.Response.BufferLength() > 0 && format == "custom" {
		return
	}

	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	g.Dump(res)
	if err == nil && r.Response.Status == 200 {
		c.responseJson(r, r.Response.Status, 1, "请求成功", res) // success
	}
	if err != nil && r.Response.Status == 200 {
		switch code.Code() {
		case 61:
			c.responseJson(r, 401, -1, err.Error(), res) // 401
		default:
			c.responseJson(r, r.Response.Status, -1, err.Error(), res) // business error
		}
	}
	if r.Response.Status == 404 {
		c.responseJson(r, r.Response.Status, -1, "未找到该资源", res) // 404
	}
	if err != nil && r.Response.Status == 500 {
		c.responseJson(r, r.Response.Status, -1, "服务器内部错误", res) // 500 has err
	}
	if err == nil && r.Response.Status == 500 {
		c.responseJson(r, r.Response.Status, -1, "未知服务器错误", res) // 500 unknown err
	}

}

func (c *cResponse) responseJson(r *ghttp.Request, httpCode int, code int, msg string, res interface{}) {
	r.Response.WriteStatus(httpCode)
	r.Response.ClearBuffer()

	internalErr := r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    res,
	})
	if internalErr != nil {
		glog.Errorf(r.Context(), `%+v`, internalErr)
	}
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
	r.Exit()

}
