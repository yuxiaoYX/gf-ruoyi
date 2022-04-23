package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 获取当前数据库所有表请求
type SysGenTablesReq struct {
	g.Meta `path:"/gen/getTables" method:"get" summary:"获取当前数据库所有表" tags:"代码生成"`
}

// 获取当前数据库所有表响应
type SysGenTablesRes struct {
	Tables []*SysGenTableInfo `json:"tables"`
}

// 获取当前表所有字段请求
type SysGenColumnsReq struct {
	g.Meta    `path:"/gen/getColumn" method:"get" summary:"获取当前表所有字段" tags:"代码生成"`
	TableName string `v:"required|length:1,100#表名不能为空！|表名长度为:{min}到{max}位" dc:"表名"` // 表名
}

// 获取当前表所有字段响应
type SysGenColumnsRes struct {
	TableInfo  *SysGenTableInfo    `json:"tableInfo"`
	ColumnList []*SysGenColumnInfo `json:"columnList"`
}

// 表信息
type SysGenTableInfo struct {
	TableName    string `json:"tableName"`
	TableComment string `json:"tableComment"`
}

// 字段信息
type SysGenColumnInfo struct {
	ColumnName    string `json:"columnName"`    // 字段名
	ColumnType    string `json:"columnType"`    // 字段类型
	ColumnComment string `json:"columnComment"` // 字段注释
	GoType        string `json:"goType"`        // go类型
	GoField       string `json:"goField"`       // go字段名
	HtmlField     string `json:"htmlField"`     // json字段名
	IsPk          string `json:"isPk"`          // 是否主键（1是）
	IsIncrement   string `json:"isIncrement"`   // 是否自增（1是）
	IsRequired    string `json:"isRequired"`    // 是否必填（1是）
	HtmlType      string `json:"htmlType"`      // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	Sort          int    `json:"sort"`          // 排序
}
