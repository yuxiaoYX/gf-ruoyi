package service

import (
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/model/entity"
	"gf-ruoyi/utility/response"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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

// 简单token中间件
func (s *sMiddleware) TokenAuth(r *ghttp.Request) {
	// 获取token，如果token有时效，可以做刷新令牌
	authHeader := r.Request.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.JsonExit(r, 1, "未登录或非法访问!")
		} else if parts[1] == "" {
			response.JsonExit(r, 1, "未登录或非法访问!")
		}
		token := parts[1]
		// 设置token到上下文信息中
		Context().SetData(r.Context(), g.Map{"token": token})
		// 验证token是否有效
		onlineInfo, _ := SysUserOnline().GetToken(r.Context(), token)
		if onlineInfo == nil {
			response.JsonExit(r, 1, "您的帐户异地登陆或令牌失效!")
			// r.SetError(gerror.New("您的帐户异地登陆或令牌失效!"))
		}
		// 设置用户信息到上下文
		userEntity, _ := SysUser().GetOne(r.Context(), model.SysUserOneInput{UserId: uint(onlineInfo.UserId)})
		var ctxUser *entity.SysUser
		gconv.Struct(userEntity, &ctxUser)
		Context().SetUser(r.Context(), ctxUser)
		// 设置角色字段列表到上下文
		roleFields, _ := SysUserRole().GetFieldList(r.Context(), userEntity.UserId)
		var ctxRoles *model.ContextRoles
		gconv.Struct(roleFields, &ctxRoles)
		Context().SetRoles(r.Context(), ctxRoles)
		r.Middleware.Next()
	} else {
		response.JsonExit(r, 1, "未登录或非法访问!")
	}
}

// 接口操作权限验证
func (s *sMiddleware) Permissions(r *ghttp.Request) {
	ctx := r.Context()

	roleIds := gconv.Ints(Context().Get(ctx).Roles.RoleIds)
	if len(roleIds) == 0 {
		response.JsonExit(r, 500, "没有角色权限")
	}
	// 角色id为1的，拥有全部权限
	for _, i := range roleIds {
		if i == 1 {
			Permissions := []string{"*:*:*"}
			return
		}
	}
	menuFields, _ := SysRoleMenu().GetFieldList(ctx, roleIds)
	Permissions := menuFields.Perms

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
