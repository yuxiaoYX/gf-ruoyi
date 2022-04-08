package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 用户登录输入
type SysUserLoginInput struct {
	UserName string // 用户名或用户昵称
	Password string // 密码
}

// 用户登录输出
type SysUserLoginOutput struct {
	UserId   uint   // 用户ID
	UserName string // 用户名
}

// 获取用户列表输入
type SysUserListInput struct {
	UserName  string // 用户账号
	NickName  string // 用户昵称
	Status    string // 用户状态；0:正常,1:停用
	DeptId    string // 部门id
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取用户列表输出
type SysUserListOutput struct {
	Rows  []*SysUserOneOutput // 列表
	Total int                 // 数据总数
}

// 获取单个用户信息输入
type SysUserOneInput struct {
	UserId uint // 用户ID
}

// 获取单个用户信息输出
type SysUserOneOutput struct {
	*entity.SysUser
}

// 新增用户输入
type SysUserCreateInput struct {
	UserName string // 用户账号
	NickName string // 用户昵称
	Password string // 登录密码
	Mobile   string // 手机号码
	Avatar   string // 用户头像地址
	Status   string // 用户状态；0:正常,1:停用
	DeptId   string // 部门id
	Remark   string // 备注
}

// 更新用户信息输入
type SysUserUpdateInput struct {
	UserId   uint   // 用户ID
	UserName string // 用户账号
	NickName string // 用户昵称
	Password string // 登录密码
	Mobile   string // 手机号码
	Avatar   string // 用户头像地址
	Status   string // 用户状态；0:正常,1:停用
	DeptId   string // 部门id
	Remark   string // 备注
}

// 删除用户输入
type SysUserDeleteInput struct {
	UserIdStr string // 需要删除的数据主键，例：1,2,3
}
