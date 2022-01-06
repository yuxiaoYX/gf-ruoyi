package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

// with特性，通过userId获取对应的角色信息
type SysUserRoleWithOutput struct {
	gmeta.Meta `orm:"table:sys_user_role"`
	UserId     int                  `json:"userId" ` // 用户ID
	RoleId     int                  `json:"roleId" ` // 角色ID
	Roles      []*SysRoleWithOutput `orm:"with:role_id"`
}
