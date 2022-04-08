package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 获取用户列表请求
type SysUserListReq struct {
	g.Meta    `path:"/user/getList" method:"post" summary:"获取用户列表" tags:"用户"`
	UserName  string `dc:"用户账户"`                                      // 用户账号
	NickName  string `dc:"用户昵称"`                                      // 用户昵称
	Status    string `dc:"用户状态；0:正常,1:禁用"`                            // 用户状态；0:正常,1:禁用
	DeptId    string `dc:"部门id"`                                      // 部门id
	BeginTime string `dc:"开始时间"`                                      // 开始时间
	EndTime   string `dc:"结束时间"`                                      // 结束时间
	PageNum   int    `d:"1" v:"min:0#分页号码错误" dc:"分页号码，默认1"`           // 分页码
	PageSize  int    `d:"10" v:"max:100#分页数量最多是100条" dc:"分页数量，最大100"` // 分页数量
}

// 获取用户列表响应
type SysUserListRes struct {
	Rows  []SysUserOneRes `json:"rows1"` // 列表
	Total int             `json:"total"` // 数据总数
}

type SysUserOneReq struct {
	g.Meta `path:"/user/getOne" method:"post" summary:"获取单个用户信息" tags:"用户"`
	UserId uint `json:"userId" v:"required#用户id不能为空" dc:"用户id" ` // 用户ID
}
type SysUserOneRes struct {
	UserId   uint   `json:"userId"`     // 用户ID
	UserName string `json:"userName"  ` // 用户账号
	NickName string `json:"nickName"  ` // 用户昵称
	// Password  string      `json:"password"  `// 登录密码
	Mobile    string      `json:"mobile"    ` // 手机号码
	Avatar    string      `json:"avatar"    ` // 用户头像地址
	Status    string      `json:"status"    ` // 用户状态；0:禁用,1:正常
	DeptId    string      `json:"deptId"    ` // 部门id
	Remark    string      `json:"remark"    ` // 备注
	LoginIp   string      `json:"loginIp"   ` // 最后登录IP
	LoginDate *gtime.Time `json:"loginDate" ` // 最后登录时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}

// 新增用户请求
type SysUserCreateReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"新增用户" tags:"用户"`
	UserName string `v:"required|length:1,60#用户名不能为空！|用户名长度为:{min}到{max}位" dc:"用户账号"` // 用户账号
	NickName string `v:"required#用户昵称不能为空！" dc:"用户昵称"`                                // 用户昵称
	Password string `v:"required|length:6,30#请输入密码！|密码长度为:{min}到{max}位" dc:"登录密码"`    // 登录密码
	Mobile   string `v:"required|phone#请输入手机号！|手机号格式错误" dc:"手机号码"`                    // 手机号码
	Avatar   string `dc:"用户头像地址"`                                                     // 用户头像地址
	Status   string `d:"0" dc:"用户状态；0:正常,1:禁用"`                                       // 用户状态；0:正常,1:禁用
	DeptId   string `dc:"部门id"`                                                       // 部门id
	Remark   string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`                         // 备注
}

// 新增用户响应
type SysUserCreateRes struct{}

// 更新用户信息请求
type SysUserUpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"修改用户" tags:"用户"`
	UserId   uint   `v:"required|length:1,5#用户id不能为空！|用户名长度为:{min}到{max}位" dc:"用户id"` // 用户ID
	UserName string `v:"required|length:1,60#用户名不能为空！|用户名长度为:{min}到{max}位" dc:"用户账号"` // 用户账号
	NickName string `v:"required#用户昵称不能为空！" dc:"用户昵称"`                                // 用户昵称
	Password string `v:"required|length:6,30#请输入密码！|密码长度为:{min}到{max}位" dc:"登录密码"`    // 登录密码
	Mobile   string `v:"required|phone#请输入手机号！|手机号格式错误" dc:"手机号码"`                    // 手机号码
	Avatar   string `dc:"用户头像地址"`                                                     // 用户头像地址
	Status   string `d:"0" dc:"用户状态；0:正常,1:禁用"`                                       // 用户状态；0:正常,1:禁用
	DeptId   string `dc:"部门id"`                                                       // 部门id
	Remark   string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`                         // 备注
}

// 更新用户信息响应
type SysUserUpdateRes struct{}

// 删除用户请求
type SysUserDeleteReq struct {
	g.Meta    `path:"/user/delete" method:"post" summary:"删除用户" tags:"用户"`
	UserIdStr string `v:"required#用户id不能为空！" dc:"用户id"` // 需要删除的数据主键，例：1,2,3
}

// 删除用户响应
type SysUserDeleteRes struct{}
