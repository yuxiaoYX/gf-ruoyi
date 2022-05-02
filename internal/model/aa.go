// ==========================================================================
// 自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-05-02 16:47:40
// ==========================================================================

package model

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// 分页获取参数配置表列表输入
type SysAaListInput struct {
	ConfigName  string      // 参数名称
	ConfigKey   string      // 参数键名
	ConfigValue string      // 参数键值
	CreatedAt   *gtime.Time // 创建时间
	BeginTime   string      // 开始时间
	EndTime     string      // 结束时间
	PageNum     int         // 分页码
	PageSize    int         // 分页数量
}

// 分页获取参数配置表列表输出
type SysAaListOutput struct {
	Rows  []*SysAaOneOutput `json:"rows"`  // 列表
	Total int               `json:"total"` // 数据总数
}

// 获取单个参数配置表信息输入
type SysAaOneInput struct {
	ConfigId int // 参数主键
}

// 获取单个参数配置表信息输出
type SysAaOneOutput struct {
	*entity.SysAa
}

// 新增参数配置表输入
type SysAaCreateInput struct {
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	ConfigType  string // 系统内置（Y是 N否）
	Remark      string // 备注
}

// 新增参数配置表输出
type SysAaCreateOutput struct{}

// 更新参数配置表信息输入
type SysAaUpdateInput struct {
	ConfigId    int    // 参数主键
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	ConfigType  string // 系统内置（Y是 N否）
	Remark      string // 备注
}

// 更新参数配置表信息输出
type SysAaUpdateOutput struct{}

// 删除参数配置表输入
type SysAaDeleteInput struct {
	ConfigIdStr string // 参数主键，例：1,2,3
}

// 删除参数配置表输出
type SysSysAaDeleteOutput struct{}
