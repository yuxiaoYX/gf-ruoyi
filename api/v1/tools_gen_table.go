package v1

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// 获取当前数据库所有表请求
type SysGenTablesReq struct {
	g.Meta `path:"/gen/getTables" method:"get" summary:"获取当前数据库所有表" tags:"代码生成"`
}

// 获取当前数据库所有表响应
type SysGenTablesRes struct {
	Tables []map[string]interface{} `json:"tables"`
}

// 获取当前表所有字段请求
type SysGenColumnsReq struct {
	g.Meta    `path:"/gen/getColumn" method:"get" summary:"获取当前表所有字段" tags:"代码生成"`
	TableName string `v:"required|length:1,100#表名不能为空！|表名长度为:{min}到{max}位" dc:"表名"` // 表名
}

// 获取当前表所有字段响应
type SysGenColumnsRes struct {
	Columns map[string]*gdb.TableField `json:"columns"`
}
