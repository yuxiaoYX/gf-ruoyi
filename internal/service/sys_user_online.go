package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
)

// 在线用户管理
var SysUserOnline = sysUserOnline{}

type sysUserOnline struct{}

// 获取所有在线用户列表,缓存
func (s *sysUserOnline) GetList(ctx context.Context) (enlinesEntity []*entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{Duration: 0, Name: "userToken"}).Scan(&enlinesEntity)
	return
}

// 判断token是否在线，返回在线状态信息
func (s *sysUserOnline) GetToken(ctx context.Context, token string) (enlineEntity *entity.SysUserOnline, err error) {
	entinesEntity, err := s.GetList(ctx)
	if err != nil {
		return nil, err
	}
	glog.Debug(ctx, entinesEntity)
	for _, v := range entinesEntity {
		if token == v.Token {
			return v, err
		}
	}
	return nil, err
}

// 添加在线用户
func (s *sysUserOnline) Create(ctx context.Context, in model.SysUserOnlineCreateInput) (err error) {
	// 删除对应用户名的在线用户
	s.Delete(ctx, model.SysUserOnlineDeleteInput{UserName: in.UserName})

	_, err = dao.SysUserOnline.Ctx(ctx).Data(in).Insert()
	return
}

// 删除指定在线用户,id OR user_name
func (s *sysUserOnline) Delete(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1, Name: "userToken",
	}).Where("id", in.Id).WhereOr("user_name", in.UserName).Delete()
	return
}
