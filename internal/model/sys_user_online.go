package model

import "gf-ruoyi/internal/model/entity"

// 获取在线用户列表请求
type SysUserOnlineListInput struct {
	UserName  string // 用户名
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取在线用户列表响应
type SysUserOnlineListOutput struct {
	Rows  []*entity.SysUserOnline `json:"rows"`  // 列表
	Total int                     `json:"total"` // 数据总数
}

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
	Ids []uint64 // 在线用户id列表
}

// 查询token是否在线输出
type SysUserOnlineGetTokenOutput struct {
	UserId int // 用户id
}
