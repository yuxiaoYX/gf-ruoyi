package model

// 删除在线用户输入参数
type SysUserOnlineDeleteInput struct {
	Id       uint64 // 在线用户id
	UserName string // 用户名
}
