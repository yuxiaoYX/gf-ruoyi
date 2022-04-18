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
	g.Meta   `path:"/job/create" method:"post" summary:"新增定时任务" tags:"定时任务"`
	JobName  string `v:"required|length:1,100#定时任务名称不能为空！|定时任务名称长度为:{min}到{max}位" dc:"定时任务名称"` // 定时任务名称
	JobKey   string `v:"required|length:1,100#定时任务键名不能为空！|定时任务键名长度为:{min}到{max}位" dc:"定时任务键名"` // 定时任务键名
	JobValue string `v:"required|length:1,500#定时任务键值不能为空！|定时任务键值长度为:{min}到{max}位" dc:"定时任务键值"` // 定时任务键值
	JobType  string `v:"max-length:1#系统内置长度错误" dc:"系统内置（Y是 N否）"`                               // 系统内置（Y是 N否）
	Remark   string `dc:"系统内置（Y是 N否）"`                                                         // 备注

	JobId          int64  `json:"jobId"          ` // 任务ID
	JobName        string `json:"jobName"        ` // 任务名称
	JobGroup       string `json:"jobGroup"       ` // 任务组名
	InvokeTarget   string `json:"invokeTarget"   ` // 调用目标字符串
	JobParams      string `json:"jobParams"      ` // 参数
	CronExpression string `json:"cronExpression" ` // cron执行表达式
	MisfirePolicy  string `json:"misfirePolicy"  ` // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Status         string `json:"status"         ` // 状态（0正常 1暂停）
	Remark         string `json:"remark"         ` // 备注信息
}

// 新增定时任务响应
type SysJobCreateRes struct{}

// 更新定时任务信息请求
type SysJobUpdateReq struct {
	g.Meta   `path:"/job/update" method:"post" summary:"修改定时任务" tags:"定时任务"`
	JobId    int    `v:"required|length:1,5#定时任务id不能为空！|定时任务id长度为:{min}到{max}位" dc:"定时任务id"`   // 定时任务主键
	JobName  string `v:"required|length:1,100#定时任务名称不能为空！|定时任务名称长度为:{min}到{max}位" dc:"定时任务名称"` // 定时任务名称
	JobKey   string `v:"required|length:1,100#定时任务键名不能为空！|定时任务键名长度为:{min}到{max}位" dc:"定时任务键名"` // 定时任务键名
	JobValue string `v:"required|length:1,500#定时任务键值不能为空！|定时任务键值长度为:{min}到{max}位" dc:"定时任务键值"` // 定时任务键值
	JobType  string `v:"max-length:1#系统内置长度错误" dc:"系统内置（Y是 N否）"`                               // 系统内置（Y是 N否）
	Remark   string `dc:"系统内置（Y是 N否）"`                                                         // 备注
}

// 更新定时任务信息响应
type SysJobUpdateRes struct{}

// 删除定时任务请求
type SysJobDeleteReq struct {
	g.Meta   `path:"/job/delete" method:"post" summary:"删除定时任务" tags:"定时任务"`
	JobIdStr string `v:"required#定时任务id不能为空！" dc:"定时任务id"` // 需要删除的数据主键，例：1,2,3
}

// 删除定时任务响应
type SysJobDeleteRes struct{}
