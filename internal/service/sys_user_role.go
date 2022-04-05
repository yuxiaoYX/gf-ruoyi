package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type sUserRole struct{}

// 用户和角色关联服务
func SysUserRole() *sUserRole {
	return &sUserRole{}
}

// 根据用户id，获取角色id列表和角色权限字符列表
func (s *sUserRole) GetFieldList(ctx context.Context, userId uint) (out model.SysUserRoleFieldsOutput, err error) {
	roleEntitys, err := s.GetRoles(ctx, userId)
	if err != nil {
		return
	}
	for _, v := range roleEntitys {
		out.RoleId = append(out.RoleId, v.RoleId)
		out.RoleName = append(out.RoleName, v.RoleName)
	}
	return
}

// 获取用户关联角色信息
// 排除被禁用的角色
func (s sUserRole) GetRoles(ctx context.Context, userId uint) (out []*model.SysRoleOneOutput, err error) {
	roleIds, err := dao.SysUserRole.Ctx(ctx).Where("user_id", userId).Array("role_id")
	if err != nil {
		return
	}
	err = dao.SysRole.Ctx(ctx).Where("status=0 AND role_id IN(?)", roleIds).Scan(&out)
	return
}

// 获取角色关联用户信息
func (s sUserRole) GetUsers(ctx context.Context, roleId uint) (out []*model.SysUserOneOutput, err error) {
	userIds, err := dao.SysUserRole.Ctx(ctx).Where("role_id", roleId).Array("user_id")
	if err != nil {
		return
	}
	err = dao.SysUser.Ctx(ctx).Where("user_id IN(?)", userIds).Scan(&out)
	return
}

// 更新用户绑定的角色
func (s sUserRole) UpdateUser(ctx context.Context, in model.SysUserRoleUpdateUInput) (err error) {
	// 删除用户的所有关联数据
	if _, err = dao.SysUserRole.Ctx(ctx).Delete("user_id", in.UserId); err != nil {
		return
	}
	if len(in.Roleids) == 0 {
		return
	}
	// 添加关联信息
	var userRoleWrite []map[string]interface{}
	for _, v := range in.Roleids {
		userRoleWrite = append(userRoleWrite, g.Map{
			"user_id": in.UserId,
			"role_id": v,
		})
	}
	_, err = dao.SysUserRole.Ctx(ctx).Data(userRoleWrite).Save()
	return
}

// 更新角色绑定的用户
func (s sUserRole) UpdateRole(ctx context.Context, in model.SysUserRoleUpdateRInput) (err error) {
	// 删除用户的所有关联数据
	if _, err = dao.SysUserRole.Ctx(ctx).Delete("role_id", in.RoleId); err != nil {
		return
	}
	if len(in.UserIds) == 0 {
		return
	}
	// 添加关联信息
	var userRoleWrite []map[string]interface{}
	for _, v := range in.UserIds {
		userRoleWrite = append(userRoleWrite, g.Map{
			"role_id": in.RoleId,
			"user_id": v,
		})
	}
	_, err = dao.SysUserRole.Ctx(ctx).Data(userRoleWrite).Save()
	return
}

// 删除关联信息
func (s *sUserRole) Delete(ctx context.Context, in model.SysUserRoleDeleteInput) (err error) {
	_, err = dao.SysUserRole.Ctx(ctx).Delete("user_id IN(?) OR role_id IN(?)", gstr.Split(in.UserIdStr, ","), gstr.Split(in.RoleIdStr, ","))
	return
}
