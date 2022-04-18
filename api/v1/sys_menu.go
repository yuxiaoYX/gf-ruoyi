package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取菜单列表请求
type SysMenuListReq struct {
	g.Meta   `path:"/menu/getList" method:"post" summary:"获取菜单列表" tags:"菜单"`
	MenuName string `v:"max-length:60#菜单名称长度最多为:{max}位" dc:"菜单名称"` // 菜单名称
	Status   string `dc:"菜单状态（0正常 1停用）" `                          // 菜单状态（0正常 1停用）
}

// 获取菜单列表响应
type SysMenuListRes []map[string]interface{}

// 单个菜单信息请求
type SysMenuOneReq struct {
	g.Meta `path:"/menu/getOne" method:"post" summary:"获取单个菜单信息" tags:"菜单"`
	MenuId int `v:"required|length:1,10#菜单id不能为空！|菜单名长度为:{min}到{max}位" dc:"菜单id"` // 菜单ID
}

// 单个菜单信息响应
type SysMenuOneRes struct {
	*entity.SysMenu
}

// 新增菜单请求
type SysMenuCreateReq struct {
	g.Meta    `path:"/menu/create" method:"post" summary:"新增菜单" tags:"菜单"`
	MenuName  string `v:"required|length:1,60#菜单名称不能为空！|菜单名称长度为:{min}到{max}位" dc:"菜单名称"` // 菜单名称
	ParentId  int    `json:"parentId" dc:"父菜单ID" `                                       // 父菜单ID
	Sort      int    `v:"required|length:1,3#排序标记不能为空！|排序标记长度为:{min}到{max}位" dc:"排序标记"`  // 排序标记
	Path      string `dc:"路由地址"`                                                         // 路由地址
	Component string `dc:"组件路径" `                                                        // 组件路径
	IsFrame   string `d:"1" dc:"是否为外链（0是 1否）" `                                          // 是否为外链（0是 1否）
	IsCache   string `d:"1" dc:"是否缓存（0缓存 1不缓存）   " `                                     // 是否缓存（0缓存 1不缓存）
	Query     string `dc:"路由参数" `                                                        // 路由参数
	MenuType  string `d:"M" dc:"菜单类型（M目录 C菜单 F按钮）" `                                     // 菜单类型（M目录 C菜单 F按钮）
	Visible   string `d:"0" dc:"显示状态（0显示 1隐藏）" `                                         // 显示状态（0显示 1隐藏）
	Status    string `d:"0" dc:"菜单状态（0正常 1停用）" `                                         // 菜单状态（0正常 1停用）
	Perms     string `dc:"权限标识" `                                                        // 权限标识
	IsAuth    string `d:"0" dc:"是否验证（0是 1否）" `                                           // 是否验证（0是 1否）
	IsLog     string `d:"0" dc:"是否记录操作日志（0是 1否）" `                                       // 是否记录操作日志（0是 1否）
	Icon      string `dc:"菜单图标" `                                                        // 菜单图标
}

// 新增菜单响应
type SysMenuCreateRes struct{}

// 更新菜单信息请求
type SysMenuUpdateReq struct {
	g.Meta    `path:"/menu/update" method:"post" summary:"修改菜单" tags:"菜单"`
	MenuId    int    `v:"required|length:1,10#菜单id不能为空！|菜单名长度为:{min}到{max}位" dc:"菜单id"`  // 菜单ID
	MenuName  string `v:"required|length:1,60#菜单名称不能为空！|菜单名称长度为:{min}到{max}位" dc:"菜单名称"` // 菜单名称
	ParentId  int    `json:"parentId" dc:"父菜单ID" `                                       // 父菜单ID
	Sort      int    `v:"required|length:1,3#排序标记不能为空！|排序标记长度为:{min}到{max}位" dc:"排序标记"`  // 排序标记
	Path      string `dc:"路由地址"`                                                         // 路由地址
	Component string `dc:"组件路径" `                                                        // 组件路径
	IsFrame   string `d:"1" dc:"是否为外链（0是 1否）" `                                          // 是否为外链（0是 1否）
	IsCache   string `d:"1" dc:"是否缓存（0缓存 1不缓存） " `                                       // 是否缓存（0缓存 1不缓存）
	Query     string `dc:"路由参数" `                                                        // 路由参数
	MenuType  string `d:"M" dc:"菜单类型（M目录 C菜单 F按钮）"`                                      // 菜单类型（M目录 C菜单 F按钮）
	Visible   string `d:"0" dc:"显示状态（0显示 1隐藏）" `                                         // 显示状态（0显示 1隐藏）
	Status    string `d:"0" dc:"菜单状态（0正常 1停用）" `                                         // 菜单状态（0正常 1停用）
	Perms     string `dc:"权限标识" `                                                        // 权限标识
	IsAuth    string `d:"0" dc:"是否验证（0是 1否）" `                                           // 是否验证（0是 1否）
	IsLog     string `d:"0" dc:"是否记录操作日志（0是 1否）" `                                       // 是否记录操作日志（0是 1否）
	Icon      string `dc:"菜单图标" `                                                        // 菜单图标
}

// 更新菜单信息响应
type SysMenuUpdateRes struct{}

// 删除菜单请求
type SysMenuDeleteReq struct {
	g.Meta    `path:"/menu/delete" method:"post" summary:"删除菜单" tags:"菜单"`
	MenuIdStr string `v:"required#菜单id不能为空！" dc:"菜单id"` // 需要删除的数据主键，例：1,2,3
}

// 删除菜单响应
type SysMenuDeleteRes struct{}

// 查询菜单树下拉结构输入
type SysMenuTreeselectReq struct {
	g.Meta `path:"/menu/treeselect" method:"post" summary:"菜单下拉树" tags:"菜单"`
	RoleId int `v:"length:1,10#角色id长度为:{min}到{max}位" dc:"角色id"` // 角色ID
}

// 查询菜单树下拉结构响应
type SysMenuTreeselectRes struct {
	CheckedKeys []int                    `json:"checkedKeys"`
	Menus       []map[string]interface{} `json:"menus"`
}
