package service

import (
	"gf-ruoyi/internal/model"
	"gf-ruoyi/utility/code"
	"gf-ruoyi/utility/response"

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

type DefaultHandlerRes struct {
	Code int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err         error
		res         interface{}
		ctx         = r.Context()
		internalErr error
	)

	res, err = r.GetHandlerResponse()
	if err != nil {
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		err.Error()

		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "操作成功", res)
		// if r.IsAjaxRequest() {
		// 	response.JsonExit(r, code.Code(), "", res)
		// } else {
		// 	// 什么都不做，业务API自行处理模板渲染的成功逻辑。
		// }
	}
}
