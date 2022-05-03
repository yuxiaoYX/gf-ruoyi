package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取定时任务列表请求
type SysJobListReq struct {
	g.Meta    `path:"/job/getList" method:"post" summary:"获取定时任务列表" tags:"定时任务"`
	JobName   string `v:"max-length:64#任务名称长度错误！" dc:"任务名称"`          // 任务名称
	JobGroup  string `v:"max-length:64#任务组名长度错误！" dc:"任务组名"`          // 任务组名
	Status    string `v:"max-length:1#状态长度错误！" dc:"状态（0正常 1暂停）"`      // 状态（0正常 1暂停）
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取定时任务列表响应
type SysJobListRes struct {
	Rows  []*entity.SysJob `json:"rows"`  // 列表
	Total int              `json:"total"` // 数据总数
}

// 单个定时任务信息请求
type SysJobOneReq struct {
	g.Meta `path:"/job/getOne" method:"post" summary:"获取单个定时任务信息" tags:"定时任务"`
	JobId  int `v:"required|length:1,20#定时任务id不能为空！|定时任务id长度为:{min}到{max}位" dc:"任务id"` // 任务ID
}

// 单个定时任务信息响应
type SysJobOneRes struct {
	*entity.SysJob
}

// 新增定时任务请求
type SysJobCreateReq struct {
	g.Meta         `path:"/job/create" method:"post" summary:"新增定时任务" tags:"定时任务"`
	JobName        string `v:"required|length:1,64#任务名称不能为空！|任务名称长度为:{min}到{max}位" dc:"任务名称"`           // 任务名称
	JobGroup       string `v:"required|length:1,64#任务组名不能为空！|任务组名长度为:{min}到{max}位" dc:"任务组名"`           // 任务组名
	InvokeTarget   string `v:"required|length:1,500#调用目标字符串不能为空！|调用目标字符串长度为:{min}到{max}位" dc:"调用目标字符串"` // 调用目标字符串
	JobParams      string `v:"max-length:200#参数长度错误" dc:"参数"`                                           // 参数
	CronExpression string `v:"max-length:255#cron执行表达式长度错误" dc:"cron执行表达式"`                             // cron执行表达式
	MisfirePolicy  string `v:"max-length:20#计划执行错误策略长度错误" dc:"计划执行错误策略（1立即执行 2执行一次 3放弃执行）"`             // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Status         string `v:"max-length:1#状态长度错误" dc:"状态（0正常 1暂停）"`                                    // 状态（0正常 1暂停）
	Remark         string `v:"max-length:500#备注信息长度错误" dc:"备注信息"`                                       // 备注信息
}

// 新增定时任务响应
type SysJobCreateRes struct{}

// 更新定时任务信息请求
type SysJobUpdateReq struct {
	g.Meta         `path:"/job/update" method:"post" summary:"修改定时任务" tags:"定时任务"`
	JobId          int64  `v:"required|length:1,20#任务ID不能为空！|任务ID长度为:{min}到{max}位" dc:"任务ID"`           // 任务ID
	JobName        string `v:"required|length:1,64#任务名称不能为空！|任务名称长度为:{min}到{max}位" dc:"任务名称"`           // 任务名称
	JobGroup       string `v:"required|length:1,64#任务组名不能为空！|任务组名长度为:{min}到{max}位" dc:"任务组名"`           // 任务组名
	InvokeTarget   string `v:"required|length:1,500#调用目标字符串不能为空！|调用目标字符串长度为:{min}到{max}位" dc:"调用目标字符串"` // 调用目标字符串
	JobParams      string `v:"max-length:200#参数长度错误" dc:"参数"`                                           // 参数
	CronExpression string `v:"max-length:255#cron执行表达式长度错误" dc:"cron执行表达式"`                             // cron执行表达式
	MisfirePolicy  string `v:"max-length:20#计划执行错误策略长度错误" dc:"计划执行错误策略（1立即执行 2执行一次 3放弃执行）"`             // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Status         string `v:"max-length:1#状态长度错误" dc:"状态（0正常 1暂停）"`                                    // 状态（0正常 1暂停）
	Remark         string `v:"max-length:500#备注信息长度错误" dc:"备注信息"`                                       // 备注信息
}

// 更新定时任务信息响应
type SysJobUpdateRes struct{}

// 删除定时任务请求
type SysJobDeleteReq struct {
	g.Meta    `path:"/job/delete" method:"post" summary:"删除定时任务" tags:"定时任务"`
	JobIdList []int64 `v:"required#定时任务id不能为空！" dc:"定时任务id"` // 需要删除的数据主键，例：[1,2,3]
}

// 删除定时任务响应
type SysJobDeleteRes struct{}
