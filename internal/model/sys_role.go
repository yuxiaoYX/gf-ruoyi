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
}

// 更新角色信息输入
type SysRoleUpdateInput struct {
	RoleId   uint   // 角色ID
	RoleName string // 角色名称
	RoleSort int    // 显示顺序
	Status   string // 角色状态；0:正常，1:禁用
	Remark   string // 备注
}

// 删除角色输入
type SysRoleDeleteInput struct {
	RoleIdStr string // 需要删除的数据主键，例：1,2,3
}
