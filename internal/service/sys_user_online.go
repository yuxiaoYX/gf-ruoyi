package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
)

// 在线用户管理
var SysUserOnline = sysUserOnline{}

type sysUserOnline struct{}

// 获取所有在线用户列表
func (s *sysUserOnline) GetList(ctx context.Context) (enlinesEntity []*entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Scan(&enlinesEntity)
	return
}

// 判断token是否在线，返回在线状态信息
func (s *sysUserOnline) GetToken(ctx context.Context, token string) (onlineEntity *entity.SysUserOnline, err error) {
	dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 12,
		Name:     "userToken-" + token,
	}).Where("token", token).Scan(&onlineEntity)
	return
}

// 添加在线用户
func (s *sysUserOnline) Create(ctx context.Context, in model.SysUserOnlineCreateInput) (err error) {
	// 单点登录，多点登录注释掉
	var onlineEntity []*entity.SysUserOnline
	if err = dao.SysUserOnline.Ctx(ctx).Where("user_name", in.UserName).Scan(&onlineEntity); err != nil {
		return err
	}
	for _, v := range onlineEntity {
		if err = s.Delete(ctx, model.SysUserOnlineDeleteInput{Token: v.Token}); err != nil {
			return err
		}
	}

	// 添加在线记录
	_, err = dao.SysUserOnline.Ctx(ctx).Data(in).Insert()
	return
}

// 删除指定在线用户,username
func (s *sysUserOnline) DeleteUserName(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1, Name: "userToken-" + in.Token,
	}).Where(in).Delete()

	return
}

// 删除指定在线用户,token
func (s *sysUserOnline) DeleteToken(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1, Name: "userToken-" + in.Token,
	}).Where(in).Delete()

	return
}
