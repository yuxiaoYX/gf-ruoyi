// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-ruoyi/internal/service/internal/dao/internal"
)

// sysDictTypeDao is the data access object for table sys_dict_type.
// You can define custom methods on it to extend its functionality as you wish.
type sysDictTypeDao struct {
	*internal.SysDictTypeDao
}

var (
	// SysDictType is globally public accessible object for table sys_dict_type operations.
	SysDictType = sysDictTypeDao{
		internal.NewSysDictTypeDao(),
	}
)

// Fill with you ideas below.