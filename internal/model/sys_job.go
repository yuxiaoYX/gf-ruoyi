package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取定时任务列表输入
type SysJobListInput struct {
	JobName   string // 任务名称
	JobGroup  string // 任务组名
	Status    string // 状态（0正常 1暂停）
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取定时任务列表输出
type SysJobListOutput struct {
	Rows  []*SysJobOneOutput // 列表
	Total int                // 数据总数
}

// 获取单个定时任务信息输入
type SysJobOneInput struct {
	JobId int // 定时任务ID
}

// 获取单个定时任务信息输出
type SysJobOneOutput struct {
	*entity.SysJob
}

// 新增定时任务输入
type SysJobCreateInput struct {
	JobName        string // 任务名称
	JobGroup       string // 任务组名
	InvokeTarget   string // 调用目标字符串
	JobParams      string // 参数
	CronExpression string // cron执行表达式
	MisfirePolicy  string // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Status         string // 状态（0正常 1暂停）
	Remark         string // 备注信息
}

// 更新定时任务信息输入
type SysJobUpdateInput struct {
	JobId          int64  // 任务ID
	JobName        string // 任务名称
	JobGroup       string // 任务组名
	InvokeTarget   string // 调用目标字符串
	JobParams      string // 参数
	CronExpression string // cron执行表达式
	MisfirePolicy  string // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Status         string // 状态（0正常 1暂停）
	Remark         string // 备注信息
}

// 删除定时任务输入
type SysJobDeleteInput struct {
	JobIdStr string // 需要删除的数据主键，例：[1,2,3]
}
