package model

import (
	"gf-ruoyi/internal/model/entity"

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
	// UserId    uint        `json:"userId"    ` // 用户ID
	// UserName  string      `json:"userName"  ` // 用户账号
	// NickName  string      `json:"nickName"  ` // 用户昵称
	// Password  string      `json:"password"  ` // 登录密码
	// Mobile    string      `json:"mobile"    ` // 手机号码
	// Avatar    string      `json:"avatar"    ` // 用户头像地址
	// Status    string      `json:"status"    ` // 用户状态；0:禁用,1:正常
	// DeptId    string      `json:"deptId"    ` // 部门id
	// Remark    string      `json:"remark"    ` // 备注
	// LoginIp   string      `json:"loginIp"   ` // 最后登录IP
	// LoginDate *gtime.Time `json:"loginDate" ` // 最后登录时间
	// CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	// UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	// DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	*entity.SysUser
}
