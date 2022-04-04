package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 角色管理
var SysRole = cRole{}

type cRole struct{}

// 获取角色列表
func (c *cRole) GetList(ctx context.Context, req *v1.SysRoleListReq) (res *v1.SysRoleListRes, err error) {
	in := &model.SysRoleListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysRole().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 获取单个角色信息
func (c *cRole) GetOne(ctx context.Context, req *v1.SysRoleOneReq) (res *v1.SysRoleOneRes, err error) {
	roleRes, err := service.SysRole().GetOne(ctx, model.SysRoleOneInput{
		RoleId: req.RoleId,
	})
	gconv.Scan(roleRes, &res)
	return
}

// 新建角色
func (c *cRole) Create(ctx context.Context, req *v1.SysRoleCreateReq) (res *v1.SysRoleCreateRes, err error) {
	in := &model.SysRoleCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysRole().Create(ctx, *in)
	return
}

// 更新角色
func (c *cRole) Update(ctx context.Context, req *v1.SysRoleUpdateReq) (res *v1.SysRoleUpdateRes, err error) {
	in := &model.SysRoleUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysRole().Update(ctx, *in)
	return
}

// 删除角色
func (c *cRole) Delete(ctx context.Context, req *v1.SysRoleDeleteReq) (res *v1.SysRoleDeleteRes, err error) {
	in := &model.SysRoleDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysRole().Delete(ctx, *in)
	return
}
