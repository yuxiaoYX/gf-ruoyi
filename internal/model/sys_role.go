package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type SysRoleWithOutput struct {
	gmeta.Meta `orm:"table:sys_role"`
	RoleId     uint        `json:"roleId"    ` // 角色ID
	RoleName   string      `json:"roleName"  ` // 角色名称
	RoleKey    string      `json:"roleKey"   ` // 权限字符
	RoleSort   int         `json:"roleSort"  ` // 显示顺序
	Status     string      `json:"status"    ` // 角色状态；0:禁用，1:正常
	Remark     string      `json:"remark"    ` // 备注
	CreatedAt  *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt" ` // 更新时间
}

type SysRoleGetIdsInput struct {
	RoleIds []int // 角色ID字符串
}
