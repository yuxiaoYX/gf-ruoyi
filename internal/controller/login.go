package controller

import (
	"context"
	"gf-ruoyi/apiv1"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (c *cLogin) Login(ctx context.Context, req *apiv1.LoginDoReq) (res *apiv1.LoginDoRes, err error) {
	// service用户名密码验证,并返回token

	// 保存用户在线状态到数据库

	res = &apiv1.LoginDoRes{}
	res.Token = "1111111111"
	return
}
