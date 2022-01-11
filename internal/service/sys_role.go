package service

import (
	"context"
	"gf-RuoYi/internal/model"
)

// 角色管理服务
var SysRole = sysRoleService{}

type sysRoleService struct{}

// 根据角色id列表获取角色信息
func (s *sysRoleService) GetIdsList(ctx context.Context, in model.SysRoleGetIdsInput) {

}
