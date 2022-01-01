package service

import (
	"gf-RuoYi/internal/model"
	"gf-RuoYi/utility/response"
	"strings"

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

// 简单token中间件
func (s *serviceMidleware) TokenAuth(r *ghttp.Request) {
	// 我们这里jwt鉴权取头部信息 Authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	token := r.Request.Header.Get("Authorization")
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.JsonExit(r, 1, "未登录或非法访问!")
		// return
	} else if parts[1] == "" {
		response.JsonExit(r, 1, "未登录或非法访问!")
		// return
	}
	token = parts[1]

	// 验证token是否有效
	onlineInfo, _ := SysUserOnline.GetToken(r.Context(), token)
	if onlineInfo == nil {
		response.JsonExit(r, 1, "您的帐户异地登陆或令牌失效!")
		// return
	}
	// 设置username到上下文
	Context.SetUser(r.Context(), &model.ContextUser{UserName: onlineInfo.UserName})
	r.Middleware.Next()
}
