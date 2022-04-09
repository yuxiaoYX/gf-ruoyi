package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 字典类型查询字典数据请求
type SysDictDataGetTypeReq struct {
	g.Meta   `path:"/dictData/type" method:"get" summary:"获取字典数据列表" tags:"字典数据"`
	DictType string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
}

// 字典类型查询字典数据响应
type SysDictDataGetTypeRes []map[string]interface{}

// 获取字典数据列表请求
type SysDictDataListReq struct {
	g.Meta    `path:"/dictData/getList" method:"post" summary:"获取字典数据列表" tags:"字典数据"`
	DictLabel string `v:"max-length:100#字典标签最大长度为{max}位" dc:"字典标签"`                       // 字典标签
	DictType  string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
	Status    string `dc:"字典数据状态；0:正常,1:停用"`                                              // 状态（0正常 1停用）
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`                               // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"`                     // 分页数量
}

// 获取字典数据列表响应
type SysDictDataListRes struct {
	Rows  []*SysDictDataOneRes `json:"rows"`  // 列表
	Total int                  `json:"total"` // 数据总数
}

// 单个字典数据信息请求
type SysDictDataOneReq struct {
	g.Meta   `path:"/dictData/getOne" method:"post" summary:"获取单个字典数据信息" tags:"字典数据"`
	DictCode uint `v:"required#字典编码不能为空" dc:"字典编码" ` // 字典编码
}

// 单个字典数据信息响应
type SysDictDataOneRes struct {
	*entity.SysDictData
}

// 新增字典数据请求
type SysDictDataCreateReq struct {
	g.Meta    `path:"/dictData/create" method:"post" summary:"新增字典数据" tags:"字典数据"`
	DictCode  int64  `v:"required|length:1,10#字典编码不能为空！|字典编码长度为:{min}到{max}位" dc:"字典编码"`  // 字典编码
	DictSort  int    `v:"required|length:1,4#字典排序不能为空！|字典排序长度为:{min}到{max}位" dc:"字典排序"`   // 字典排序
	DictLabel string `v:"required|length:1,100#字典标签不能为空！|字典标签长度为:{min}到{max}位" dc:"字典标签"` // 字典标签
	DictValue string `v:"required|length:1,100#字典键值不能为空！|字典键值长度为:{min}到{max}位" dc:"字典键值"` // 字典键值
	DictType  string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
	CssClass  string ` dc:"样式属性（其他样式扩展）"`                                                 // 样式属性（其他样式扩展）
	ListClass string ` dc:"表格回显样式"`                                                       // 表格回显样式
	IsDefault string ` dc:"是否默认（0是 1否） "`                                                 // 是否默认（0是 1否）
	Status    string `d:"0" dc:"字典数据状态；0:正常,1:停用"`                                        // 状态（0正常 1停用）
	Remark    string `v:"max-length:500#备注最多为{max}个字符！" dc:"备注"`                          // 备注
}

// 新增字典数据响应
type SysDictDataCreateRes struct{}

// 更新字典数据信息请求
type SysDictDataUpdateReq struct {
	g.Meta    `path:"/dictData/update" method:"post" summary:"修改字典数据" tags:"字典数据"`
	DictCode  int64  `v:"required|length:1,10#字典编码不能为空！|字典编码长度为:{min}到{max}位" dc:"字典编码"`  // 字典编码
	DictSort  int    `v:"required|length:1,4#字典排序不能为空！|字典排序长度为:{min}到{max}位" dc:"字典排序"`   // 字典排序
	DictLabel string `v:"required|length:1,100#字典标签不能为空！|字典标签长度为:{min}到{max}位" dc:"字典标签"` // 字典标签
	DictValue string `v:"required|length:1,100#字典键值不能为空！|字典键值长度为:{min}到{max}位" dc:"字典键值"` // 字典键值
	DictType  string `v:"required|length:1,100#字典类型不能为空！|字典类型长度为:{min}到{max}位" dc:"字典类型"` // 字典类型
	CssClass  string ` dc:"样式属性（其他样式扩展）"`                                                 // 样式属性（其他样式扩展）
	ListClass string ` dc:"表格回显样式"`                                                       // 表格回显样式
	IsDefault string ` dc:"是否默认（0是 1否） "`                                                 // 是否默认（0是 1否）
	Status    string `d:"0" dc:"字典数据状态；0:正常,1:停用"`                                        // 状态（0正常 1停用）
	Remark    string `v:"max-length:500#备注最多为{max}个字符！" dc:"备注"`                          // 备注
}

// 更新字典数据信息响应
type SysDictDataUpdateRes struct{}

// 删除字典数据请求
type SysDictDataDeleteReq struct {
	g.Meta      `path:"/dictData/delete" method:"post" summary:"删除字典数据" tags:"字典数据"`
	DictCodeStr string `v:"required#字典编码不能为空！" dc:"字典编码"` // 需要删除的数据主键，例：1,2,3
}

// 删除字典数据响应
type SysDictDataDeleteRes struct{}
