package service

import (
	"context"
	"errors"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDictType struct{}

// 字典类型管理服务
func SysDictType() *sDictType {
	return &sDictType{}
}

// 获取字典类型列表
func (s *sDictType) GetList(ctx context.Context, in model.SysDictTypeListInput) (out model.SysDictTypeListOutput, err error) {
	m := dao.SysDictType.Ctx(ctx).OmitEmpty().Where(g.Map{
		"dict_name": in.DictName,
		"dict_type": in.DictType,
		"status":    in.Status,
	})
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
	return
}

// 获取字典类型详细信息,缓存10小时
func (s *sDictType) GetOne(ctx context.Context, in model.SysDictTypeOneInput) (out *model.SysDictTypeOneOutput, err error) {
	err = dao.SysDictType.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "dictId-" + gconv.String(in.DictId),
		Force:    false,
	}).Where("dict_id", in.DictId).Scan(&out)
	return
}

// 新增字典类型
func (s *sDictType) Create(ctx context.Context, in model.SysDictTypeCreateInput) (err error) {
	dictTypeCount, err := dao.SysDictType.Ctx(ctx).Where("dict_name=? AND dict_type=?", in.DictName, in.DictType).Count()
	if err != nil {
		return err
	}
	if dictTypeCount > 0 {
		return errors.New("字典名称或字典类型已存在！")
	}
	_, err = dao.SysDictType.Ctx(ctx).Insert(in)
	return
}

// 更新字典类型,并删除缓存
func (s *sDictType) Update(ctx context.Context, in model.SysDictTypeUpdateInput) (err error) {
	var dictTypeEntity *model.SysDictTypeOneOutput
	if err = dao.SysDictType.Ctx(ctx).Where("dict_id IN(?)", in.DictId).Scan(&dictTypeEntity); err != nil {
		return
	}
	// 更新字典数据，并清空缓存
	if _, err = dao.SysDictData.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "dictType-" + dictTypeEntity.DictType,
	}).Data("dict_type", in.DictType).Where("dict_type", dictTypeEntity.DictType).Update(); err != nil {
		return
	}

	// 更新字典类型
	_, err = dao.SysDictType.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "dictId-" + gconv.String(in.DictId),
	}).Data(in).Where("dict_id", in.DictId).Update()
	return
}

// 删除字典类型,并删除缓存
func (s *sDictType) Delete(ctx context.Context, in model.SysDictTypeDeleteInput) (err error) {
	var dictTypeList []*model.SysDictTypeOneOutput
	if err = dao.SysDictType.Ctx(ctx).Where("dict_id IN(?)", gstr.Split(in.DictIdStr, ",")).Scan(&dictTypeList); err != nil {
		return
	}
	for _, v := range dictTypeList {
		// 删除字典数据
		if _, err = dao.SysDictData.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "dictType-" + v.DictType,
		}).Delete("dict_type=?", v.DictType); err != nil {
			return
		}
		// 删除字典类型
		if _, err = dao.SysDictType.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "dictId-" + gconv.String(v.DictId),
		}).Delete("dictType_id=?", v); err != nil {
			return
		}
	}
	return
}
