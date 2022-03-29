package cmd

import (
	"context"
	"gf-ruoyi/internal/controller"
	"gf-ruoyi/internal/service"

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
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)

				group.Bind(controller.Login)
				group.Middleware(
					service.Middleware().TokenAuth,
				)

				group.Bind(
					controller.Hello,
					controller.SysUser,
				)
			})
			s.Run()
			return nil
		},
	}
)
