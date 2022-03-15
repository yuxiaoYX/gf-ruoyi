package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"
	"gf-ruoyi/utility/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mssola/user_agent"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (c *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	// service用户名密码验证,并返回token
	out, err := service.SysUser().Login(ctx, model.SysUserLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 保存用户在线状态到数据库
	token := guid.S([]byte(out.UserName))
	r := g.RequestFromCtx(ctx)
	userAgent := r.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	explorer, _ := ua.Browser()
	service.SysUserOnline().Create(ctx, model.SysUserOnlineCreateInput{
		Token:    token,
		UserId:   int(out.UserId),
		UserName: out.UserName,
		Ip:       utils.GetClientIp(r),
		Os:       ua.OS(),
		Explorer: explorer,
	})

	res.Token = token
	return
}
