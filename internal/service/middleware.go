package service

import (
	"gf-RuoYi/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件管理服务
var (
	Middleware = serviceMidleware{}
)

type serviceMidleware struct {
	LoginUrl string //登录路由地址,不用可以删除
}

// 自定义上下文对象
func (s *serviceMidleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		// Session: r.Session,
		Data: make(g.Map),
	}
	Context.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
