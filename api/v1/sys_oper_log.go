package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取操作日志列表请求
type SysOperLogListReq struct {
	g.Meta       `path:"/operLog/getList" method:"post" summary:"获取操作日志列表" tags:"操作日志"`
	Title        string `v:"max-length:50#模块标题长度错误！" dc:"模块标题"`                     // 模块标题
	BusinessType int    `v:"max-length:2#业务类型长度错误！" dc:"业务类型（0其它 1新增 2删除 3修改 4查询）"` // 业务类型（0其它 1新增 2删除 3修改 4查询）
	OperName     string `v:"max-length:50#操作人员长度错误！" dc:"操作人员"`                     // 操作人员
	Status       string `v:"max-length:1#操作状态长度错误！" dc:"操作状态（0正常 1异常）"`             // 操作状态（0正常 1异常）
	BeginTime    string `dc:"开始时间"`                                                 // 开始时间
	EndTime      string `dc:"结束时间"`                                                 // 结束时间
	PageNum      int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`                      // 分页码
	PageSize     int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"`            // 分页数量
}

// 获取操作日志列表响应
type SysOperLogListRes struct {
	Rows  []*entity.SysOperLog `json:"rows"`  // 列表
	Total int                  `json:"total"` // 数据总数
}

// 删除操作日志请求
type SysOperLogDeleteReq struct {
	g.Meta    `path:"/operLog/delete" method:"post" summary:"删除操作日志" tags:"操作日志"`
	OperIdStr string `v:"required#操作日志id不能为空！" dc:"操作日志id"` // 需要删除的数据主键，例：1,2,3
}

// 删除操作日志响应
type SysOperLogDeleteRes struct{}

// 清空操作日志请求
type SysOperLogCleanReq struct {
	g.Meta `path:"/operLog/clean" method:"get" summary:"清空操作日志" tags:"操作日志"`
}

// 清空操作日志响应
type SysOperLogCleanRes struct{}
