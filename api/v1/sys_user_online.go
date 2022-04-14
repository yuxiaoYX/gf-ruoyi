package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 获取登录日志列表请求
type SysUserOnlineListReq struct {
	g.Meta    `path:"/loginLog/getList" method:"post" summary:"获取登录日志列表" tags:"登录日志"`
	UserName  string `v:"max-length:50#登录账号长度错误！" dc:"登录账号"`         // 登录账号
	Ipaddr    string `v:"max-length:128#登录IP地址长度错误！" dc:"登录IP地址"`    // 登录IP地址
	Status    string `v:"max-length:1#登录状态长度错误！" dc:"登录状态（0成功 1失败）"` // 登录状态（0成功 1失败）
	BeginTime string `dc:"开始时间"`                                     // 开始时间
	EndTime   string `dc:"结束时间"`                                     // 结束时间

	Id        uint64      `json:"id"        ` //
	Token     string      `json:"token"     ` // 用户token
	UserId    int         `json:"userId"    ` // 用户id
	UserName  string      `json:"userName"  ` // 用户名
	Ip        string      `json:"ip"        ` // 登录ip
	Explorer  string      `json:"explorer"  ` // 浏览器
	Os        string      `json:"os"        ` // 操作系统
	CreatedAt *gtime.Time `json:"createdAt" ` // 登录时间

	PageNum  int `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize int `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取登录日志列表响应
type SysUserOnlineListRes struct {
	Rows  []*entity.SysUserOnline `json:"rows"`  // 列表
	Total int                     `json:"total"` // 数据总数
}

// 删除登录日志请求
type SysUserOnlineDeleteReq struct {
	g.Meta    `path:"/loginLog/delete" method:"post" summary:"删除登录日志" tags:"登录日志"`
	InfoIdStr string `v:"required#访问id不能为空！" dc:"登录日志id"` // 需要删除的数据主键，例：1,2,3
}

// 删除登录日志响应
type SysUserOnlineDeleteRes struct{}
