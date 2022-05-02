// ==========================================================================
// 自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-05-02 16:47:40
// ==========================================================================

package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 参数配置表管理
var SysAa = cSysAa{}

type cSysAa struct{}

// 分页获取参数配置表列表
func (c *cSysAa) GetList(ctx context.Context, req *v1.SysAaListReq) (res *v1.SysAaListRes, err error) {
	sysAaRes, err := service.SysAa().GetList(ctx, model.SysAaListInput{
		ConfigName:  req.ConfigName,
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		BeginTime:   req.BeginTime,
		EndTime:     req.EndTime,
		PageNum:     req.PageNum,
		PageSize:    req.PageSize,
	})
	gconv.Scan(sysAaRes, &res)
	return
}

// 获取单个参数配置表信息
func (c *cSysAa) GetOne(ctx context.Context, req *v1.SysAaOneReq) (res *v1.SysAaOneRes, err error) {
	sysAaRes, err := service.SysAa().GetOne(ctx, model.SysAaOneInput{
		ConfigId: req.ConfigId,
	})
	gconv.Scan(sysAaRes, &res)
	return
}

// 新建参数配置表
func (c *cSysAa) Create(ctx context.Context, req *v1.SysAaCreateReq) (res *v1.SysAaCreateRes, err error) {
	err = service.SysAa().Create(ctx, model.SysAaCreateInput{
		ConfigName:  req.ConfigName,
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		ConfigType:  req.ConfigType,
		Remark:      req.Remark,
	})
	return
}

// 更新参数配置表
func (c *cSysAa) Update(ctx context.Context, req *v1.SysAaUpdateReq) (res *v1.SysAaUpdateRes, err error) {
	err = service.SysAa().Update(ctx, model.SysAaUpdateInput{
		ConfigId:    req.ConfigId,
		ConfigName:  req.ConfigName,
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		ConfigType:  req.ConfigType,
		Remark:      req.Remark,
	})
	return
}

// 删除参数配置表
func (c *cSysAa) Delete(ctx context.Context, req *v1.SysAaDeleteReq) (res *v1.SysAaDeleteRes, err error) {
	err = service.SysAa().Delete(ctx, model.SysAaDeleteInput{
		ConfigIdStr: req.ConfigIdStr,
	})
	return
}
