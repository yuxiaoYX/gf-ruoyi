package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 操作日志管理
var SysOperLog = cOperLog{}

type cOperLog struct{}

// 获取操作日志列表
func (c *cOperLog) GetList(ctx context.Context, req *v1.SysOperLogListReq) (res *v1.SysOperLogListRes, err error) {
	in := &model.SysOperLogListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysOperLog().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 删除操作日志
func (c *cOperLog) Delete(ctx context.Context, req *v1.SysOperLogDeleteReq) (res *v1.SysOperLogDeleteRes, err error) {
	in := &model.SysOperLogDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysOperLog().Delete(ctx, *in)
	return
}

// 清空操作日志
func (c *cOperLog) Clean(ctx context.Context, req *v1.SysOperLogCleanReq) (res *v1.SysOperLogCleanRes, err error) {
	err = service.SysOperLog().Clean(ctx)
	return
}
