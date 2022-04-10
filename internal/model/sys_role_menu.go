package model

// 根据角色id表，获取菜单字段列表输出
type SysRoleMenuFieldsOutput struct {
	MenuId []int    // 菜单ID
	Perms  []string // 权限标识
}

// 更新角色绑定的菜单输入
type SysRoleMenuUpdateRInput struct {
	RoleId  int   // 角色id
	MenuIds []int // 菜单id列表
}
