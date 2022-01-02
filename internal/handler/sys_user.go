package handler

import (
	"context"
	"gf-RuoYi/apiv1"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service"

	"github.com/gogf/gf/v2/os/glog"
)

// 用户管理
var SysUser = handlerUser{}

type handlerUser struct{}

func (h *handlerUser) Login(ctx context.Context, req *apiv1.LoginDoReq) (res *apiv1.LoginDoRes, err error) {
	out, err := service.SysUser.Login(ctx, model.SysUserLoginInput{
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res = &apiv1.LoginDoRes{}
	res.OK = false
	glog.Debug(ctx, out)
	// if len(out) == 0 {
	// 	err = errors.New("账户或密码错误！")
	// 	return
	// }
	return
}
