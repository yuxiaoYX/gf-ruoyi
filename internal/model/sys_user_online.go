package model

// 添加在线用户输入参数
type SysUserOnlineCreateInput struct {
	Token    string // 用户token
	UserId   int    // 用户id
	UserName string // 用户名
	Ip       string // 登录ip
	Explorer string // 浏览器
	Os       string // 操作系统
}

// 删除在线用户输入参数
type SysUserOnlineDeleteInput struct {
	Id       uint64
	Token    string // 用户token
	UserName string // 用户名
}
