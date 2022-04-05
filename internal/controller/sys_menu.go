package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 菜单管理
var SysMenu = cMenu{}

type cMenu struct{}

// 获取菜单列表
func (c *cMenu) GetList(ctx context.Context, req *v1.SysMenuListReq) (res *v1.SysMenuListRes, err error) {
	in := &model.SysMenuListInput{}
	gconv.Scan(req, &in)

	menuRes, err := service.SysMenu().GetList(ctx, *in)
	gconv.Scan(menuRes, &res)
	return
}

// 获取单个菜单信息
func (c *cMenu) GetOne(ctx context.Context, req *v1.SysMenuOneReq) (res *v1.SysMenuOneRes, err error) {
	menuRes, err := service.SysMenu().GetOne(ctx, model.SysMenuOneInput{
		MenuId: req.MenuId,
	})
	gconv.Scan(menuRes, &res)
	return
}

// 新建菜单
func (c *cMenu) Create(ctx context.Context, req *v1.SysMenuCreateReq) (res *v1.SysMenuCreateRes, err error) {
	in := &model.SysMenuCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysMenu().Create(ctx, *in)
	return
}

// 更新菜单
func (c *cMenu) Update(ctx context.Context, req *v1.SysMenuUpdateReq) (res *v1.SysMenuUpdateRes, err error) {
	in := &model.SysMenuUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysMenu().Update(ctx, *in)
	return
}

// 删除菜单
func (c *cMenu) Delete(ctx context.Context, req *v1.SysMenuDeleteReq) (res *v1.SysMenuDeleteRes, err error) {
	in := &model.SysMenuDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysMenu().Delete(ctx, *in)
	return
}
