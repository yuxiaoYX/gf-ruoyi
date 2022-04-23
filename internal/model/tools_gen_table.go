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
	ColumnName    string // 字段名
	ColumnType    string // 字段类型
	ColumnComment string // 字段注释
	GoType        string // go类型
	GoField       string // go字段名
	HtmlField     string // json字段名
	IsPk          string // 是否主键（1是）
	IsIncrement   string // 是否自增（1是）
	IsRequired    string // 是否必填（1是）
	IsInsert      string // 是否为插入字段（1是）
	IsEdit        string // 是否编辑字段（1是）
	IsList        string // 是否列表字段（1是）
	IsQuery       string // 是否查询字段（1是）
	QueryType     string // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string // 字典类型
	Sort          int    // 排序
}
