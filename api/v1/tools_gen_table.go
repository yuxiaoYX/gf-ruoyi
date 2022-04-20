package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取代码生成页列表请求
type SysToolsGenTableListReq struct {
	g.Meta       `path:"/gen/getList" method:"post" summary:"获取代码生成页列表" tags:"代码生成"`
	TableName    string `dc:"表名称"`                                       //表名称
	TableComment string `dc:"表描述"`                                       //表描述
	BeginTime    string `dc:"开始时间"`                                      // 开始时间
	EndTime      string `dc:"结束时间"`                                      // 结束时间
	PageNum      int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize     int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取代码生成页列表响应
type SysToolsGenTableListRes struct {
	Rows  []*entity.SysJob `json:"rows"`  // 列表
	Total int              `json:"total"` // 数据总数
}
