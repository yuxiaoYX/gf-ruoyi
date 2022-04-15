package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 登录日志管理
var SysLoginLog = cLoginLog{}

type cLoginLog struct{}

// 获取登录日志列表
func (c *cLoginLog) GetList(ctx context.Context, req *v1.SysLoginLogListReq) (res *v1.SysLoginLogListRes, err error) {
	in := &model.SysLoginLogListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysLoginLog().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 删除登录日志
func (c *cLoginLog) Delete(ctx context.Context, req *v1.SysLoginLogDeleteReq) (res *v1.SysLoginLogDeleteRes, err error) {
	in := &model.SysLoginLogDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysLoginLog().Delete(ctx, *in)
	return
}

// 清空登录日志
func (c *cLoginLog) Clean(ctx context.Context, req *v1.SysLoginLogCleanReq) (res *v1.SysLoginLogCleanRes, err error) {
	err = service.SysLoginLog().Clean(ctx)
	return
}
