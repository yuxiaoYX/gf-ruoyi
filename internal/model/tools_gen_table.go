package model

// 获取当前数据库所有表输出
type SysGenTablesOutput struct {
	TableName string `json:"tableName"`
}

// 获取当前表所有字段输入
type SysGenColumnInput struct {
	TableName string // 表名
}

// 获取当前表所有字段输出
type SysGenColumnOutput struct {
	ColumnComment string // 注释
	ColumnName    string // 字段名
	DataType      string // 字段类型
	DataTypeLong  string // 字段长度
}
