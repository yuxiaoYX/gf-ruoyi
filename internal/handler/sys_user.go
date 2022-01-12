package handler

import (
	"context"
	"gf-RuoYi/apiv1"
)

// 用户管理
var SysUser = handlerUser{}

type handlerUser struct{}

// 用户注销登录,删除数据库中的在线用户就可以了
// func (h *handlerUser) Logout(ctx context.Context, req *apiv1.SysUserLogoutReq) (res *apiv1.SysUserLogoutRes, err error) {
// 	err = service.SysUserOnline.Delete(ctx, model.SysUserOnlineDeleteInput{Token: service.Context.Get(ctx).User.Token})
// 	return
// }

// 获取登录的用户信息
func (h *handlerUser) GetInfo(ctx context.Context, req *apiv1.SysUserInfoReq) (res *apiv1.SysUserInfoRes, err error) {
	// res = &apiv1.SysUserInfoRes{}
	// res.User, err = service.SysUser.GetInfo(ctx, model.SysUserGetInfoInput{UserId: service.Context.Get(ctx).User.UserId})

	// userRoleList, err := SysUserRole.GetByIdList(ctx, model.SysUserRoleListInput{UserId: int(in.UserId)})
	// SysRole.GetIdsList(ctx, model.SysRoleGetIdsInput{RoleIds: userRoleList.RoleIds})
	return
}
