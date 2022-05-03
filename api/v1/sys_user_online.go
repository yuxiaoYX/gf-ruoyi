package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取在线用户列表请求
type SysUserOnlineListReq struct {
	g.Meta    `path:"/online/getList" method:"post" summary:"获取在线用户列表" tags:"在线用户"`
	UserName  string `v:"max-length:50#用户名长度错误！" dc:"用户名"`            // 用户名
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取在线用户列表响应
type SysUserOnlineListRes struct {
	Rows  []*entity.SysUserOnline `json:"rows"`  // 列表
	Total int                     `json:"total"` // 数据总数
}

// 删除在线用户请求
type SysUserOnlineDeleteReq struct {
	g.Meta `path:"/online/delete" method:"post" summary:"删除在线用户" tags:"在线用户"`
	IdList []uint64 `v:"required#在线用户id不能为空！" dc:"在线用户id"` // 需要删除的数据主键，例：1,2,3
}

// 删除在线用户响应
type SysUserOnlineDeleteRes struct{}
