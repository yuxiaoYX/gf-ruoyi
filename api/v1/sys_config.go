package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取参数列表请求
type SysConfigListReq struct {
	g.Meta      `path:"/config/getList" method:"post" summary:"获取参数列表" tags:"参数"`
	ConfigName  string `v:"max-length:100#参数名称长度错误！" dc:"参数名称"`         // 参数名称
	ConfigKey   string `v:"max-length:100#参数键名长度错误！" dc:"参数键名"`         // 参数键名
	ConfigValue string `v:"max-length:500#参数键值长度错误！" dc:"参数键值"`         // 参数键值
	BeginTime   string `dc:"开始时间"`                                      // 开始时间
	EndTime     string `dc:"结束时间"`                                      // 结束时间
	PageNum     int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize    int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取参数列表响应
type SysConfigListRes struct {
	Rows  []*SysConfigOneRes `json:"rows"`  // 列表
	Total int                `json:"total"` // 数据总数
}

// 单个参数信息请求
type SysConfigOneReq struct {
	g.Meta   `path:"/config/getOne" method:"post" summary:"获取单个参数信息" tags:"参数"`
	ConfigId int `v:"required|length:1,5#参数id不能为空！|参数id长度为:{min}到{max}位" dc:"参数id"` // 参数主键
}

// 单个参数信息响应
type SysConfigOneRes struct {
	*entity.SysConfig
}

// 新增参数请求
type SysConfigCreateReq struct {
	g.Meta      `path:"/config/create" method:"post" summary:"新增参数" tags:"参数"`
	ConfigName  string `v:"required|length:1,100#参数名称不能为空！|参数名称长度为:{min}到{max}位" dc:"参数名称"` // 参数名称
	ConfigKey   string `v:"required|length:1,100#参数键名不能为空！|参数键名长度为:{min}到{max}位" dc:"参数键名"` // 参数键名
	ConfigValue string `v:"required|length:1,500#参数键值不能为空！|参数键值长度为:{min}到{max}位" dc:"参数键值"` // 参数键值
	ConfigType  string `v:"max-length:1#系统内置长度错误" dc:"系统内置（Y是 N否）"`                         // 系统内置（Y是 N否）
	Remark      string `dc:"备注"`                                                            // 备注
}

// 新增参数响应
type SysConfigCreateRes struct{}

// 更新参数信息请求
type SysConfigUpdateReq struct {
	g.Meta      `path:"/config/update" method:"post" summary:"修改参数" tags:"参数"`
	ConfigId    int    `v:"required|length:1,5#参数id不能为空！|参数id长度为:{min}到{max}位" dc:"参数id"`   // 参数主键
	ConfigName  string `v:"required|length:1,100#参数名称不能为空！|参数名称长度为:{min}到{max}位" dc:"参数名称"` // 参数名称
	ConfigKey   string `v:"required|length:1,100#参数键名不能为空！|参数键名长度为:{min}到{max}位" dc:"参数键名"` // 参数键名
	ConfigValue string `v:"required|length:1,500#参数键值不能为空！|参数键值长度为:{min}到{max}位" dc:"参数键值"` // 参数键值
	ConfigType  string `v:"max-length:1#系统内置长度错误" dc:"系统内置（Y是 N否）"`                         // 系统内置（Y是 N否）
	Remark      string `dc:"备注"`                                                            // 备注
}

// 更新参数信息响应
type SysConfigUpdateRes struct{}

// 删除参数请求
type SysConfigDeleteReq struct {
	g.Meta       `path:"/config/delete" method:"post" summary:"删除参数" tags:"参数"`
	ConfigIdList []int `v:"required#参数id不能为空！" dc:"参数id"` // 需要删除的数据主键，例：[1,2,3]
}

// 删除参数响应
type SysConfigDeleteRes struct{}

// 根据参数键名查询参数值请求
type SysConfigKeyReq struct {
	g.Meta    `path:"/config/configKey" method:"get" summary:"参数键名查询参数值" tags:"参数"`
	ConfigKey string `v:"required|length:1,100#参数键名不能为空！|参数键名长度为:{min}到{max}位" dc:"参数键名"` // 参数键名
}

// 根据参数键名查询参数值响应
type SysConfigKeyRes struct {
	ConfigValue string `json:"configValue" ` // 参数键值
}
