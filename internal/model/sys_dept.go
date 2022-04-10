package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取部门列表输入
type SysDeptListInput struct {
	DeptId   int64  // 部门id
	DeptName string // 部门名称
	Status   string // 部门状态；0:正常，1:禁用
}

// 获取单个部门信息输入
type SysDeptOneInput struct {
	DeptId int64 // 部门id
}

// 获取单个部门信息输出
type SysDeptOneOutput struct {
	*entity.SysDept
}

// 新增部门输入
type SysDeptCreateInput struct {
	ParentId int64  // 父部门id
	DeptName string // 部门名称
	OrderNum int    // 显示顺序
	Leader   string // 负责人
	Phone    string // 手机号码
	Email    string // 邮箱
	Status   string // 部门状态（0正常 1停用）
}

// 更新部门信息输入
type SysDeptUpdateInput struct {
	DeptId   int64  // 部门id
	ParentId int64  // 父部门id
	DeptName string // 部门名称
	OrderNum int    // 显示顺序
	Leader   string // 负责人
	Phone    string // 手机号码
	Email    string // 邮箱
	Status   string // 部门状态（0正常 1停用）
}

// 删除部门输入
type SysDeptDeleteInput struct {
	DeptIdStr string // 需要删除的数据主键，例：1,2,3
}
