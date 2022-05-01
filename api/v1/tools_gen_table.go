package v1

import (
	"gf-ruoyi/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取当前数据库所有表请求
type SysGenTablesReq struct {
	g.Meta `path:"/gen/getTables" method:"get" summary:"获取当前数据库所有表" tags:"代码生成"`
}

// 获取当前数据库所有表响应
type SysGenTablesRes struct {
	Tables []*model.SysGenTableInfo `json:"tables"`
}

// 获取当前表所有字段请求
type SysGenColumnsReq struct {
	g.Meta    `path:"/gen/getColumn" method:"get" summary:"获取当前表所有字段" tags:"代码生成"`
	TableName string `v:"required|length:1,100#表名不能为空！|表名长度为:{min}到{max}位" dc:"表名"` // 表名
}

// 获取当前表所有字段响应
type SysGenColumnsRes struct {
	TableInfo  *model.SysGenTableInfo    `json:"tableInfo"`
	ColumnList []*model.SysGenColumnInfo `json:"columnList"`
}

// 预览代码请求
type SysGenPreviewCodeReq struct {
	g.Meta             `path:"/gen/preview" method:"post" summary:"预览代码" tags:"代码生成"`
	StructName         string                    `v:"required#Struct名称不能为空！" dc:"Struct名称"`     // Struct名称
	Description        string                    `v:"required#Struct中文名称不能为空！" dc:"Struct中文名称"` // Struct中文名称
	ModuleName         string                    `v:"required#Struct模块名不能为空！" dc:"模块名"`         // 模块名
	BusinessName       string                    `v:"required#Struct业务名不能为空！" dc:"业务名"`         // 业务名
	ApiFile            string                    `v:"required#api文件夹不能为空！" dc:"api文件夹"`         // api文件夹
	TplCategory        string                    `v:"required#使用的模板不能为空！" dc:"使用的模板"`           // 使用的模板（crud单表操作 tree树表操作）
	AutoCreateApiToSql bool                      `dc:"自动创建api"`                                 // 自动创建api
	AutoMoveFile       bool                      `dc:"自动移动文件"`                                  // 自动移动文件
	Fields             []*model.SysGenColumnInfo `dc:"列信息"`                                     // 列信息
}

// 预览代码响应
type SysGenPreviewCodeRes struct {
	ServerApi        string `json:"server_api"`
	ServerController string `json:"server_controller"`
	ServerModel      string `json:"server_model"`
	ServerService    string `json:"server_server"`
	WebApi           string `json:"web_api"`
	WebVueList       string `json:"web_vueList"`
	Sql              string `json:"sql"`
}
