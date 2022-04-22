package service

import (
	"context"
	"gf-ruoyi/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
func (s *sGenTable) GetGenColumns(ctx context.Context, in model.SysGenColumnInput) (out map[string]*gdb.TableField, err error) {
	db := g.DB()
	sql := "select table_name as table_name from information_schema.tables where table_schema = 'gf_ruoyi'"
	err = db.GetScan(ctx, &out, sql)
	return
}
