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
func (s *sMenu) GetList(ctx context.Context, in model.SysMenuListInput) (out model.SysMenuListOutput, err error) {
	m := dao.SysMenu.Ctx(ctx).OmitEmpty().Where(g.Map{
		"menu_name like ?": "%" + in.MenuName + "%",
		"status":           in.Status,
	})
	if m.Scan(&out.Rows); err != nil {
		return
	}
	out.Total = len(out.Rows)
	return
}

// 获取菜单详细信息,缓存10小时
func (s *sMenu) GetOne(ctx context.Context, in model.SysMenuOneInput) (out *model.SysMenuOneOutput, err error) {
	err = dao.SysMenu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "menuid-" + gconv.String(in.MenuId),
		Force:    false,
	}).Where("menu_id", in.MenuId).Scan(&out)
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
func (s *sMenu) Update(ctx context.Context, in model.SysMenuUpdateInput) (err error) {
	_, err = dao.SysMenu.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "menuid-" + gconv.String(in.MenuId),
	}).Data(in).Where("menu_id", in.MenuId).Update()
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
