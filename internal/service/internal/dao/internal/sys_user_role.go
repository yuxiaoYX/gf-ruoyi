// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserRoleDao is the data access object for table sys_user_role.
type SysUserRoleDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns SysUserRoleColumns // Columns contains all the column names of Table for convenient usage.
}

// SysUserRoleColumns defines and stores column names for table sys_user_role.
type SysUserRoleColumns struct {
	UserId  string // 用户ID  
    RoleId  string // 角色ID
}

//  sysUserRoleColumns holds the columns for table sys_user_role.
var sysUserRoleColumns = SysUserRoleColumns{
	UserId: "user_id",  
            RoleId: "role_id",
}

// NewSysUserRoleDao creates and returns a new DAO object for table data access.
func NewSysUserRoleDao() *SysUserRoleDao {
	return &SysUserRoleDao{
		Group:   "default",
		Table:   "sys_user_role",
		Columns: sysUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserRoleDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}