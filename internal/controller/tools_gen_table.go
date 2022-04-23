package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 代码生成管理
var SysGenTable = cGenTable{}

type cGenTable struct{}

// 获取当前数据库所有表
func (c *cGenTable) GetTableNames(ctx context.Context, req *v1.SysGenTablesReq) (res v1.SysGenTablesRes, err error) {
	Tables, err := service.SysGenTable().GetGenTables(ctx)
	gconv.Scan(Tables, &res.Tables)
	return
}

// 获取当前表所有字段
func (c *cGenTable) GetGenColumns(ctx context.Context, req *v1.SysGenColumnsReq) (res v1.SysGenColumnsRes, err error) {
	_, err = service.SysGenTable().GetGenColumns(ctx, model.SysGenColumnInput{TableName: req.TableName})
	return
}
