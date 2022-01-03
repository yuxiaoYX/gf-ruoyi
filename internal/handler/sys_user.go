package handler

import (
	"context"
	"gf-RuoYi/apiv1"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service"
)

// 用户管理
var SysUser = handlerUser{}

type handlerUser struct{}

// 用户注销登录,删除数据库中的在线用户就可以了
func (h *handlerUser) Logout(ctx context.Context, req *apiv1.SysUserLogoutReq) (res *apiv1.SysUserLogoutRes, err error) {
	err = service.SysUserOnline.Delete(ctx, model.SysUserOnlineDeleteInput{UserName: service.Context.Get(ctx).User.UserName})
	return
}

// 获取登录的用户信息、角色权限、菜单等
func (h *handlerUser) GetInfo(ctx context.Context, req *apiv1.SysUserInfoReq) (res *apiv1.SysUserInfoRes, err error) {
	res = &apiv1.SysUserInfoRes{}
	_, err = service.SysUser.GetById(ctx, model.SysUserGetNameInput{UserName: service.Context.Get(ctx).User.UserName})
	return
}
