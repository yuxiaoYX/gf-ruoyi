package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
)

type sUserOnline struct{}

// 在线用户管理
func UserOnline() *sUserOnline {
	return &sUserOnline{}
}

// 添加在线用户
func (s *sUserOnline) Create(ctx context.Context, in model.SysUserOnlineCreateInput) (err error) {

	// 单点登录，多点登录注释掉
	// var onlineEntity []*entity.SysUserOnline
	// if err = dao.SysUserOnline.Ctx(ctx).Where("user_name", in.UserName).Scan(&onlineEntity); err != nil {
	// 	return err
	// }
	// for _, v := range onlineEntity {
	// 	if err = s.Delete(ctx, model.SysUserOnlineDeleteInput{Token: v.Token}); err != nil {
	// 		return err
	// 	}
	// }

	// 添加在线记录
	_, err = dao.SysUserOnline.Ctx(ctx).Data(in).Insert()
	return
}
