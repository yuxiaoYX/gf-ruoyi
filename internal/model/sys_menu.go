package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 获取菜单列表输入
type SysMenuListInput struct {
	MenuName string // 菜单名称
	Status   string // 菜单状态（0正常 1停用）
}

// 获取菜单列表输出
type SysMenuListOutput struct {
	Rows  []*SysMenuOneOutput // 列表
	Total int                 // 数据总数
}

// 获取单个菜单信息输入
type SysMenuOneInput struct {
	MenuId int // 菜单ID
}

// 获取单个菜单信息输出
type SysMenuOneOutput struct {
	MenuId    int         // 菜单ID
	MenuName  string      // 菜单名称
	ParentId  int         // 父菜单ID
	Sort      int         // 排序标记
	Path      string      // 路由地址
	Component string      // 组件路径
	IsFrame   string      // 是否为外链（0是 1否）
	IsCache   string      // 是否缓存（0缓存 1不缓存）
	Query     string      // 路由参数
	MenuType  string      // 菜单类型（M目录 C菜单 F按钮）
	Visible   string      // 显示状态（0显示 1隐藏）
	Status    string      // 菜单状态（0正常 1停用）
	Perms     string      // 权限标识
	Icon      string      // 菜单图标
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}

// 新增菜单输入
type SysMenuCreateInput struct {
	MenuName  string // 菜单名称
	ParentId  int    // 父菜单ID
	Sort      int    // 排序标记
	Path      string // 路由地址
	Component string // 组件路径
	IsFrame   string // 是否为外链（0是 1否）
	IsCache   string // 是否缓存（0缓存 1不缓存）
	Query     string // 路由参数
	MenuType  string // 菜单类型（M目录 C菜单 F按钮）
	Visible   string // 显示状态（0显示 1隐藏）
	Status    string // 菜单状态（0正常 1停用）
	Perms     string // 权限标识
	Icon      string // 菜单图标
}

// 更新菜单信息输入
type SysMenuUpdateInput struct {
	MenuId    int    // 菜单ID
	MenuName  string // 菜单名称
	ParentId  int    // 父菜单ID
	Sort      int    // 排序标记
	Path      string // 路由地址
	Component string // 组件路径
	IsFrame   string // 是否为外链（0是 1否）
	IsCache   string // 是否缓存（0缓存 1不缓存）
	Query     string // 路由参数
	MenuType  string // 菜单类型（M目录 C菜单 F按钮）
	Visible   string // 显示状态（0显示 1隐藏）
	Status    string // 菜单状态（0正常 1停用）
	Perms     string // 权限标识
	Icon      string // 菜单图标
}

// 删除菜单输入
type SysMenuDeleteInput struct {
	MenuIdStr string // 需要删除的数据主键，例：1,2,3
}
