package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 定时任务管理
var SysJob = cJob{}

type cJob struct{}

// 获取定时任务列表
func (c *cJob) GetList(ctx context.Context, req *v1.SysJobListReq) (res *v1.SysJobListRes, err error) {
	in := &model.SysJobListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysJob().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 获取单个定时任务信息
func (c *cJob) GetOne(ctx context.Context, req *v1.SysJobOneReq) (res *v1.SysJobOneRes, err error) {
	roleRes, err := service.SysJob().GetOne(ctx, model.SysJobOneInput{
		JobId: req.JobId,
	})
	gconv.Scan(roleRes, &res)
	return
}

// 新建定时任务
func (c *cJob) Create(ctx context.Context, req *v1.SysJobCreateReq) (res *v1.SysJobCreateRes, err error) {
	in := &model.SysJobCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysJob().Create(ctx, *in)
	return
}

// 更新定时任务
func (c *cJob) Update(ctx context.Context, req *v1.SysJobUpdateReq) (res *v1.SysJobUpdateRes, err error) {
	in := &model.SysJobUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysJob().Update(ctx, *in)
	return
}

// 删除定时任务
func (c *cJob) Delete(ctx context.Context, req *v1.SysJobDeleteReq) (res *v1.SysJobDeleteRes, err error) {
	in := &model.SysJobDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysJob().Delete(ctx, *in)
	return
}
