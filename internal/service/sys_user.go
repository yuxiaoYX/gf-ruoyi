package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// 用户管理服务
var SysUser = serviceUser{}

type serviceUser struct{}

// 执行登录
func (s *serviceUser) Login(ctx context.Context, in model.SysUserLoginInput) (out *model.SysUserLoginOutput, err error) {
	var userEntity []*entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		"user_name": in.UserName,
		"nick_name": in.NickName,
	}).Where("password", in.Password).OmitEmpty().Scan(userEntity)
	if len(userEntity) == 0 {

	}
	return
}
