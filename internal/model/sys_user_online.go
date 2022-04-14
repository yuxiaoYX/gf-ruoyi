package model

// 添加在线用户输入参数
type SysUserOnlineCreateInput struct {
	Token         string `json:"token"         ` // 用户token
	UserId        int64  `json:"userId"        ` // 用户id
	UserName      string `json:"userName"      ` // 用户名
	Os            string `json:"os"            ` // 操作系统
	Ipaddr        string `json:"ipaddr"        ` // 登录IP地址
	LoginLocation string `json:"loginLocation" ` // 登录地点
	Browser       string `json:"browser"       ` // 浏览器类型
}

// 删除在线用户
type SysUserOnlineDeleteInput struct {
	Id     uint64
	Token  string // 用户token
	UserId int    // 用户id
}

// 查询token是否在线输出
type SysUserOnlineGetTokenOutput struct {
	UserId int // 用户id
}
