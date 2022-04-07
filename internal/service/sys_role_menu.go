package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRoleMenu struct{}

// 角色菜单关联服务
func SysRoleMenu() *sRoleMenu {
	return &sRoleMenu{}
}

// 根据角色id表，获取菜单字段列表输出
// 排除被禁用的菜单
func (s *sRoleMenu) GetFieldList(ctx context.Context, roleIds []int) (out model.SysRoleMenuFieldsOutput, err error) {
	// 获取角色对应的菜单id
	// menuIds, err := dao.SysRoleMenu.Ctx(ctx).Where("role_id IN(?)", roleId).Array("menu_id")
	menuIds, err := s.GetMenuIds(ctx, roleIds)
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
		if v.Perms != "" {
			out.Perms = append(out.Perms, v.Perms)
		}
	}
	return
}

// 根据角色id列表，获取对应菜单id列表
func (s *sRoleMenu) GetMenuIds(ctx context.Context, roleIds []int) (out []int, err error) {
	// 获取角色对应的菜单id
	menuEntity, err := dao.SysRoleMenu.Ctx(ctx).Where("role_id IN(?)", roleIds).Array("menu_id")
	if err != nil {
		return
	}
	out = gconv.Ints(menuEntity)
	return
}

// 获取角色路由表
func (s *sRoleMenu) GetTreeRoute(ctx context.Context, roleIds []int) (out []map[string]interface{}, err error) {
	//是否管管理员
	isAdmin := false
	for _, i := range roleIds {
		if i == 1 {
			isAdmin = true
			break
		}
	}
	// 角色id为1的管理员，获取所有路由
	var menuEntitys []*model.SysMenuOneOutput
	if isAdmin {
		if err = dao.SysMenu.Ctx(ctx).Where("status=0").Scan(&menuEntitys); err != nil {
			return
		}
	} else {
		var menuIds []int
		menuIds, err = s.GetMenuIds(ctx, roleIds)
		if err != nil {
			return
		}
		// 获取菜单信息
		if err = dao.SysMenu.Ctx(ctx).Where("status=0 AND menu_id IN(?)", menuIds).Scan(&menuEntitys); err != nil {
			return
		}
	}
	if out, err = s.formTreeRouter(ctx, 0, menuEntitys); err != nil {
		return
	}
	return
}

// 构造树形菜单列表
func (s *sRoleMenu) formTreeRouter(ctx context.Context, parentId int, entities []*model.SysMenuOneOutput) (treeRoute []map[string]interface{}, err error) {
	for _, entity := range entities {
		if entity.ParentId == parentId && entity.MenuType != "F" {
			subTree, err := s.formTreeRouter(ctx, entity.MenuId, entities)
			if err != nil {
				return nil, err
			}
			children := s.setMapRouter(ctx, entity)
			children["children"] = subTree
			treeRoute = append(treeRoute, children)
		}
	}
	return treeRoute, err
}

//组合返回menu前端数据
func (s *sRoleMenu) setMapRouter(ctx context.Context, entity *model.SysMenuOneOutput) (menu map[string]interface{}) {

	menu = make(map[string]interface{})
	menu["name"] = gstr.UcFirst(entity.Path)

	metaMap := g.Map{
		"title": entity.MenuName,
		"icon":  entity.Icon,
	}
	if entity.IsCache == "1" {
		metaMap["noCache"] = true
	} else {
		metaMap["noCache"] = false
	}
	menu["meta"] = metaMap

	menu["hidden"] = gconv.Bool(entity.Visible)

	if entity.IsFrame == "1" && entity.MenuType == "M" && entity.ParentId == 0 {
		// 主菜单
		menu["alwaysShow"] = true
		menu["component"] = "Layout"
		menu["path"] = "/" + entity.Path
	} else if entity.IsFrame == "1" && entity.MenuType == "M" {
		// 次级菜单
		menu["alwaysShow"] = true
		menu["component"] = "ParentView"
		menu["path"] = "/" + entity.Path
	} else if entity.IsFrame == "0" {
		// 外链
		menu["alwaysShow"] = false
		menu["component"] = "Layout"
		menu["path"] = entity.Path
	} else {
		// 页面
		menu["component"] = entity.Component
		menu["path"] = entity.Path
	}
	return menu
}
