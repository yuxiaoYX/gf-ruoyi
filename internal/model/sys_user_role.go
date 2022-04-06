package model

// 根据用户id，获取角色id列表和角色名称列表输出
type SysUserRoleFieldsOutput struct {
	RoleIds   []uint   // 角色ID
	RoleNames []string // 角色名称

}

// 更新用户绑定的角色输入
type SysUserRoleUpdateUInput struct {
	UserId  int   // 用户id
	Roleids []int // 角色id列表
}

// 更新角色绑定的角色输入
type SysUserRoleUpdateRInput struct {
	UserIds []int // 用户id列表
	RoleId  int   // 角色id
}

// 删除关联信息输入
type SysUserRoleDeleteInput struct {
	UserIdStr string // 需要删除的数据主键，例：1,2,3
	RoleIdStr string // 需要删除的数据主键，例：1,2,3
}
