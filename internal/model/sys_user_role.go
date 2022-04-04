package model

// 更新用户绑定的角色输入
type SysUserRoleUpdateUInput struct {
	UserId  int   `v:"required|min-length:1#角色id不能为空|角色id不能为空！"`
	Roleids []int `v:"required|min-length:1#菜单id不能为空|菜单id不能为空！"`
}

// 更新角色绑定的角色输入
type SysUserRoleUpdateRInput struct {
	UserIds []int `v:"required|min-length:1#角色id不能为空|角色id不能为空！"`
	Roleid  int   `v:"required|min-length:1#菜单id不能为空|菜单id不能为空！"`
}

// 删除关联信息输入
type SysUserRoleDeleteInput struct {
	UserIdStr string // 需要删除的数据主键，例：1,2,3
	RoleIdStr string // 需要删除的数据主键，例：1,2,3
}
