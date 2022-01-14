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
			s.Group("/api", func(group *ghttp.RouterGroup) {
				handler.GfToken.Middleware(ctx, group)
				group.Middleware(
					service.Middleware.Ctx,
					service.Middleware.ResponseHandler,
				)

				group.Middleware(service.Middleware.TokenAuth)
				group.Bind(
					handler.SysUser, // 用户管理
				)
			})
			s.Run()
			return nil
		},
	}
)
