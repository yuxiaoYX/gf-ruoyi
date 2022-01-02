package apiv1

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	UserName string `dc:"用户名"`
	NickName string `v:"required-without:UserName" dc:"用户昵称"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

type LoginDoRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}
