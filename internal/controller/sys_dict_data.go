package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 字典数据管理
var SysDictData = cDictData{}

type cDictData struct{}

// 根据字典类型查询字典数据,缓存10小时
func (c *cDictData) GetType(ctx context.Context, req *v1.SysDictDataGetTypeReq) (res v1.SysDictDataGetTypeRes, err error) {
	dictDataRes, err := service.SysDictData().GetType(ctx, req.DictType)
	res = gconv.Maps(dictDataRes)
	return
}

// 获取字典数据列表
func (c *cDictData) GetList(ctx context.Context, req *v1.SysDictDataListReq) (res *v1.SysDictDataListRes, err error) {
	in := &model.SysDictDataListInput{}
	gconv.Scan(req, &in)

	dictDataRes, err := service.SysDictData().GetList(ctx, *in)
	gconv.Scan(dictDataRes, &res)
	return
}

// 获取单个字典数据信息
func (c *cDictData) GetOne(ctx context.Context, req *v1.SysDictDataOneReq) (res *v1.SysDictDataOneRes, err error) {
	dictDataRes, err := service.SysDictData().GetOne(ctx, model.SysDictDataOneInput{
		DictCode: int64(req.DictCode),
	})
	gconv.Scan(dictDataRes, &res)
	return
}

// 新建字典数据
func (c *cDictData) Create(ctx context.Context, req *v1.SysDictDataCreateReq) (res *v1.SysDictDataCreateRes, err error) {
	in := &model.SysDictDataCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysDictData().Create(ctx, *in)
	return
}

// 更新字典数据
func (c *cDictData) Update(ctx context.Context, req *v1.SysDictDataUpdateReq) (res *v1.SysDictDataUpdateRes, err error) {
	in := &model.SysDictDataUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysDictData().Update(ctx, *in)
	return
}

// 删除字典数据
func (c *cDictData) Delete(ctx context.Context, req *v1.SysDictDataDeleteReq) (res *v1.SysDictDataDeleteRes, err error) {
	in := &model.SysDictDataDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysDictData().Delete(ctx, *in)
	return
}
