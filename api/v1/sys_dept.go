package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取部门列表请求
type SysDeptListReq struct {
	g.Meta   `path:"/dept/getList" method:"post" summary:"获取部门列表" tags:"部门"`
	DeptId   int64  `v:"length:1,10#部门名长度为:{min}到{max}位" dc:"部门id"` // 部门id
	DeptName string `v:"max-length:60#部门名称长度最多为:{max}位" dc:"部门名称"`  // 部门名称
	Status   string `dc:"部门状态（0正常 1停用）" `                           // 部门状态（0正常 1停用）
}

// 获取部门列表响应
type SysDeptListRes []map[string]interface{}

// 单个部门信息请求
type SysDeptOneReq struct {
	g.Meta `path:"/dept/getOne" method:"post" summary:"获取单个部门信息" tags:"部门"`
	DeptId int64 `v:"required|length:1,10#部门id不能为空！|部门名长度为:{min}到{max}位" dc:"部门id"` // 部门id
}

// 单个部门信息响应
type SysDeptOneRes struct {
	*entity.SysDept
}

// 新增部门请求
type SysDeptCreateReq struct {
	g.Meta   `path:"/dept/create" method:"post" summary:"新增部门" tags:"部门"`
	ParentId int64  `v:"required|length:1,10#父部门id不能为空！|父部门id长度为:{min}到{max}位" dc:"父部门id"` // 父部门id
	DeptName string `v:"required|length:1,60#部门名称不能为空！|部门名称长度为:{min}到{max}位" dc:"部门名称"`    // 部门名称
	OrderNum int    `v:"required|length:1,3#显示顺序不能为空！|显示顺序长度为:{min}到{max}位" dc:"排序标记"`     // 显示顺序
	Leader   string `dc:"负责人" `                                                            // 负责人
	Phone    string `v:"phone#手机号码错误！" dc:"联系电话" `                                         // 手机号码
	Email    string `v:"email#邮箱格式错误！" dc:"邮箱" `                                           // 邮箱
	Status   string `dc:"部门状态（0正常 1停用）" `                                                  // 部门状态（0正常 1停用）
}

// 新增部门响应
type SysDeptCreateRes struct{}

// 更新部门信息请求
type SysDeptUpdateReq struct {
	g.Meta   `path:"/dept/update" method:"post" summary:"修改部门" tags:"部门"`
	DeptId   int64  `v:"required|length:1,10#部门id不能为空！|部门名长度为:{min}到{max}位" dc:"部门id"`     // 部门id
	ParentId int64  `v:"required|length:1,10#父部门id不能为空！|父部门id长度为:{min}到{max}位" dc:"父部门id"` // 父部门id
	DeptName string `v:"required|length:1,60#部门名称不能为空！|部门名称长度为:{min}到{max}位" dc:"部门名称"`    // 部门名称
	OrderNum int    `v:"required|length:1,3#显示顺序不能为空！|显示顺序长度为:{min}到{max}位" dc:"排序标记"`     // 显示顺序
	Leader   string `dc:"负责人" `                                                            // 负责人
	Phone    string `v:"phone#手机号码错误！" dc:"联系电话" `                                         // 手机号码
	Email    string `v:"email#邮箱格式错误！" dc:"邮箱" `                                           // 邮箱
	Status   string `dc:"部门状态（0正常 1停用）" `                                                  // 部门状态（0正常 1停用）
}

// 更新部门信息响应
type SysDeptUpdateRes struct{}

// 删除部门请求
type SysDeptDeleteReq struct {
	g.Meta     `path:"/dept/delete" method:"post" summary:"删除部门" tags:"部门"`
	DeptIdList []int `v:"required#部门id不能为空！" dc:"部门id"` // 需要删除的数据主键，例：[1,2,3]
}

// 删除部门响应
type SysDeptDeleteRes struct{}

// 查询部门树结构输入
type SysDeptTreeselectReq struct {
	g.Meta `path:"/dept/treeselect" method:"get" summary:"查询部门树结构" tags:"部门"`
}

// 查询部门树结构响应
type SysDeptTreeselectRes struct {
	Depts []map[string]interface{} `json:"depts"`
}
