package v1

import (
	"gf-ruoyi/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取角色列表请求
type SysRoleListReq struct {
	g.Meta    `path:"/role/getList" method:"post" summary:"获取角色列表" tags:"角色"`
	RoleName  string `v:"max-length:60#角色名称长度最多为:{max}位" dc:"角色名称"`   // 角色名称
	Status    string `d:"0" dc:"角色状态；0:正常,1:禁用"`                      // 角色状态；0:正常，1:禁用
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取角色列表响应
type SysRoleListRes struct {
	Rows  []*SysRoleOneRes `json:"rows"`  // 列表
	Total int              `json:"total"` // 数据总数
}

// 单个角色信息请求
type SysRoleOneReq struct {
	g.Meta `path:"/role/getOne" method:"post" summary:"获取单个角色信息" tags:"角色"`
	RoleId uint `json:"roleId" v:"required#角色id不能为空" dc:"角色id" ` // 角色ID
}

// 单个角色信息响应
type SysRoleOneRes struct {
	*entity.SysRole
}

// 新增角色请求
type SysRoleCreateReq struct {
	g.Meta   `path:"/role/create" method:"post" summary:"新增角色" tags:"角色"`
	RoleName string `v:"required|length:1,60#角色名称不能为空！|角色名称长度为:{min}到{max}位" dc:"角色名称"` // 角色名称
	RoleSort int    `d:"显示顺序" `                                                         // 显示顺序
	Status   string `d:"0" dc:"角色状态；0:正常,1:禁用"`                                         // 角色状态；0:正常，1:禁用
	Remark   string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`                           // 备注
}

// 新增角色响应
type SysRoleCreateRes struct{}

// 更新角色信息请求
type SysRoleUpdateReq struct {
	g.Meta   `path:"/role/update" method:"post" summary:"修改角色" tags:"角色"`
	RoleId   uint   `v:"required|length:1,5#角色id不能为空！|角色名长度为:{min}到{max}位" dc:"角色id"`   // 角色ID
	RoleName string `v:"required|length:1,60#角色名称不能为空！|角色名称长度为:{min}到{max}位" dc:"角色名称"` // 角色名称
	RoleSort int    `d:"显示顺序" `                                                         // 显示顺序
	Status   string `d:"0" dc:"角色状态；0:正常,1:禁用"`                                         // 角色状态；0:正常，1:禁用
	Remark   string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`                           // 备注
}

// 更新角色信息响应
type SysRoleUpdateRes struct{}

// 删除角色请求
type SysRoleDeleteReq struct {
	g.Meta    `path:"/role/delete" method:"post" summary:"删除角色" tags:"角色"`
	RoleIdStr string `v:"required#角色id不能为空！" dc:"角色id"` // 需要删除的数据主键，例：1,2,3
}

// 删除角色响应
type SysRoleDeleteRes struct{}
