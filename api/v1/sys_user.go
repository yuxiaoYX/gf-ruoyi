package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysUserLogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"执行用户注销接口" tags:"用户"`
}

type SysUserLogoutRes struct{}

type SysUserOneReq struct {
	g.Meta `path:"/user/getOne" method:"post" summary:"获取单个用户信息" tags:"用户"`
	UserId uint `json:"userId"   v:"required#用户id不能为空" dc:"用户id" ` // 用户ID
}
type SysUserOneRes struct {
	UserId   uint   `json:"userId"    ` // 用户ID
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
