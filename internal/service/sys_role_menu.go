package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
)

type sRoleMenu struct{}

// 角色菜单关联服务
func SysRoleMenu() *sRoleMenu {
	return &sRoleMenu{}
}

// 根据角色id表，获取菜单字段列表输出
// 排除被禁用的菜单
func (s *sRoleMenu) GetFieldList(ctx context.Context, roleId []int) (out model.SysRoleMenuFieldsOutput, err error) {
	// 获取角色对应的菜单id
	menuIds, err := dao.SysRoleMenu.Ctx(ctx).Where("role_id IN(?)", roleId).Array("menu_id")
	if err != nil {
		return
	}
	// 获取菜单信息
	var menuEntitys []*model.SysMenuOneOutput
	if err = dao.SysMenu.Ctx(ctx).Where("status=0 AND menu_id IN(?)", menuIds).Scan(&menuEntitys); err != nil {
		return
	}
	for _, v := range menuEntitys {
		out.MenuId = append(out.MenuId, v.MenuId)
		out.Perms = append(out.Perms, v.Perms)
	}
	return
}

// // 更新用户绑定的角色
// func (s sRoleMenu) UpdateUser(ctx context.Context, in model.SysRoleMenuUpdateUInput) (err error) {
// 	// 删除用户的所有关联数据
// 	if _, err = dao.SysRoleMenu.Ctx(ctx).Delete("role_id", in.UserId); err != nil {
// 		return
// 	}
// 	if len(in.Roleids) == 0 {
// 		return
// 	}
// 	// 添加关联信息
// 	var userRoleWrite []map[string]interface{}
// 	for _, v := range in.Roleids {
// 		userRoleWrite = append(userRoleWrite, g.Map{
// 			"role_id": in.UserId,
// 			"role_id": v,
// 		})
// 	}
// 	_, err = dao.SysRoleMenu.Ctx(ctx).Data(userRoleWrite).Save()
// 	return
// }

// // 更新角色绑定的用户
// func (s sRoleMenu) UpdateRole(ctx context.Context, in model.SysRoleMenuUpdateRInput) (err error) {
// 	// 删除用户的所有关联数据
// 	if _, err = dao.SysRoleMenu.Ctx(ctx).Delete("role_id", in.Roleid); err != nil {
// 		return
// 	}
// 	if len(in.UserIds) == 0 {
// 		return
// 	}
// 	// 添加关联信息
// 	var userRoleWrite []map[string]interface{}
// 	for _, v := range in.UserIds {
// 		userRoleWrite = append(userRoleWrite, g.Map{
// 			"role_id": in.Roleid,
// 			"role_id": v,
// 		})
// 	}
// 	_, err = dao.SysRoleMenu.Ctx(ctx).Data(userRoleWrite).Save()
// 	return
// }

// // 删除关联信息
// func (s *sRoleMenu) Delete(ctx context.Context, in model.SysRoleMenuDeleteInput) (err error) {
// 	_, err = dao.SysRoleMenu.Ctx(ctx).Delete("role_id IN(?) OR role_id IN(?)", gstr.Split(in.UserIdStr, ","), gstr.Split(in.RoleIdStr, ","))
// 	return
// }
