package v1

import "github.com/gogf/gf/v2/frame/g"

type UserLogoutReq struct {
	g.Meta `path:"/user/logout" method:"get" summary:"执行用户注销接口" tags:"个人"`
}

type UserLogoutRes struct{}
