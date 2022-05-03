package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取字典类型列表请求
type SysDictTypeListReq struct {
	g.Meta    `path:"/dictType/getList" method:"post" summary:"获取字典类型列表" tags:"字典类型"`
	DictName  string `v:"max-length:100#字典名称最大长度为{max}位" dc:"字典名称"`   // 字典名称
	DictType  string `v:"max-length:100#字典类型最大长度为{max}位" dc:"字典类型"`   // 字典类型
	Status    string `dc:"字典类型状态；0:正常,1:禁用"`                          // 状态（0正常 1停用）
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取字典类型列表响应
type SysDictTypeListRes struct {
	Rows  []*SysDictTypeOneRes `json:"rows"`  // 列表
	Total int                  `json:"total"` // 数据总数
}

// 单个字典类型信息请求
type SysDictTypeOneReq struct {
	g.Meta `path:"/dictType/getOne" method:"post" summary:"获取单个字典类型信息" tags:"字典类型"`
	DictId uint `v:"required#字典主键不能为空" dc:"字典主键" ` // 字典主键ID
}

// 单个字典类型信息响应
type SysDictTypeOneRes struct {
	*entity.SysDictType
}

// 新增字典类型请求
type SysDictTypeCreateReq struct {
	g.Meta   `path:"/dictType/create" method:"post" summary:"新增字典类型" tags:"字典类型"`
	DictName string `v:"required|length:1,100#字典名称不能为空！|字典名称长度为:{min}到{max}位" dc:"字典名称"` // 字典名称
	DictType string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
	Status   string `d:"0" dc:"字典类型状态；0:正常,1:禁用"`                                        // 状态（0正常 1停用）
	Remark   string `v:"max-length:500#备注最多为{max}个字符！" dc:"备注"`                          // 备注
}

// 新增字典类型响应
type SysDictTypeCreateRes struct{}

// 更新字典类型信息请求
type SysDictTypeUpdateReq struct {
	g.Meta   `path:"/dictType/update" method:"post" summary:"修改字典类型" tags:"字典类型"`
	DictId   int64  `v:"required|length:1,10#字典主键不能为空！|字典主键长度为:{min}到{max}位" dc:"字典主键"`  // 字典主键
	DictName string `v:"required|length:1,100#字典名称不能为空！|字典名称长度为:{min}到{max}位" dc:"字典名称"` // 字典名称
	DictType string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
	Status   string `d:"0" dc:"字典类型状态；0:正常,1:禁用"`                                        // 状态（0正常 1停用）
	Remark   string `v:"max-length:500#备注最多为{max}个字符！" dc:"备注"`                          // 备注
}

// 更新字典类型信息响应
type SysDictTypeUpdateRes struct{}

// 删除字典类型请求
type SysDictTypeDeleteReq struct {
	g.Meta     `path:"/dictType/delete" method:"post" summary:"删除字典类型" tags:"字典类型"`
	DictIdList []int64 `v:"required#字典主键不能为空！" dc:"字典主键"` // 需要删除的数据主键，例：[1,2,3]
}

// 删除字典类型响应
type SysDictTypeDeleteRes struct{}
