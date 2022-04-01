package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sUser struct{}

// 用户管理服务
func SysUser() *sUser {
	return &sUser{}
}

// 登录验证
func (s *sUser) Login(ctx context.Context, in model.SysUserLoginInput) (out *model.SysUserLoginOutput, err error) {
	if err = dao.SysUser.Ctx(ctx).Where("password", in.Password).Where(
		"nick_name=? or user_name=?", in.UserName, in.UserName,
	).Scan(&out); err != nil {
		return
	}
	if out == nil {
		err = gerror.New("用户名或密码错误！")
		return
	}
	return
}

// 获取用户列表
func (s *sUser) GetList(ctx context.Context, in model.SysUserListInput) (out *model.SysUserListOutput, err error) {
	m := dao.SysUser.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": r.UserName,
		"status":    r.Status,
		"role_id":   r.RoleId,
	})
	if r.BeginTime != "" && r.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", r.BeginTime, r.EndTime)
	}
	if err = m.Page(r.PageNum, r.PageSize).Scan(&userRes.Rows); err != nil {
		return
	}
	userRes.Total, err = m.Count()
	return
}

func Paginate(r *ghttp.Request) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		type Pagination struct {
			Page int
			Size int
		}
		var pagination Pagination
		_ = r.Parse(&pagination)
		switch {
		case pagination.Size > 100:
			pagination.Size = 100

		case pagination.Size <= 0:
			pagination.Size = 10
		}
		return m.Page(pagination.Page, pagination.Size)
	}
}

// 获取用户详细信息,缓存10小时
func (s *sUser) GetOne(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {
	err = dao.SysUser.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "user-" + string(in.UserId),
		Force:    false,
	}).Where("user_id", in.UserId).Scan(&out)
	return
}

// 获取用户详细信息
func (s *sUser) Create(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {

}

// 获取用户详细信息
func (s *sUser) Update(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {

}

// 获取用户详细信息
func (s *sUser) Delete(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {

}
