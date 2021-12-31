package service

import (
	"context"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
)

// 在线用户管理
var SysUserOnline = sysUserOnline{}

type sysUserOnline struct{}

// 通过token查询用户是否在线
func (s *sysUserOnline) GetOnlineByToken(ctx context.Context, token string) (online *entity.SysUserOnline, err error) {
	r, err := dao.SysUserOnline.Ctx(ctx).Where("token", token).All()
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, err
	}
	return nil, err
}

// 获取所有在线用户列表
func (s *sysUserOnline) GetList(ctx context.Context) (enlineEntity []*entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{Duration: time.Hour, Name: "token"}).Scan(&enlineEntity)
	return
}
