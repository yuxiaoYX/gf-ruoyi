package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取参数列表输入
type SysConfigListInput struct {
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	BeginTime   string // 开始时间
	EndTime     string // 结束时间
	PageNum     int    // 分页码
	PageSize    int    // 分页数量
}

// 获取参数列表输出
type SysConfigListOutput struct {
	Rows  []*SysConfigOneOutput // 列表
	Total int                   // 数据总数
}

// 获取单个参数信息输入
type SysConfigOneInput struct {
	ConfigId int // 参数ID
}

// 获取单个参数信息输出
type SysConfigOneOutput struct {
	*entity.SysConfig
}

// 新增参数输入
type SysConfigCreateInput struct {
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	ConfigType  string // 系统内置（Y是 N否）
	Remark      string // 备注
}

// 更新参数信息输入
type SysConfigUpdateInput struct {
	ConfigId    int    // 参数主键
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	ConfigType  string // 系统内置（Y是 N否）
	Remark      string // 备注
}

// 删除参数输入
type SysConfigDeleteInput struct {
	ConfigIdList []int // 需要删除的数据主键，例：[1,2,3]
}

// 根据参数键名查询参数值输入
type SysConfigKeyInput struct {
	ConfigKey string // 参数键名
}

// 根据参数键名查询参数值输出
type SysConfigKeyOutput struct {
	ConfigValue string `json:"configValue" ` // 参数键值
}
