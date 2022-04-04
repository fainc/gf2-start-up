package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf2-start-up/internal/controller"
	"gf2-start-up/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddlewareDefault(middleware.Response.HandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Bind(
					controller.Hello,
				)
			})
			s.Run()
			return nil
		},
	}
)
