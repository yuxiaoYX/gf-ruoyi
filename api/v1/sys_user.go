package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
	Rows  []SysUserOneRes `json:"rows"`  // 列表
	Total int             `json:"total"` // 数据总数
}

// 获取单个用户信息输入
type SysUserOneReq struct {
	g.Meta `path:"/user/getOne" method:"post" summary:"获取单个用户信息" tags:"用户"`
	UserId uint `json:"userId" v:"required#用户id不能为空" dc:"用户id" ` // 用户ID
}

// 获取单个用户信息响应
type SysUserOneRes struct {
	UserId   uint   `json:"userId"`     // 用户ID
	UserName string `json:"userName"  ` // 用户账号
	NickName string `json:"nickName"  ` // 用户昵称
	// Password  string      `json:"password"  `// 登录密码
	Mobile    string           `json:"mobile"    ` // 手机号码
	Avatar    string           `json:"avatar"    ` // 用户头像地址
	Status    string           `json:"status"    ` // 用户状态；0:禁用,1:正常
	DeptId    int64            `json:"deptId"    ` // 部门id
	Remark    string           `json:"remark"    ` // 备注
	LoginIp   string           `json:"loginIp"   ` // 最后登录IP
	LoginDate *gtime.Time      `json:"loginDate" ` // 最后登录时间
	CreatedAt *gtime.Time      `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time      `json:"updatedAt" ` // 更新时间
	DeletedAt *gtime.Time      `json:"deletedAt" ` // 删除时间
	Dept      *SysDeptOneRes   `json:"dept" `      // 部门信息
	RoleIds   []uint           `json:"roleIds" `   // 角色id列表
	Roles     []*SysRoleOneRes `json:"roles" `     // 角色信息列表
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
	DeptId   int64  `d:"100" dc:"部门id"`                                               // 部门id
	Remark   string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`                         // 备注
	RoleIds  []uint // 角色选中id列表
}

// 新增用户响应
type SysUserCreateRes struct{}

// 更新用户信息请求
type SysUserUpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"修改用户" tags:"用户"`
	UserId   uint   `v:"required|length:1,10#用户id不能为空！|用户名长度为:{min}到{max}位" dc:"用户id"` // 用户ID
	UserName string `v:"required|length:1,60#用户名不能为空！|用户名长度为:{min}到{max}位" dc:"用户账号"`  // 用户账号
	NickName string `v:"required#用户昵称不能为空！" dc:"用户昵称"`                                 // 用户昵称
	// Password string `v:"required|length:6,30#请输入密码！|密码长度为:{min}到{max}位" dc:"登录密码"`    // 登录密码
	Mobile  string `v:"required|phone#请输入手机号！|手机号格式错误" dc:"手机号码"` // 手机号码
	Avatar  string `dc:"用户头像地址"`                                  // 用户头像地址
	Status  string `d:"0" dc:"用户状态；0:正常,1:禁用"`                    // 用户状态；0:正常,1:禁用
	DeptId  int64  `dc:"部门id"`                                    // 部门id
	Remark  string `v:"max-length:200#备注最多为200个字符！" dc:"备注"`      // 备注
	RoleIds []uint // 角色选中id列表
}

// 更新用户信息响应
type SysUserUpdateRes struct{}

// 删除用户请求
type SysUserDeleteReq struct {
	g.Meta     `path:"/user/delete" method:"post" summary:"删除用户" tags:"用户"`
	UserIdList []uint `v:"required#用户id不能为空！" dc:"用户id"` // 需要删除的数据主键，例：[1,2,3]
}

// 删除用户响应
type SysUserDeleteRes struct{}

// 用户分配角色请求
type SysUserSelectRoleReq struct {
	g.Meta  `path:"/user/userSelectRole" method:"post" summary:"更新用户绑定的角色" tags:"用户"`
	UserId  uint   // 用户id
	RoleIds []uint // 角色选中id列表
}

// 用户分配角色响应
type SysUserSelectRoleRes struct{}

// 用户密码重置请求
type SysUserResetPwdReq struct {
	g.Meta   `path:"/user/resetPwd" method:"post" summary:"重置密码" tags:"用户"`
	UserId   uint   `v:"required|length:1,10#用户id不能为空！|用户名长度为:{min}到{max}位" dc:"用户id"` // 用户ID
	Password string `v:"required|length:6,30#请输入密码！|密码长度为:{min}到{max}位" dc:"登录密码"`     // 登录密码
}

// 用户密码重置响应
type SysUserResetPwdRes struct{}

// 用户状态修改请求
type SysUserChangeStatusReq struct {
	g.Meta `path:"/user/changeStatus" method:"post" summary:"用户状态修改" tags:"用户"`
	UserId uint   `v:"required|length:1,10#用户id不能为空！|用户名长度为:{min}到{max}位" dc:"用户id"`      // 用户ID
	Status string `v:"required|max-length:1#用户状态不能为空！|用户状态长度为:{max}" dc:"用户状态；0:正常,1:禁用"` // 用户状态；0:正常,1:禁用
}

// 用户状态修改响应
type SysUserChangeStatusRes struct{}

// 用户查询个人信息请求
type SysUserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" summary:"用户查询个人信息" tags:"用户"`
}

// 用户查询个人信息响应
type SysUserProfileRes struct {
	*SysUserOneRes
}

// 用户修改个人信息请求
type SysUserUpdateProfileReq struct {
	g.Meta   `path:"/user/updateProfile" method:"post" summary:"用户修改个人信息" tags:"用户"`
	NickName string `v:"required#用户昵称不能为空！" dc:"用户昵称"`             // 用户昵称
	Mobile   string `v:"required|phone#请输入手机号！|手机号格式错误" dc:"手机号码"` // 手机号码
}

// 用户修改个人信息响应
type SysUserUpdateProfileRes struct{}

// 用户修改个人密码请求
type SysUserUpdatePwdReq struct {
	g.Meta      `path:"/user/profile/updatePwd" method:"post" summary:"用户修改个人密码" tags:"用户"`
	OldPassword string `v:"required|length:6,30#请输入旧密码！|密码长度为:{min}到{max}位" dc:"旧密码"` // 旧密码
	NewPassword string `v:"required|length:6,30#请输入新密码！|密码长度为:{min}到{max}位" dc:"新密码"` // 新密码
}

// 用户修改个人密码响应
type SysUserUpdatePwdRes struct{}

// 用户头像上传请求
type SysUserUpdateAvatarReq struct {
	g.Meta     `path:"/user/profile/avatar" method:"post" summary:"修改个人头像" tags:"用户"`
	Avatarfile *ghttp.UploadFile `json:"avatarfile" type:"file" dc:"选择上传文件"`
}

// 用户头像上传响应
type SysUserUpdateAvatarRes struct {
	ImgUrl string `json:"imgUrl"` // 头像地址
}
