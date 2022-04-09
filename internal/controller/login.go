package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"
	"gf-ruoyi/utility/utils"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mssola/user_agent"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

// 用户登录，并写入token
// TODO 考虑改成jwt
func (c *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	// service用户名密码验证,并返回token
	out, err := service.SysUser().Login(ctx, model.SysUserLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 保存用户在线状态到数据库
	token := guid.S([]byte(out.UserName))
	r := g.RequestFromCtx(ctx)
	userAgent := r.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	explorer, _ := ua.Browser()
	service.SysUserOnline().Create(ctx, model.SysUserOnlineCreateInput{
		Token:    token,
		UserId:   int(out.UserId),
		UserName: out.UserName,
		Ip:       utils.GetClientIp(r),
		Os:       ua.OS(),
		Explorer: explorer,
	})

	res.Token = token
	return
}

// 用户注销登录,删除数据库中的在线用户就可以了
func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := g.RequestFromCtx(ctx)
	authHeader := r.Request.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return
		} else if parts[1] == "" || parts[1] == "undefined" {
			return
		}
		token := parts[1]
		err = service.SysUserOnline().Delete(ctx, model.SysUserOnlineDeleteInput{Token: token})
	}
	return
}

// 登录后获取用户信息
func (c *cLogin) GetInfo(ctx context.Context, req *v1.LoginUserInfoReq) (res v1.LoginUserInfoRes, err error) {
	// 从上下文，获取用户信息
	userEntity := service.Context().Get(ctx).User
	gconv.Scan(userEntity, &res.User)
	// 从上下文，获取角色权限字符
	res.Roles = service.Context().Get(ctx).Roles.RoleNames
	// 获取菜单权限标识
	roleIds := gconv.Ints(service.Context().Get(ctx).Roles.RoleIds)
	// 角色id为1的，拥有全部权限
	for _, i := range roleIds {
		if i == 1 {
			res.Permissions = []string{"*/*/*"}
			return
		}
	}
	// 取非管理员角色的所有权限
	menuFields, _ := service.SysRoleMenu().GetFieldList(ctx, roleIds)
	res.Permissions = menuFields.Perms
	return
}

// 登录后获取用户路由表
func (c *cLogin) GetRouters(ctx context.Context, req *v1.LoginUserRouterReq) (res v1.LoginUserRouterRes, err error) {
	roleIds := gconv.Ints(service.Context().Get(ctx).Roles.RoleIds)
	res, err = service.SysRoleMenu().GetTreeRoute(ctx, roleIds)
	return
}
