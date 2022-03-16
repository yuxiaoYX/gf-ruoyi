package controller

import (
	"context"
	"fmt"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"
)

// 用户管理
var SysUser = cUser{}

type cUser struct{}

// 用户注销登录,删除数据库中的在线用户就可以了
func (h *cUser) Logout(ctx context.Context) (err error) {
	fmt.Println(service.Context().Get(ctx).Data)
	err = service.SysUserOnline().Delete(ctx, model.SysUserOnlineDeleteInput{Token: "11111"})
	return
}
