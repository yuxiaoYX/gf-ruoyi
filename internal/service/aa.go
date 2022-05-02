// ==========================================================================
// 自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-05-02 16:47:40
// ==========================================================================

package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/text/gstr"
)

type sSysAa struct{}

// 参数配置表管理服务
func SysAa() *sSysAa {
	return &sSysAa{}
}

// 获取参数配置表列表
func (s *sSysAa) GetList(ctx context.Context, in model.SysAaListInput) (out model.SysAaListOutput, err error) {
	m := dao.SysAa.Ctx(ctx).OmitEmpty()
	if in.ConfigName != "" {
		m = m.WhereLike("config_name", "%"+in.ConfigName+"%")
	}
	m = m.Where("config_key", in.ConfigKey)
	m = m.Where("config_value", in.ConfigValue)
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
	return
}

// 获取参数配置表详细信息
func (s *sSysAa) GetOne(ctx context.Context, in model.SysAaOneInput) (out *model.SysAaOneOutput, err error) {
	err = dao.SysAa.Ctx(ctx).OmitEmpty().Where(in).Scan(&out)
	return
}

// 新增参数配置表
func (s *sSysAa) Create(ctx context.Context, in model.SysAaCreateInput) (err error) {
	_, err = dao.SysAa.Ctx(ctx).OmitEmpty().Insert(in)
	return
}

// 更新参数配置表
func (s *sSysAa) Update(ctx context.Context, in model.SysAaUpdateInput) (err error) {
	_, err = dao.SysAa.Ctx(ctx).OmitEmpty().Data(in).Where(
		in.ConfigId,
	).Update()
	return
}

// 删除参数配置表
func (s *sSysAa) Delete(ctx context.Context, in model.SysAaDeleteInput) (err error) {
	configIdList := gstr.Split(in.ConfigIdStr, ",")
	for _, v := range configIdList {
		if _, err = dao.SysAa.Ctx(ctx).Delete("config_id=?", v); err != nil {
			return
		}
	}
	return
}
