package v1

import "github.com/gogf/gf/v2/frame/g"

// 登录请求
type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	UserName string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码(明文)"`
}

// 登录响应
type LoginDoRes struct {
	Token string `json:"token" dc:"用户token"`
}

// 注销请求
type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"执行用户注销接口" tags:"用户"`
}

// 注销响应
type LogoutRes struct{}
