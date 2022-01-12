package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"
)

// 角色管理服务
var SysRole = sysRoleService{}

type sysRoleService struct{}

// 根据角色id列表获取角色信息
func (s *sysRoleService) GetIdsList(ctx context.Context, in model.SysRoleGetIdsInput) {

}

// 根据角色id，获取角色权限字符列表
func (s *sysRoleService) RoleKeyList(ctx context.Context, roleIds []int) (roleKey []string, err error) {
	var roleEntity []*entity.SysRole
	err = dao.SysRole.Ctx(ctx).WherePri(roleIds).Where("status=1").Scan(&roleEntity)
	if err != nil {
		return nil, err
	}
	for _, v := range roleEntity {
		roleKey = append(roleKey, v.RoleKey)
	}
	return
}
