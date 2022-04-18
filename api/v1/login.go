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
	g.Meta `path:"/logout" method:"post" summary:"执行用户注销接口" tags:"登录"`
}

// 注销响应
type LogoutRes struct{}

// 登录后获取用户信息请求
type LoginUserInfoReq struct {
	g.Meta `path:"/getInfo" method:"post" summary:"登录后获取用户信息" tags:"登录"`
}

// 登录后获取用户信息响应
type LoginUserInfoRes struct {
	User        *SysUserOneRes `json:"user"`        // 用户信息
	Roles       []string       `json:"roles"`       // 角色权限字符列表
	Permissions []string       `json:"permissions"` // 菜单权限标识
}

// 登录后获取用户路由表请求
type LoginUserRouterReq struct {
	g.Meta `path:"/getRouters" method:"post" summary:"登录后获取用户路由表" tags:"登录"`
}

// 登录后获取用户路由表响应
type LoginUserRouterRes []map[string]interface{}
