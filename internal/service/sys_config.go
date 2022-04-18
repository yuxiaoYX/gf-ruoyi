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

type sConfig struct{}

// 参数管理服务
func SysConfig() *sConfig {
	return &sConfig{}
}

// 获取参数列表
func (s *sConfig) GetList(ctx context.Context, in model.SysConfigListInput) (out model.SysConfigListOutput, err error) {
	m := dao.SysConfig.Ctx(ctx).OmitEmpty().Where(g.Map{
		"config_name":  in.ConfigName,
		"config_key":   in.ConfigKey,
		"config_value": in.ConfigValue,
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

// 获取参数详细信息,缓存10小时
func (s *sConfig) GetOne(ctx context.Context, in model.SysConfigOneInput) (out *model.SysConfigOneOutput, err error) {
	err = dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "configId-" + gconv.String(in.ConfigId),
		Force:    false,
	}).Where("config_id", in.ConfigId).Scan(&out)
	return
}

// 新增参数
func (s *sConfig) Create(ctx context.Context, in model.SysConfigCreateInput) (err error) {
	roleCount, err := dao.SysConfig.Ctx(ctx).Where("config_name=? AND config_key=?", in.ConfigName, in.ConfigKey).Count()
	if err != nil {
		return err
	}
	if roleCount > 0 {
		return errors.New("参数名称或参数键名已存在！")
	}
	_, err = dao.SysConfig.Ctx(ctx).Insert(in)
	return
}

// 更新参数,并删除缓存
func (s *sConfig) Update(ctx context.Context, in model.SysConfigUpdateInput) (err error) {
	_, err = dao.SysConfig.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "configId-" + gconv.String(in.ConfigId),
	}).Data(in).Where("config_id", in.ConfigId).Update()
	return
}

// 删除参数,并删除缓存
func (s *sConfig) Delete(ctx context.Context, in model.SysConfigDeleteInput) (err error) {
	configIdList := gstr.Split(in.ConfigIdStr, ",")
	for _, v := range configIdList {
		if _, err = dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "configId-" + v,
		}).Delete("config_id=?", v); err != nil {
			return
		}
	}
	return
}

// 根据参数键名查询参数值
func (s *sConfig) ConfigKey(ctx context.Context, in model.SysConfigKeyInput) (out *model.SysConfigKeyOutput, err error) {
	err = dao.SysConfig.Ctx(ctx).Where("config_key", in.ConfigKey).Scan(&out)
	return
}
