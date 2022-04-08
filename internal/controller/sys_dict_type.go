package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 字典类型管理
var SysDictType = cDictType{}

type cDictType struct{}

// 获取字典类型列表
func (c *cDictType) GetList(ctx context.Context, req *v1.SysDictTypeListReq) (res *v1.SysDictTypeListRes, err error) {
	in := &model.SysDictTypeListInput{}
	gconv.Scan(req, &in)

	dictTypeRes, err := service.SysDictType().GetList(ctx, *in)
	gconv.Scan(dictTypeRes, &res)
	return
}

// 获取单个字典类型信息
func (c *cDictType) GetOne(ctx context.Context, req *v1.SysDictTypeOneReq) (res *v1.SysDictTypeOneRes, err error) {
	dictTypeRes, err := service.SysDictType().GetOne(ctx, model.SysDictTypeOneInput{
		DictId: req.DictId,
	})
	gconv.Scan(dictTypeRes, &res)
	return
}

// 新建字典类型
func (c *cDictType) Create(ctx context.Context, req *v1.SysDictTypeCreateReq) (res *v1.SysDictTypeCreateRes, err error) {
	in := &model.SysDictTypeCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysDictType().Create(ctx, *in)
	return
}

// 更新字典类型
func (c *cDictType) Update(ctx context.Context, req *v1.SysDictTypeUpdateReq) (res *v1.SysDictTypeUpdateRes, err error) {
	in := &model.SysDictTypeUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysDictType().Update(ctx, *in)
	return
}

// 删除字典类型
func (c *cDictType) Delete(ctx context.Context, req *v1.SysDictTypeDeleteReq) (res *v1.SysDictTypeDeleteRes, err error) {
	in := &model.SysDictTypeDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysDictType().Delete(ctx, *in)
	return
}
