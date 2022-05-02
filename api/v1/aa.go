// ==========================================================================
// 自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2022-05-02 16:47:40
// ==========================================================================

package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 分页获取参数配置表列表请求
type SysAaListReq struct {
	g.Meta      `path:"/sysAa/getList" method:"post" summary:"分页获取参数配置表列表" tags:"参数配置表"`
	ConfigName  string `dc:"参数名称"`                                      // 参数名称
	ConfigKey   string `dc:"参数键名"`                                      // 参数键名
	ConfigValue string `dc:"参数键值"`                                      // 参数键值
	BeginTime   string `dc:"开始时间"`                                      // 开始时间
	EndTime     string `dc:"结束时间"`                                      // 结束时间
	PageNum     int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize    int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 分页获取参数配置表列表响应
type SysAaListRes struct {
	Rows  []*SysAaOneRes `json:"rows"`  // 列表
	Total int            `json:"total"` // 数据总数
}

// 获取单个参数配置表信息请求
type SysAaOneReq struct {
	g.Meta   `path:"/sysAa/getOne" method:"get" summary:"获取单个参数配置表信息" tags:"参数配置表"`
	ConfigId int `v:"required#参数主键不能为空！" dc:"参数主键"` // 参数主键
}

// 获取单个参数配置表信息响应
type SysAaOneRes struct {
	*entity.SysAa
}

// 新增参数配置表请求
type SysAaCreateReq struct {
	g.Meta      `path:"/sysAa/create" method:"post" summary:"新增参数配置表" tags:"参数配置表"`
	ConfigName  string `dc:"参数名称"`                                      // 参数名称
	ConfigKey   string `dc:"参数键名"`                                      // 参数键名
	ConfigValue string `dc:"参数键值"`                                      // 参数键值
	ConfigType  string `v:"required#系统内置（Y是 N否）不能为空！" dc:"系统内置（Y是 N否）"` // 系统内置（Y是 N否）
	Remark      string `dc:"备注"`                                        // 备注
}

// 新增参数配置表响应
type SysAaCreateRes struct{}

// 更新参数配置表信息请求
type SysAaUpdateReq struct {
	g.Meta      `path:"/sysAa/update" method:"post" summary:"修改参数配置表" tags:"参数配置表"`
	ConfigId    int    `v:"required#参数主键不能为空！" dc:"参数主键"`               // 参数主键
	ConfigName  string `dc:"参数名称"`                                      // 参数名称
	ConfigKey   string `dc:"参数键名"`                                      // 参数键名
	ConfigValue string `dc:"参数键值"`                                      // 参数键值
	ConfigType  string `v:"required#系统内置（Y是 N否）不能为空！" dc:"系统内置（Y是 N否）"` // 系统内置（Y是 N否）
	Remark      string `dc:"备注"`                                        // 备注
}

// 更新参数配置表信息响应
type SysAaUpdateRes struct{}

// 删除参数配置表请求
type SysAaDeleteReq struct {
	g.Meta      `path:"/sysAa/delete" method:"get" summary:"删除参数配置表" tags:"参数配置表"`
	ConfigIdStr string `v:"required#参数主键不能为空！" dc:"参数主键"` // 参数主键，例：1,2,3
}

// 删除参数配置表响应
type SysAaDeleteRes struct{}
