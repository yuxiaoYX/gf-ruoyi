package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取字典类型列表输入
type SysDictTypeListInput struct {
	DictName  string // 字典名称
	DictType  string // 字典类型
	Status    string // 状态（0正常 1停用）
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取字典类型列表输出
type SysDictTypeListOutput struct {
	Rows  []*SysDictTypeOneOutput // 列表
	Total int                     // 数据总数
}

// 获取单个字典类型信息输入
type SysDictTypeOneInput struct {
	DictId uint // 字典类型ID
}

// 获取单个字典类型信息输出
type SysDictTypeOneOutput struct {
	*entity.SysDictType
}

// 新增字典类型输入
type SysDictTypeCreateInput struct {
	DictName string // 字典名称
	DictType string // 字典类型
	Status   string // 状态（0正常 1停用）
	Remark   string // 备注
}

// 更新字典类型信息输入
type SysDictTypeUpdateInput struct {
	DictId   int64  // 字典主键
	DictName string // 字典名称
	DictType string // 字典类型
	Status   string // 状态（0正常 1停用）
	Remark   string // 备注
}

// 删除字典类型输入
type SysDictTypeDeleteInput struct {
	DictIdStr string // 需要删除的数据主键，例：1,2,3
}
