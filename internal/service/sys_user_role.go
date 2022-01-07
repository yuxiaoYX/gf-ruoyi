package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/glog"
)

// 用户和角色管理表，管理服务
var SysUserRole = sysUserRoleService{}

type sysUserRoleService struct{}

// 根据用户id查询对应角色id
func (s *sysUserRoleService) GetByUserId(ctx context.Context, in *model.SysUserRoleUserIdInput) (out []*model.SysUserRoleUserIdOutput, err error) {
	err = dao.SysUserRole.Ctx(ctx).OmitEmpty().Where(&in).Scan(&out)
	glog.Debug(ctx, out)
	return
}
