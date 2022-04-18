package service

import (
	"context"
	"fmt"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type sUserRole struct{}

// 用户和角色关联服务
func SysUserRole() *sUserRole {
	return &sUserRole{}
}

// // 根据用户id，获取角色id列表和角色权限字符列表
// func (s *sUserRole) GetFieldList(ctx context.Context, userId uint) (out model.SysUserRoleFieldsOutput, err error) {
// 	roleEntitys, err := s.GetRoles(ctx, userId)
// 	if err != nil {
// 		return
// 	}
// 	for _, v := range roleEntitys {
// 		out.RoleIds = append(out.RoleIds, v.RoleId)
// 		out.RoleNames = append(out.RoleNames, v.RoleName)
// 	}
// 	return
// }

// 获取用户关联角色信息
// 排除被禁用的角色
func (s sUserRole) GetRoles(ctx context.Context, userId uint) (out model.SysUserRolesOutput, err error) {
	roleIds, err := dao.SysUserRole.Ctx(ctx).Where("user_id", userId).Array("role_id")
	if err != nil {
		return
	}
	err = dao.SysRole.Ctx(ctx).Where("status=0 AND role_id IN(?)", roleIds).Scan(&out.Roles)
	if err != nil {
		return
	}
	for _, v := range out.Roles {
		out.RoleIds = append(out.RoleIds, v.RoleId)
		out.RoleNames = append(out.RoleNames, v.RoleName)
	}
	return
}

// 查询角色已授权或未授权用户列表
func (s *sUserRole) GetAllocatedList(ctx context.Context, in model.SysUserRoleAllocatedListInput) (out model.SysUserRoleAllocatedListOutput, err error) {
	// 获取角色已授权或未授权用户id列表
	var userIds []*gvar.Var
	userIds, err = dao.SysUserRole.Ctx(ctx).Where("role_id", in.RoleId).Array("user_id")
	if err != nil || len(userIds) == 0 {
		return
	}
	fmt.Println(userIds)
	// 获取用户信息列表
	m := dao.SysUser.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": in.UserName,
		"mobile":    in.Mobile,
	})
	if in.IsAllocated {
		m = m.Where("user_id IN(?)", userIds)
	} else {
		m = m.WhereNotIn("user_id", userIds)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
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

// 角色分配用户
func (s sUserRole) RoleSelectUser(ctx context.Context, in model.SysRoleSelectUserInput) (err error) {
	userIdList := gstr.Split(in.UserIds, ",")
	if len(userIdList) == 0 {
		return
	}
	// 添加关联信息
	var userRoleWrite []map[string]interface{}
	for _, v := range userIdList {
		userRoleWrite = append(userRoleWrite, g.Map{
			"role_id": in.RoleId,
			"user_id": v,
		})
	}
	_, err = dao.SysUserRole.Ctx(ctx).Data(userRoleWrite).Save()
	return
}

// 角色取消用户绑定
func (s sUserRole) RoleCancelUser(ctx context.Context, in model.SysRoleCancelUserInput) (err error) {
	userIdList := gstr.Split(in.UserIds, ",")
	_, err = dao.SysUserRole.Ctx(ctx).Where("role_id=? AND user_id IN(?)", in.RoleId, userIdList).Delete()
	return
}

// 删除关联信息
func (s *sUserRole) Delete(ctx context.Context, in model.SysUserRoleDeleteInput) (err error) {
	m := dao.SysUserRole.Ctx(ctx)
	if in.UserIdStr != nil {
		m = m.WhereOr("user_id IN(?)", in.UserIdStr)
	}
	if in.RoleIdStr != nil {
		m = m.WhereOr("role_id IN(?)", in.RoleIdStr)
	}
	_, err = m.Delete()
	return
}
