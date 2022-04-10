package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取角色列表输入
type SysRoleListInput struct {
	RoleName  string // 角色名称
	Status    string // 角色状态；0:正常，1:禁用
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取角色列表输出
type SysRoleListOutput struct {
	Rows  []*SysRoleOneOutput // 列表
	Total int                 // 数据总数
}

// 获取单个角色信息输入
type SysRoleOneInput struct {
	RoleId uint // 角色ID
}

// 获取单个角色信息输出
type SysRoleOneOutput struct {
	*entity.SysRole
}

// 新增角色输入
type SysRoleCreateInput struct {
	RoleName string // 角色名称
	RoleSort int    // 显示顺序
	Status   string // 角色状态；0:正常，1:禁用
	Remark   string // 备注
	MenuIds  []int  // 菜单选中id列表
}

// 更新角色信息输入
type SysRoleUpdateInput struct {
	RoleId   uint   // 角色ID
	RoleName string // 角色名称
	RoleSort int    // 显示顺序
	Status   string // 角色状态；0:正常，1:禁用
	Remark   string // 备注
	MenuIds  []int  // 菜单选中id列表
}

// 删除角色输入
type SysRoleDeleteInput struct {
	RoleIdStr []int // 需要删除的数据主键，例：[1,2,3]
}

// 查询角色已授权用户列表请求
type SysRoleAllocatedListReq struct {
	NickName string // 用户昵称
	Mobile   string // 手机号码
	RoleId   uint   // 角色ID
	PageNum  int    // 分页码
	PageSize int    // 分页数量
}

// 查询角色已授权用户列表响应
type SysRoleAllocatedListRes struct {
	Rows  []*entity.SysUser `json:"rows"`  // 列表
	Total int               `json:"total"` // 数据总数
}
