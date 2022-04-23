package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

type sGenTable struct{}

// 代码生成管理服务
func SysGenTable() *sGenTable {
	return &sGenTable{}
}

// 获取当前数据库所有表
func (s *sGenTable) GetGenTables(ctx context.Context) (out []*model.SysGenTablesOutput, err error) {
	db := g.DB()
	sql := "select table_name as table_name from information_schema.tables where table_schema = 'gf_ruoyi'"
	err = db.GetScan(ctx, &out, sql)
	return
}

// 获取当前表所有字段
func (s *sGenTable) GetGenColumns(ctx context.Context, in model.SysGenColumnInput) (out []*model.SysGenColumnOutput, err error) {
	db := g.DB()
	// sql := "select * from information_schema.columns where table_schema = 'gf_ruoyi' AND table_name = 'sys_user'"
	// err = db.GetScan(ctx, out, sql)

	fieldMap, err := db.TableFields(ctx, "sys_user", "gf_ruoyi")
	// array := make([]map[string]interface{}, len(fieldMap))
	out = make([]*model.SysGenColumnOutput, len(fieldMap))
	for _, v := range fieldMap {
		column, err := s.initColumnField(v)
		if err != nil {
			return nil, err
		}
		g.Log().Info(ctx, v)
		out[v.Index] = &column
	}
	g.Log().Info(ctx, &out)
	return
}

// 初始化列属性字段
func (s *sGenTable) initColumnField(field *gdb.TableField) (column model.SysGenColumnOutput, err error) {
	column.ColumnName = field.Name
	column.ColumnType = field.Type
	column.ColumnComment = field.Comment
	column.GoType = generateStructFieldDefinition(field)
	column.GoField = gstr.CaseCamel(field.Name)
	column.HtmlField = gstr.CaseCamelLower(field.Name)

	return
}

// generateStructFieldForModel generates and returns the attribute definition for specified field.
func generateStructFieldDefinition(field *gdb.TableField) string {
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
