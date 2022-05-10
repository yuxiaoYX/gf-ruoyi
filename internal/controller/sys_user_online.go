package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 在线用户管理
var SysUserOnline = cUserOnline{}

type cUserOnline struct{}

// 获取在线用户列表
func (c *cUserOnline) GetList(ctx context.Context, req *v1.SysUserOnlineListReq) (res *v1.SysUserOnlineListRes, err error) {
	in := &model.SysUserOnlineListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysUserOnline().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 删除在线用户
func (c *cUserOnline) Delete(ctx context.Context, req *v1.SysUserOnlineDeleteReq) (res *v1.SysUserOnlineDeleteRes, err error) {
	err = service.SysUserOnline().Delete(ctx, model.SysUserOnlineDeleteInput{IdList: req.IdList})
	return
}
