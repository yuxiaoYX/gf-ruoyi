package service

import (
	"context"
	"errors"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
func (s *sUser) GetList(ctx context.Context, in model.SysUserListInput) (out model.SysUserListOutput, err error) {
	m := dao.SysUser.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": in.UserName,
		"nick_name": in.NickName,
		"status":    in.Status,
		"dept_id":   in.DeptId,
	})
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total = len(out.Rows)
	return
}

// 获取用户详细信息,缓存10小时
func (s *sUser) GetOne(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {
	err = dao.SysUser.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "userid-" + gconv.String(in.UserId),
		Force:    false,
	}).Where("user_id", in.UserId).Scan(&out)
	return
}

// 新增用户
func (s *sUser) Create(ctx context.Context, in model.SysUserCreateInput) (err error) {
	userCount, err := dao.SysUser.Ctx(ctx).Where("user_name=? OR nick_name=?", in.UserName, in.NickName).Count()
	if err != nil {
		return err
	}
	if userCount > 0 {
		return errors.New("账户或昵称已存在！")
	}
	_, err = dao.SysUser.Ctx(ctx).Insert(in)
	return
}

// 更新用户,并删除缓存
func (s *sUser) Update(ctx context.Context, in model.SysUserUpdateInput) (err error) {
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userid-" + gconv.String(in.UserId),
	}).Data(in).Where("user_id", in.UserId).Update()
	return
}

// 删除用户,并删除缓存
func (s *sUser) Delete(ctx context.Context, in model.SysUserDeleteInput) (err error) {
	userIdList := gstr.Split(in.UserIdStr, ",")
	dao.SysUser.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userid-" + gconv.String("123"),
	}).Delete("user_id", userIdList)
	return
}
