package model

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	// Session *ghttp.Session // 当前Session管理对象
	UserInfo *ContextUser // 上下文用户信息
	// Roles    *ContextRoles // 上下文角色字段列表
	Data g.Map // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	User      *entity.SysUser   // 用户信息
	Dept      *entity.SysDept   // 部门信息
	RoleIds   []uint            // 角色id列表
	RoleNames []string          // 角色名称
	Roles     []*entity.SysRole // 角色信息列表
}
