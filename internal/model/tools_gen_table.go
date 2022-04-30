package model

// 获取当前数据库所有表输出
type SysGenTablesOutput struct {
	Tables []*SysGenTableInfo
}

// 获取当前表所有字段输入
type SysGenColumnInput struct {
	TableName string // 表名
}

// 获取当前表所有字段输出
type SysGenColumnOutput struct {
	TableInfo  *SysGenTableInfo
	ColumnList []*SysGenColumnInfo
}

// 预览代码输入
type SysGenPreviewCodeInput struct {
	*SysGenAutoCodeInfo
}

// 预览代码输出
type SysGenPreviewCodeOutput struct {
	ServerApi        string `json:"server_api"`
	ServerController string `json:"server_controller"`
	ServerModel      string `json:"server_model"`
	ServerService    string `json:"server_server"`
	WebApi           string `json:"web_api"`
	WebVueList       string `json:"web_vueList"`
	Sql              string `json:"sql"`
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
	IsInsert      string `json:"isInsert"`      // 是否为插入字段（1是）
	IsEdit        string `json:"isEdit"`        // 是否编辑字段（1是）
	IsList        string `json:"isList"`        // 是否列表字段（1是）
	IsQuery       string `json:"isQuery"`       // 是否查询字段（1是）
	QueryType     string `json:"queryType"`     // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string `json:"htmlType"`      // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string `json:"dictType"`      // 字典类型
	Sort          int    `json:"sort"`          // 排序
}

// 表信息
type SysGenAutoCodeInfo struct {
	StructName         string              // Struct名称
	Abbreviation       string              // Struct简称
	Description        string              // Struct中文名称
	ApiFile            string              // api文件夹
	TplCategory        string              // 使用的模板（crud单表操作 tree树表操作）
	AutoCreateApiToSql bool                // 自动创建api
	AutoMoveFile       bool                // 自动移动文件
	Fields             []*SysGenColumnInfo // 列信息
}
