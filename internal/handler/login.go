package handler

import (
	"context"
	"gf-RuoYi/apiv1"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service"
	"gf-RuoYi/utility/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mssola/user_agent"
)

// 登录管理
var Login = handlerLogin{}

type handlerLogin struct{}

// 用户登录验证
func (h *handlerLogin) Login(ctx context.Context, req *apiv1.LoginDoReq) (res *apiv1.LoginDoRes, err error) {
	out, err := service.SysUser.Login(ctx, model.SysUserLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	//保存用户在线状态token到数据库
	token := guid.S([]byte(out.UserName))
	r := g.RequestFromCtx(ctx)
	userAgent := r.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	explorer, _ := ua.Browser()
	service.SysUserOnline.Create(ctx, model.SysUserOnlineCreateInput{
		Token:    token,
		UserName: out.UserName,
		Ip:       utils.GetClientIp(r),
		Os:       ua.OS(),
		Explorer: explorer,
	})
	res = &apiv1.LoginDoRes{}
	res.Token = token
	return
}
