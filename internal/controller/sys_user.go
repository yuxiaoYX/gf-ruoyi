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

// 获取用户列表
func (c *cUser) GetList(ctx context.Context, req *v1.SysUserListReq) (res *v1.SysUserListRes, err error) {
	in := &model.SysUserListInput{}
	gconv.Scan(req, &in)

	userRes, err := service.SysUser().GetList(ctx, *in)
	gconv.Scan(userRes, &res)
	return
}

// 获取单个用户信息
func (c *cUser) GetOne(ctx context.Context, req *v1.SysUserOneReq) (res *v1.SysUserOneRes, err error) {
	// res = &v1.SysUserOneRes{}
	userRes, err := service.SysUser().GetOne(ctx, model.SysUserOneInput{
		UserId: req.UserId,
	})
	gconv.Scan(userRes, &res)
	return
}

// 新建用户
func (c *cUser) Create(ctx context.Context, req *v1.SysUserCreateReq) (res *v1.SysUserCreateRes, err error) {
	in := &model.SysUserCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Create(ctx, *in)
	return
}

// 更新用户
func (c *cUser) Update(ctx context.Context, req *v1.SysUserUpdateReq) (res *v1.SysUserUpdateRes, err error) {
	in := &model.SysUserUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Update(ctx, *in)
	return
}

// 删除用户
func (c *cUser) Delete(ctx context.Context, req *v1.SysUserDeleteReq) (res *v1.SysUserDeleteRes, err error) {
	in := &model.SysUserDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Delete(ctx, *in)
	return
}
