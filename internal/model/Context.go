package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Context 请求上下文结构
type Context struct {
	// Session *ghttp.Session // 当前Session管理对象
	User *ContextUser // 上下文用户信息
	Data g.Map        // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	// UserId   uint   // 用户ID
	// UserName string // 用户账号
	// Token string // 用户token
	// UserId    uint        `json:"userId"    ` // 用户ID
	// UserName  string      `json:"userName"  ` // 用户账号
	// NickName  string      `json:"nickName"  ` // 用户昵称
	// Mobile    string      `json:"mobile"    ` // 手机号
	// Avatar    string      `json:"avatar"    ` // 用户头像地址
	// Status    string      `json:"status"    ` // 用户状态；0:禁用,1:正常
	// RoleIds   string      `json:"roleIds"   ` // 角色id字符串
	// ShopIds   string      `json:"shopIds"   ` // 店铺ID字符串
	// DeptId    string      `json:"deptId"    ` // 部门ID
	// Remark    string      `json:"remark"    ` // 备注
	// CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	// UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	// DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	*SysUserInfoOutput
}
