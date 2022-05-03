package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取操作日志列表输入
type SysOperLogListInput struct {
	Title        string // 模块标题
	BusinessType int    // 业务类型（0其它 1新增 2删除 3修改 4查询）
	OperName     string // 操作人员
	Status       string // 操作状态（0正常 1异常）
	BeginTime    string // 开始时间
	EndTime      string // 结束时间
	PageNum      int    // 分页码
	PageSize     int    // 分页数量
}

// 获取操作日志列表输出
type SysOperLogListOutput struct {
	Rows  []*entity.SysOperLog // 列表
	Total int                  // 数据总数
}

// 删除操作日志输入
type SysOperLogDeleteInput struct {
	OperIdList []int64 // 需要删除的数据主键，例：[1,2,3]
}
