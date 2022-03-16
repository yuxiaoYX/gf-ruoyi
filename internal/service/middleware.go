package service

import (
	"gf-ruoyi/internal/model"
	"gf-ruoyi/utility/response"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

// 中间件管理服务
func Middleware() *sMiddleware {
	return &sMiddleware{}
}

// 简单token中间件
func (s *sMiddleware) TokenAuth(r *ghttp.Request) {
	// 获取token，如果token有时效，可以做刷新令牌
	authHeader := r.Request.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// response.JsonExit(r, 1, "未登录或非法访问!")
			gerror.New("未登录或非法访问")
		} else if parts[1] == "" {
			response.JsonExit(r, 1, "未登录或非法访问!")
			gerror.New("未登录或非法访问")
		}
		token := parts[1]

		Context().SetData(r.Context(), g.Map{"token": token})
	}

	r.Middleware.Next()
}

// 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		// Session: r.Session,
		Data: make(g.Map),
	}
	Context().Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)

	if err != nil {
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}

		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "操作成功", res)
	}
}
