package service

import (
	"context"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/model/entity"
	"gf-RuoYi/internal/service/internal/dao"
	"gf-RuoYi/utility/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
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
	for _, v := range entinesEntity {
		if token == v.Token {
			return v, err
		}
	}
	return nil, err
}

// 添加在线用户
func (s *sysUserOnline) Create(ctx context.Context, user_name string) (err error) {
	// 删除对应用户名的在线用户
	s.DeleteUserName(ctx, model.SysUserOnlineDeleteInput{UserName: user_name})

	token := guid.S([]byte(user_name))
	_, err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{Duration: -1, Name: "userToken"}).Data(g.Map{
		"token":     token,
		"user_name": user_name,
		"ip":        utils.GetClientIp,
		"explorer":  "Edge",
		"os":        "win11",
	}).Replace()
	return
}

// 删除指定用户名的在线用户
func (s *sysUserOnline) DeleteUserName(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).OmitEmpty().Where("id", in.Id).WhereOr("user_name", in.UserName).Delete()
	return
}
