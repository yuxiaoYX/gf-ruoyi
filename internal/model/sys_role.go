package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 获取角色列表输入
type SysRoleListInput struct {
	RoleName  string // 角色名称
	RoleKey   string // 权限字符
	Status    string // 角色状态；0:禁用，1:正常
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
	RoleId    uint        // 角色ID
	RoleName  string      // 角色名称
	RoleKey   string      // 权限字符
	RoleSort  int         // 显示顺序
	Status    string      // 角色状态；0:禁用，1:正常
	Remark    string      // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}

// 新增角色输入
type SysRoleCreateInput struct {
	RoleName string // 角色名称
	RoleKey  string // 权限字符
	RoleSort int    // 显示顺序
	Status   string // 角色状态；0:禁用，1:正常
	Remark   string // 备注
}

// 更新角色信息输入
type SysRoleUpdateInput struct {
	RoleId   uint   // 角色ID
	RoleName string // 角色名称
	RoleKey  string // 权限字符
	RoleSort int    // 显示顺序
	Status   string // 角色状态；0:禁用，1:正常
	Remark   string // 备注
}

// 删除角色输入
type SysRoleDeleteInput struct {
	RoleIdStr string // 需要删除的数据主键，例：1,2,3
}
