package service

import (
	"context"
	"errors"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

// 菜单管理服务
func SysMenu() *sMenu {
	return &sMenu{}
}

// 获取菜单列表
// TODO OmitEmpty支持特殊字符空值过滤，相关函数可以改写了
func (s *sMenu) GetList(ctx context.Context, in model.SysMenuListInput) (out []*model.SysMenuOneOutput, err error) {
	err = dao.SysMenu.Ctx(ctx).OmitEmpty().Where(g.Map{
		"menu_name like ?": "%" + in.MenuName + "%",
		"status":           in.Status,
	}).Scan(&out)
	return
}

// 获取菜单详细信息,缓存10小时
func (s *sMenu) GetOne(ctx context.Context, in model.SysMenuOneInput) (out *model.SysMenuOneOutput, err error) {
	err = dao.SysMenu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "menuid-" + gconv.String(in.MenuId),
		Force:    false,
	}).Where("menu_id", in.MenuId).Scan(&out)
	s.Treeselect(ctx)
	return
}

// 新增菜单
func (s *sMenu) Create(ctx context.Context, in model.SysMenuCreateInput) (err error) {
	menuCount, err := dao.SysMenu.Ctx(ctx).Where("path=?", in.Path).Count()
	if err != nil {
		return err
	}
	if menuCount > 0 {
		return errors.New("路由地址已存在！")
	}
	_, err = dao.SysMenu.Ctx(ctx).Insert(in)
	return
}

// 更新菜单,并删除缓存
// TODO 有BUG删除菜单时，子菜单的缓存并不会删除
func (s *sMenu) Update(ctx context.Context, in model.SysMenuUpdateInput) (err error) {
	_, err = dao.SysMenu.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "menuid-" + gconv.String(in.MenuId),
	}).Data(in).Where("menu_id=? or parent_id=?", in.MenuId, in.MenuId).Update()
	return
}

// 删除菜单,并删除缓存
func (s *sMenu) Delete(ctx context.Context, in model.SysMenuDeleteInput) (err error) {
	menuIdList := gstr.Split(in.MenuIdStr, ",")
	for _, v := range menuIdList {
		if _, err = dao.SysMenu.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "menuid-" + v,
		}).Delete("menu_id=?", v); err != nil {
			return
		}
	}
	return
}

// 查询菜单下拉树结构
func (s *sMenu) Treeselect(ctx context.Context) (err error) {
	menuEntitys, err := s.GetList(ctx, model.SysMenuListInput{Status: "0"})
	aa, err := s.formTreeRouter(ctx, menuEntitys)
	g.Log().Info(ctx, aa)
	return
}

// 构造树形菜单列表
func (s *sMenu) formTreeRouter(ctx context.Context, entities []*model.SysMenuOneOutput) (treeRoute []map[string]interface{}, err error) {
	for _, entity := range entities {
		if entity.ParentId == 0 && entity.MenuType != "F" {
			subTree, err := s.formTreeRouter(ctx, entities)
			if err != nil {
				return nil, err
			}
			children := make(map[string]interface{})
			children["children"] = subTree
			treeRoute = append(treeRoute, children)
		}

	}
	return treeRoute, err
}
