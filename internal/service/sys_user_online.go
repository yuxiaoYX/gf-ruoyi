package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/model/entity"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sUserOnline struct{}

// 在线用户管理
func SysUserOnline() *sUserOnline {
	return &sUserOnline{}
}

// 获取在线用户列表
func (s *sUserOnline) GetList(ctx context.Context, in model.SysUserOnlineListInput) (out model.SysUserOnlineListOutput, err error) {
	m := dao.SysUserOnline.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": in.UserName,
	})
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
	return
}

// 添加在线用户
func (s *sUserOnline) Create(ctx context.Context, in model.SysUserOnlineCreateInput) (err error) {
	// 删除该用户id的在线记录，多点登录注释掉
	// if _, err = dao.SysUserOnline.Ctx(ctx).OmitEmpty().Where("user_id",in.UserId).Delete(); err != nil {
	// 	return err
	// }

	// 添加在线记录
	_, err = dao.SysUserOnline.Ctx(ctx).Data(in).Insert()
	return
}

// 删除在线用户
func (s *sUserOnline) Delete(ctx context.Context, in model.SysUserOnlineDeleteInput) (err error) {
	var onlineEntity []*entity.SysUserOnline
	err = dao.SysUserOnline.Ctx(ctx).Where("id IN(?)", in.Ids).Scan(&onlineEntity)
	if err != nil {
		return err
	}
	for _, v := range onlineEntity {
		if _, err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "token-" + v.Token,
			Force:    false,
		}).Delete("id", v.Id); err != nil {
			return
		}
	}
	return
}

// 使用token删除在线用户
func (s *sUserOnline) TokenDelete(ctx context.Context, token string) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "token-" + token,
		Force:    false,
	}).Delete("token=?", token)
	return
}

// 判断token是否在线，返回在线状态信息
// TODO 查询方式可以改成缓存模式
// TODO 可以使用用户名或用户id查询数据库是否存在token
func (s *sUserOnline) GetToken(ctx context.Context, token string) (out *model.SysUserOnlineGetTokenOutput, err error) {
	dao.SysUserOnline.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "token-" + token,
		Force:    false,
	}).Where("token", token).Scan(&out)
	return
}
