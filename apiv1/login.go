package apiv1

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	UserName string `v:"required#用户名不能为空" dc:"用户名"`
	Password string `v:"required#密码不能为空" dc:"密码"`
}

type LoginDoRes struct {
	Token string `json:"token" dc:"token"`
}
