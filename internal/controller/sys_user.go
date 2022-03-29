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
func (h *cUser) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	in := &model.SysUserOnlineDeleteInput{}
	_ = gconv.Struct(service.Context().Get(ctx).Data, in)
	err = service.SysUserOnline().Delete(ctx, *in)
	return
}
