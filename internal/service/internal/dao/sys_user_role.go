// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-RuoYi/internal/service/internal/dao/internal"
)

// sysUserRoleDao is the data access object for table sys_user_role.
// You can define custom methods on it to extend its functionality as you wish.
type sysUserRoleDao struct {
	*internal.SysUserRoleDao
}

var (
	// SysUserRole is globally public accessible object for table sys_user_role operations.
	SysUserRole = sysUserRoleDao{
		internal.NewSysUserRoleDao(),
	}
)

// Fill with you ideas below.