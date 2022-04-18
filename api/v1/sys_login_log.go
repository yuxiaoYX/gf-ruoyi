package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取登录日志列表请求
type SysLoginLogListReq struct {
	g.Meta    `path:"/loginLog/getList" method:"post" summary:"获取登录日志列表" tags:"登录日志"`
	UserName  string `v:"max-length:50#登录账号长度错误！" dc:"登录账号"`          // 登录账号
	Ipaddr    string `v:"max-length:128#登录IP地址长度错误！" dc:"登录IP地址"`     // 登录IP地址
	Status    string `v:"max-length:1#登录状态长度错误！" dc:"登录状态（0成功 1失败）"`  // 登录状态（0成功 1失败）
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取登录日志列表响应
type SysLoginLogListRes struct {
	Rows  []*entity.SysLoginLog `json:"rows"`  // 列表
	Total int                   `json:"total"` // 数据总数
}

// 删除登录日志请求
type SysLoginLogDeleteReq struct {
	g.Meta    `path:"/loginLog/delete" method:"post" summary:"删除登录日志" tags:"登录日志"`
	InfoIdStr string `v:"required#访问id不能为空！" dc:"登录日志id"` // 需要删除的数据主键，例：1,2,3
}

// 删除登录日志响应
type SysLoginLogDeleteRes struct{}

// 清空登录日志请求
type SysLoginLogCleanReq struct {
	g.Meta `path:"/loginLog/clean" method:"get" summary:"清空登录日志" tags:"登录日志"`
}

// 清空登录日志响应
type SysLoginLogCleanRes struct{}
