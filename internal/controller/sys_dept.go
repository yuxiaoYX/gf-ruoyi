package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 部门管理
var SysDept = cDept{}

type cDept struct{}

// 获取部门列表
func (c *cDept) GetList(ctx context.Context, req *v1.SysDeptListReq) (res v1.SysDeptListRes, err error) {
	in := &model.SysDeptListInput{}
	gconv.Scan(req, &in)
	deptRes, err := service.SysDept().GetList(ctx, *in)
	res = gconv.Maps(deptRes)
	return
}

// 获取单个部门信息
func (c *cDept) GetOne(ctx context.Context, req *v1.SysDeptOneReq) (res *v1.SysDeptOneRes, err error) {
	deptRes, err := service.SysDept().GetOne(ctx, model.SysDeptOneInput{
		DeptId: req.DeptId,
	})
	gconv.Scan(deptRes, &res)
	return
}

// 新建部门
func (c *cDept) Create(ctx context.Context, req *v1.SysDeptCreateReq) (res *v1.SysDeptCreateRes, err error) {
	in := &model.SysDeptCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysDept().Create(ctx, *in)
	return
}

// 更新部门
func (c *cDept) Update(ctx context.Context, req *v1.SysDeptUpdateReq) (res *v1.SysDeptUpdateRes, err error) {
	in := &model.SysDeptUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysDept().Update(ctx, *in)
	return
}

// 删除部门
func (c *cDept) Delete(ctx context.Context, req *v1.SysDeptDeleteReq) (res *v1.SysDeptDeleteRes, err error) {
	in := &model.SysDeptDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysDept().Delete(ctx, *in)
	return
}

// 查询部门下拉树结构
func (c *cDept) Treeselect(ctx context.Context, req *v1.SysDeptTreeselectReq) (res v1.SysDeptTreeselectRes, err error) {
	res.Depts, err = service.SysDept().Treeselect(ctx)
	return
}
