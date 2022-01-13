package handler

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
)

// GToken登录验证管理
var Auth = hAuth{
	multiLogin: g.Cfg().GetAdapter(),
}

type hAuth struct {
	multiLogin bool
	gfToken    *gtoken.GfToken
}
