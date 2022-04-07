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
					service.Middleware().CORS,
					service.Middleware().ResponseHandler,
				)
				// 登录和注销
				group.Bind(
					controller.Login.Login,
					controller.Login.Logout,
				)
				// token验证，并把用户信息和角色字段列表保存到上下文中
				group.Middleware(service.Middleware().TokenAuth)
				// 登录后获取用户信息和路由
				group.Bind(
					controller.Login.GetInfo,
					controller.Login.GetRouters,
				)
				// 权限效验
				group.Middleware(service.Middleware().Auth)

				group.Bind(
					controller.SysUser,
					controller.SysRole,
					controller.SysMenu,
				)
			})
			s.Run()
			return nil
		},
	}
)
