package service

import (
	"context"
	"errors"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/glog"
)

// 用户管理服务
var SysUser = serviceUser{}

type serviceUser struct{}

// 登录验证
func (s *serviceUser) Login(ctx context.Context, in model.SysUserLoginInput) (out *model.SysUserLoginOutput, err error) {
	if err = dao.SysUser.Ctx(ctx).Where("password", in.Password).Where(
		"nick_name=? or user_name=?", in.UserName, in.UserName,
	).Scan(&out); err != nil {
		return
	}
	if out == nil {
		err = errors.New("用户名或密码错误！")
		return
	}
	return
}

// 获取用户详细信息
func (s *serviceUser) GetById(ctx context.Context, in model.SysUserGetNameInput) (out *model.SysUserGetRoleOutput, err error) {
	err = dao.SysUser.Ctx(ctx).Where("user_name", in.UserName).Scan(&out)
	glog.Debug(ctx, out)
	return
}
