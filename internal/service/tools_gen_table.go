package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sGenTable struct{}

// 代码生成管理服务
func SysGenTable() *sGenTable {
	return &sGenTable{}
}

// 获取当前数据库所有表
func (s *sGenTable) GetGenTables(ctx context.Context) (out model.SysGenTablesOutput, err error) {
	db := g.DB()
	sql := "select * from information_schema.tables where table_schema = 'gf_ruoyi'"
	err = db.GetScan(ctx, &out.Tables, sql)
	return
}

// 获取当前表所有字段
func (s *sGenTable) GetGenColumns(ctx context.Context, in model.SysGenColumnInput) (out model.SysGenColumnOutput, err error) {
	db := g.DB()
	// 表信息
	sql := "select * from information_schema.tables where table_schema = 'gf_ruoyi'"
	err = db.GetScan(ctx, &out.TableInfo, sql)
	if err != nil {
		return
	}
	// 表名转换成大驼峰
	out.TableInfo.TableName = gstr.CaseCamel(out.TableInfo.TableName)
	// 字段信息
	fieldMap, err := db.TableFields(ctx, in.TableName, "gf_ruoyi")
	if err != nil {
		return
	}
	columnList := make([]*model.SysGenColumnInfo, len(fieldMap))
	for _, v := range fieldMap {
		column := s.initColumnField(v)
		columnList[v.Index] = &column
	}
	out.ColumnList = columnList

	return
}

// 预览代码
func (s *sGenTable) PreviewCode(ctx context.Context, in model.SysGenPreviewCodeInput) (out *model.SysGenPreviewCodeOutput, err error) {
	view := gview.New()
	view.SetConfigWithMap(g.Map{
		// "Paths":      g.Cfg().MustGet(ctx, "gen.templatePath").String(),
		"Delimiters": []string{"{{", "}}"},
	})
	view.BindFuncMap(g.Map{
		// "UcFirst": func(str string) string { // 首字母大写
		// 	return gstr.UcFirst(str)
		// },
		// "Sum": func(a, b int) int {
		// 	return a + b
		// },
		"CaseCamelLower": gstr.CaseCamelLower, //首字母小写驼峰
		"CaseCamel":      gstr.CaseCamel,      //首字母大写驼峰
		// "HasSuffix":      gstr.HasSuffix,      //是否存在后缀
		// "ContainsI":      gstr.ContainsI,      //是否包含子字符串
		// "VueTag": func(t string) string {
		// 	return t
		// },
		"ProjectName": func() string {
			return g.Cfg().MustGet(ctx, "gen.projectName").String()
		},
	})

	tplData := g.Map{"table": in}
	// api代码
	apiValue := ""
	var tmpApi string
	if tmpApi, err = view.Parse(ctx, "vm/go/api.template", tplData); err == nil {
		apiValue = tmpApi
		apiValue, _ = s.trimBreak(apiValue)
	} else {
		return
	}
	// controller代码
	controllerValue := ""
	var tmpController string
	if tmpController, err = view.Parse(ctx, "vm/go/controller.template", tplData); err == nil {
		controllerValue = tmpController
		controllerValue, _ = s.trimBreak(controllerValue)
	} else {
		return
	}
	out = &model.SysGenPreviewCodeOutput{
		ServerApi:        apiValue,
		ServerController: controllerValue,
	}
	return
}

// 初始化列属性字段
func (s *sGenTable) initColumnField(field *gdb.TableField) (column model.SysGenColumnInfo) {
	column.ColumnName = field.Name
	column.ColumnType = field.Type
	column.ColumnComment = field.Comment
	column.GoType = s.generateStructFieldDefinition(field)
	column.GoField = gstr.CaseCamel(field.Name)
	column.HtmlField = gstr.CaseCamelLower(field.Name)
	if field.Key == "PRI" {
		column.IsPk = "1"
	} else {
		column.IsPk = "0"
	}
	if field.Extra == "auto_increment" {
		column.IsIncrement = "1"
	} else {
		column.IsIncrement = "0"
	}
	if field.Null {
		column.IsRequired = "0"
	} else {
		column.IsRequired = "1"
	}
	column.Sort = field.Index + 1
	if column.GoType == "string" && s.getColumnLength(field.Type) >= 500 {
		column.HtmlType = "textarea"
	} else if gstr.ContainsI(column.GoType, "time") {
		column.HtmlType = "datetime"
	} else {
		column.HtmlType = "input"
	}
	return
}

// 返回字段的属性定义
func (s *sGenTable) generateStructFieldDefinition(field *gdb.TableField) string {
	var (
		typeName string
	)
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint"
		} else {
			typeName = "int"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}

	case "real":
		typeName = "float32"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float64"

	case "bool":
		typeName = "bool"

	case "datetime", "timestamp", "date", "time":
		typeName = "*gtime.Time"
	case "json", "jsonb":
		typeName = "*gjson.Json"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "int"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "float64"
		case strings.Contains(t, "bool"):
			typeName = "bool"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "[]byte"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			typeName = "*gtime.Time"
		default:
			typeName = "string"
		}
	}

	return typeName
}

// GetColumnLength 获取字段长度
func (s *sGenTable) getColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := ""
	if start >= 0 && end >= 0 {
		result = columnType[start+1 : end-1]
	}
	return gconv.Int(result)
}

//剔除多余的换行
func (s *sGenTable) trimBreak(str string) (rStr string, err error) {
	var b []byte
	if b, err = gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str)); err != nil {
		return
	}
	if b, err = gregex.Replace("(([\\s\t]*)/{4}\r?\n)", []byte("$2\n\n"), b); err == nil {
		rStr = gconv.String(b)
	}
	return
}
