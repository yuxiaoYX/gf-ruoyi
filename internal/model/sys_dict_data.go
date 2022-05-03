package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取字典数据列表输入
type SysDictDataListInput struct {
	DictLabel string // 字典标签
	DictType  string // 字典类型
	Status    string // 状态（0正常 1停用）
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取字典数据列表输出
type SysDictDataListOutput struct {
	Rows  []*SysDictDataOneOutput // 列表
	Total int                     // 数据总数
}

// 获取单个字典数据信息输入
type SysDictDataOneInput struct {
	DictCode int64 // 字典编码
}

// 获取单个字典数据信息输出
type SysDictDataOneOutput struct {
	*entity.SysDictData
}

// 新增字典数据输入
type SysDictDataCreateInput struct {
	DictSort  int    // 字典排序
	DictLabel string // 字典标签
	DictValue string // 字典键值
	DictType  string // 字典类型
	CssClass  string // 样式属性（其他样式扩展）
	ListClass string // 表格回显样式
	IsDefault string // 是否默认（0是 1否）
	Status    string // 状态（0正常 1停用）
	Remark    string // 备注
}

// 更新字典数据信息输入
type SysDictDataUpdateInput struct {
	DictCode  int64  // 字典编码
	DictSort  int    // 字典排序
	DictLabel string // 字典标签
	DictValue string // 字典键值
	DictType  string // 字典类型
	CssClass  string // 样式属性（其他样式扩展）
	ListClass string // 表格回显样式
	IsDefault string // 是否默认（0是 1否）
	Status    string // 状态（0正常 1停用）
	Remark    string // 备注
}

// 删除字典数据输入
type SysDictDataDeleteInput struct {
	DictCodeList []int64 // 需要删除的数据主键，例：[1,2,3]
}
