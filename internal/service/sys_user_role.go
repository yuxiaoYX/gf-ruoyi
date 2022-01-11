package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/glog"
)

// 用户和角色管理表，管理服务
var SysUserRole = sysUserRoleService{}

type sysUserRoleService struct{}

// 根据用户id查询对应角色id
func (s *sysUserRoleService) GetByIdList(ctx context.Context, in model.SysUserRoleListInput) (out *model.SysUserRoleListOutput, err error) {
	var userRoleEntity []*entity.SysUserRole
	err = dao.SysUserRole.Ctx(ctx).OmitEmpty().Where(&in).Scan(&userRoleEntity)
	for _, v := range userRoleEntity {
		out.UserIds = append(out.UserIds, v.UserId)
		out.RoleIds = append(out.RoleIds, v.RoleId)
	}
	glog.Debug(ctx, out)
	return
}
