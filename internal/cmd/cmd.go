package cmd

import (
	"context"

	"gf-RuoYi/internal/handler"
	"gf-RuoYi/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware.Ctx)
				group.Bind(
					handler.Hello,
				)
				group.Middleware(service.Middleware.TokenAuth)
			})
			s.Run()
			return nil
		},
	}
)
