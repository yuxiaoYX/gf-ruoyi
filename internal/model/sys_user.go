package model

// 用户登录输入
type SysUserLoginInput struct {
	UserName string // 用户名
	NickName string // 用户昵称
	Password string // 密码
}

// 用户登录输出
type SysUserLoginOutput struct {
	Msg   string // 提示消息
	Token string // 用户token
}
