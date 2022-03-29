package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
)

type sUserOnline struct{}

// 在线用户管理
func SysUserOnline() *sUserOnline {
	return &sUserOnline{}
}

// 添加在线用户
func (s *sUserOnline) Create(ctx context.Context, in model.SysUserOnlineCreateInput) (err error) {
	// 删除该用户id的在线记录，多点登录注释掉
	if err = s.Delete(ctx, model.SysUserOnlineDeleteInput{UserId: in.UserId}); err != nil {
		return err
	}
	// 添加在线记录
	_, err = dao.SysUserOnline.Ctx(ctx).Data(in).Insert()
	return
}

// 删除在线用户
func (s *sUserOnline) Delete(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).OmitEmpty().Where(in).Delete()
	return
}

// 判断token是否在线，返回在线状态信息
// TODO 查询方式可以改成缓存模式
func (s *sUserOnline) GetToken(ctx context.Context, token string) (out *model.SysUserOnlineGetTokenOutput, err error) {
	dao.SysUserOnline.Ctx(ctx).Where("token", token).Scan(&out)
	return
}
