package service

import (
	"context"
	"errors"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sDictData struct{}

// 字典数据管理服务
func SysDictData() *sDictData {
	return &sDictData{}
}

// 根据字典类型查询字典数据,缓存10小时
func (s *sDictData) GetType(ctx context.Context, dictType string) (out []*model.SysDictDataOneOutput, err error) {
	err = dao.SysDictData.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "dictType-" + dictType,
		Force:    false,
	}).Where("dict_type", dictType).OrderAsc("dict_sort").Scan(&out)
	return
}

// 获取字典数据列表
func (s *sDictData) GetList(ctx context.Context, in model.SysDictDataListInput) (out model.SysDictDataListOutput, err error) {
	m := dao.SysDictData.Ctx(ctx).OmitEmpty().Where(g.Map{
		"dict_label": in.DictLabel,
		"dict_type":  in.DictType,
		"status":     in.Status,
	})
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).OrderAsc("dict_sort").Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
	return
}

// 获取字典数据详细信息
func (s *sDictData) GetOne(ctx context.Context, in model.SysDictDataOneInput) (out *model.SysDictDataOneOutput, err error) {
	err = dao.SysDictData.Ctx(ctx).Where("dict_code", in.DictCode).Scan(&out)
	return
}

// 新增字典数据
func (s *sDictData) Create(ctx context.Context, in model.SysDictDataCreateInput) (err error) {
	dictDataCount, err := dao.SysDictData.Ctx(ctx).Where("dict_label=? AND dict_value=?", in.DictLabel, in.DictValue).Count()
	if err != nil {
		return err
	}
	if dictDataCount > 0 {
		return errors.New("字典标签或字典键值已存在！")
	}
	_, err = dao.SysDictData.Ctx(ctx).Insert(in)
	return
}

// 更新字典数据,并删除缓存
func (s *sDictData) Update(ctx context.Context, in model.SysDictDataUpdateInput) (err error) {
	_, err = dao.SysDictData.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "dictType-" + in.DictType,
	}).Data(in).Where("dict_code", in.DictCode).Update()
	return
}

// 删除字典数据,并删除缓存
func (s *sDictData) Delete(ctx context.Context, in model.SysDictDataDeleteInput) (err error) {
	var dictDataList []*model.SysDictDataOneOutput
	if err = dao.SysDictData.Ctx(ctx).Where("dict_code IN(?)", in.DictCodeList).Scan(&dictDataList); err != nil {
		return
	}
	for _, v := range dictDataList {
		if _, err = dao.SysDictData.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "dictType-" + v.DictType,
		}).Delete("dict_code=?", v.DictCode); err != nil {
			return
		}
	}
	return
}
