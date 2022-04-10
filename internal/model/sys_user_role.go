package model

// 根据用户id，获取角色id列表和角色名称列表输出
type SysUserRoleFieldsOutput struct {
	RoleIds   []uint   // 角色ID
	RoleNames []string // 角色名称

}

// 查询角色已授权或未授权用户列表输入
type SysUserRoleAllocatedListInput struct {
	UserName    string // 用户名
	Mobile      string // 手机号码
	RoleId      uint   // 角色ID
	IsAllocated bool   // 是否获取授权用户;false:未授权,true:已授权
	PageNum     int    // 分页码
	PageSize    int    // 分页数量
}

// 查询角色已授权用户列表输出
type SysUserRoleAllocatedListOutput struct {
	Rows  []*SysUserOneOutput // 列表
	Total int                 // 数据总数
}

// 更新用户绑定的角色输入
type SysUserRoleUpdateUInput struct {
	UserId  int   // 用户id
	Roleids []int // 角色id列表
}

// 角色分配用户输入
type SysRoleSelectUserInput struct {
	UserIds string // 用户id，例：1,2,3
	RoleId  string // 角色id
}

// 角色取消用户输入
type SysRoleCancelUserInput struct {
	UserIds string // 用户id，例：1,2,3
	RoleId  string // 角色id
}

// 删除关联信息输入
type SysUserRoleDeleteInput struct {
	UserIdStr []int // 需要删除的数据主键，例：[1,2,3]
	RoleIdStr []int // 需要删除的数据主键，例：[1,2,3]
}
