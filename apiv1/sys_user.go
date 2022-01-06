package apiv1

import (
	"gf-RuoYi/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SysUserLogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"执行用户注销接口" tags:"个人"`
}
type SysUserLogoutRes struct {
}

type SysUserInfoReq struct {
	g.Meta `path:"/getInfo" method:"get" summary:"登录成功后获取用户信息" tags:"用户"`
}
type SysUserInfoRes struct {
	User *model.SysUserGetRoleOutput `json:"user" dc:"用户信息"`
}
