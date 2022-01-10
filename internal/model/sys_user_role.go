package model

// 用户id查询输入
type SysUserRoleGetInput struct {
	UserId int // 用户ID
	RoleId int // 角色ID
}

// 用户id查询输入
type SysUserRoleGetOutput struct {
	UserId int   // 用户ID
	RoleId []int // 角色ID
}
