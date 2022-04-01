package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 用户管理
var SysUser = cUser{}

type cUser struct{}

// 用户注销登录,删除数据库中的在线用户就可以了
func (c *cUser) Logout(ctx context.Context, req *v1.SysUserLogoutReq) (res *v1.SysUserLogoutRes, err error) {
	in := &model.SysUserOnlineDeleteInput{}
	_ = gconv.Struct(service.Context().Get(ctx).Data, in)
	err = service.SysUserOnline().Delete(ctx, *in)
	return
}

// 获取单个用户信息
func (c *cUser) GetOne(ctx context.Context, req *v1.SysUserOneReq) (res *v1.SysUserOneRes, err error) {
	res = &v1.SysUserOneRes{}
	userRes, err := service.SysUser().GetOne(ctx, model.SysUserOneInput{
		UserId: req.UserId,
	})
	gconv.Struct(userRes, &res)
	return
}
