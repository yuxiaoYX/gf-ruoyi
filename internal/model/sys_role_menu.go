package model

// 根据角色id表，获取菜单字段列表输出
type SysRoleMenuFieldsOutput struct {
	MenuId []int    // 菜单ID
	Perms  []string // 权限标识

}

// 更新用户绑定的菜单输入
type SysRoleMenuUpdateUInput struct {
	RoleId  int   // 用户id
	Menuids []int // 菜单id列表
}

// 更新菜单绑定的菜单输入
type SysRoleMenuUpdateRInput struct {
	RoleIds []int // 用户id列表
	MenuId  int   // 菜单id
}

// 删除关联信息输入
type SysRoleMenuDeleteInput struct {
	RoleIdStr string // 需要删除的数据主键，例：1,2,3
	MenuIdStr string // 需要删除的数据主键，例：1,2,3
}
