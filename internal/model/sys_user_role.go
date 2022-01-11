package model

// 用户id查询输入
type SysUserRoleListInput struct {
	UserId int // 用户ID
	RoleId int // 角色ID
}

// 用户id查询输入
type SysUserRoleListOutput struct {
	UserIds []int // 用户ID
	RoleIds []int // 角色ID
}
