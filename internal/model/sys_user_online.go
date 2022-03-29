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

// 删除在线用户
type SysUserOnlineDeleteInput struct {
	Id     uint64
	Token  string // 用户token
	UserId int    // 用户id
}

// 查询token是否在线返回
type SysUserOnlineGetTokenOutput struct {
	Token string
}
