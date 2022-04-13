package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 参数管理
var SysConfig = cConfig{}

type cConfig struct{}

// 获取参数列表
func (c *cConfig) GetList(ctx context.Context, req *v1.SysConfigListReq) (res *v1.SysConfigListRes, err error) {
	in := &model.SysConfigListInput{}
	gconv.Scan(req, &in)

	roleRes, err := service.SysConfig().GetList(ctx, *in)
	gconv.Scan(roleRes, &res)
	return
}

// 获取单个参数信息
func (c *cConfig) GetOne(ctx context.Context, req *v1.SysConfigOneReq) (res *v1.SysConfigOneRes, err error) {
	roleRes, err := service.SysConfig().GetOne(ctx, model.SysConfigOneInput{
		ConfigId: req.ConfigId,
	})
	gconv.Scan(roleRes, &res)
	return
}

// 新建参数
func (c *cConfig) Create(ctx context.Context, req *v1.SysConfigCreateReq) (res *v1.SysConfigCreateRes, err error) {
	in := &model.SysConfigCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysConfig().Create(ctx, *in)
	return
}

// 更新参数
func (c *cConfig) Update(ctx context.Context, req *v1.SysConfigUpdateReq) (res *v1.SysConfigUpdateRes, err error) {
	in := &model.SysConfigUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysConfig().Update(ctx, *in)
	return
}

// 删除参数
func (c *cConfig) Delete(ctx context.Context, req *v1.SysConfigDeleteReq) (res *v1.SysConfigDeleteRes, err error) {
	in := &model.SysConfigDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysConfig().Delete(ctx, *in)
	return
}

// 查询参数已授权或未授权用户列表
func (c *cConfig) ConfigKey(ctx context.Context, req *v1.SysConfigKeyReq) (res *v1.SysConfigKeyRes, err error) {
	res = &v1.SysConfigKeyRes{}
	userConfigRes, err := service.SysConfig().ConfigKey(ctx, model.SysConfigKeyInput{ConfigKey: req.ConfigKey})
	res.ConfigValue = userConfigRes.ConfigValue
	return
}
