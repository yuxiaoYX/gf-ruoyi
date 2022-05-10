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
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUser struct{}

// 用户管理服务
func SysUser() *sUser {
	return &sUser{}
}

// 登录验证
func (s *sUser) Login(ctx context.Context, in model.SysUserLoginInput) (out *model.SysUserLoginOutput, err error) {
	if err = dao.SysUser.Ctx(ctx).Where("password=? AND status=0", in.Password).Where(
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
		"status":  in.Status,
		"dept_id": in.DeptId,
	})
	if in.UserName != "" {
		m = m.Where("user_name LIKE ?", "%"+in.UserName+"%")
	}
	if in.NickName != "" {
		m = m.Where("nick_name LIKE ?", "%"+in.NickName+"%")
	}
	if in.BeginTime != "" && in.EndTime != "" {
		m = m.Where("created_at>? and created_at<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total, err = m.Count()
	return
}

// 获取用户详细信息,缓存10小时
func (s *sUser) GetOne(ctx context.Context, in model.SysUserOneInput) (out *model.SysUserOneOutput, err error) {
	err = dao.SysUser.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     "userId-" + gconv.String(in.UserId),
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
	lastInsertId, err := dao.SysUser.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return err
	}
	// 添加用户和角色关联信息
	SysUserRole().UpdateUser(ctx, model.SysUserRoleUpdateUInput{UserId: uint(lastInsertId), Roleids: in.RoleIds})
	return
}

// 更新用户,并删除缓存
func (s *sUser) Update(ctx context.Context, in model.SysUserUpdateInput) (err error) {
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userId-" + gconv.String(in.UserId),
	}).Data(in).Where("user_id", in.UserId).Update()
	// 更新用户和角色关联信息
	SysUserRole().UpdateUser(ctx, model.SysUserRoleUpdateUInput{UserId: in.UserId, Roleids: in.RoleIds})
	return
}

// 删除用户,并删除缓存
func (s *sUser) Delete(ctx context.Context, in model.SysUserDeleteInput) (err error) {
	// userIdList := gstr.Split(in.UserIdStr, ",")
	for _, v := range in.UserIdList {
		if _, err = dao.SysUser.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "userId-" + string(rune(v)),
		}).Delete("user_id=?", v); err != nil {
			return
		}
	}
	// 删除用户和角色的管理信息
	if err = SysUserRole().Delete(ctx, model.SysUserRoleDeleteInput{UserIdList: in.UserIdList}); err != nil {
		return
	}
	// 删除用户token
	SysUserOnline().Delete(ctx, model.SysUserOnlineDeleteInput{IdList: gconv.Uint64s(in.UserIdList)})
	return
}

// 用户密码重置,只是修改密码，无需删除缓存
func (s *sUser) ResetPwd(ctx context.Context, in model.SysUserResetPwdInput) (err error) {
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Data(in).Where("user_id", in.UserId).Update()
	return
}

// 用户状态修改，并删除缓存
func (s *sUser) ChangeStatus(ctx context.Context, in model.SysUserChangeStatusInput) (err error) {
	if _, err = dao.SysUser.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userId-" + gconv.String(in.UserId),
	}).Data(in).Where("user_id", in.UserId).Update(); err != nil {
		return
	}

	// 删除用户token
	SysUserOnline().Delete(ctx, model.SysUserOnlineDeleteInput{IdList: gconv.Uint64s(in.UserId)})
	return
}

// 用户修改个人信息
func (s *sUser) UpdateProfile(ctx context.Context, in model.SysUserUpdateProfileInput) (err error) {
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userId-" + gconv.String(in.UserId),
	}).Data(in).Where("user_id", in.UserId).Update()
	return
}

// 用户修改个人密码,只是修改密码，无需删除缓存
func (s *sUser) UpdatePwd(ctx context.Context, in model.SysUserUpdatePwdInput) (err error) {
	r, err := dao.SysUser.Ctx(ctx).Where("user_id=? AND password=?", in.UserId, in.OldPassword).One()
	if err != nil {
		return err
	}
	if len(r) == 0 {
		return gerror.New("旧密码错误！")
	}
	_, err = dao.SysUser.Ctx(ctx).Data("password", in.NewPassword).Where("user_id", in.UserId).Update()
	return
}

// 用户修改头像
func (s *sUser) UpdateAvatar(ctx context.Context, in model.SysUserUpdateAvatarInput) (err error) {
	// 删除历史头像文件
	userEntity, err := s.GetOne(ctx, model.SysUserOneInput{UserId: in.UserId})
	if err != nil {
		return err
	}
	gfile.Remove(userEntity.Avatar)
	// 保存头像链接
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "userId-" + gconv.String(in.UserId),
	}).Data(in).Where("user_id", in.UserId).Update()
	return
}
