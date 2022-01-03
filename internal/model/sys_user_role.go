package model

import "github.com/gogf/gf/v2/util/gmeta"

type SysUserRoleOutput struct {
	gmeta.Meta `orm:"table:sys_user_role"`
	UserId     int `json:"userId" ` // 用户ID
	RoleId     int `json:"roleId" ` // 角色ID
}
