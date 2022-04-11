package controller

import (
	"context"
	v1 "gf-ruoyi/api/v1"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// 用户管理
var SysUser = cUser{}

type cUser struct{}

// 获取用户列表
func (c *cUser) GetList(ctx context.Context, req *v1.SysUserListReq) (res *v1.SysUserListRes, err error) {
	in := &model.SysUserListInput{}
	gconv.Scan(req, &in)

	userRes, err := service.SysUser().GetList(ctx, *in)
	gconv.Scan(userRes, &res)
	return
}

// 获取单个用户信息
func (c *cUser) GetOne(ctx context.Context, req *v1.SysUserOneReq) (res *v1.SysUserOneRes, err error) {
	// 用户信息
	userRes, err := service.SysUser().GetOne(ctx, model.SysUserOneInput{
		UserId: req.UserId,
	})
	if err != nil {
		return
	}
	gconv.Scan(userRes, &res)
	// 角色信息
	roleRes, err := service.SysUserRole().GetRoles(ctx, userRes.UserId)
	if err != nil {
		return
	}
	gconv.Scan(roleRes, &res.Roles)
	for _, v := range roleRes {
		res.RoleIds = append(res.RoleIds, v.RoleId)
	}
	// 部门信息
	deptRes, err := service.SysDept().GetOne(ctx, model.SysDeptOneInput{DeptId: userRes.DeptId})
	gconv.Scan(deptRes, &res.Dept)
	return
}

// 新建用户
func (c *cUser) Create(ctx context.Context, req *v1.SysUserCreateReq) (res *v1.SysUserCreateRes, err error) {
	in := &model.SysUserCreateInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Create(ctx, *in)
	return
}

// 更新用户
func (c *cUser) Update(ctx context.Context, req *v1.SysUserUpdateReq) (res *v1.SysUserUpdateRes, err error) {
	in := &model.SysUserUpdateInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Update(ctx, *in)
	return
}

// 删除用户
func (c *cUser) Delete(ctx context.Context, req *v1.SysUserDeleteReq) (res *v1.SysUserDeleteRes, err error) {
	in := &model.SysUserDeleteInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().Delete(ctx, *in)
	return
}

// 角色分配用户
func (c *cUser) UserSelectRole(ctx context.Context, req *v1.SysUserSelectRoleReq) (res *v1.SysUserSelectRoleRes, err error) {
	in := &model.SysUserRoleUpdateUInput{}
	gconv.Scan(req, &in)
	err = service.SysUserRole().UpdateUser(ctx, *in)
	return
}

// 用户重置密码
func (c *cUser) ResetPwd(ctx context.Context, req *v1.SysUserResetPwdReq) (res *v1.SysUserResetPwdRes, err error) {
	in := &model.SysUserResetPwdInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().ResetPwd(ctx, *in)
	return
}

// 用户修改状态
func (c *cUser) ChangeStatus(ctx context.Context, req *v1.SysUserChangeStatusReq) (res *v1.SysUserChangeStatusRes, err error) {
	in := &model.SysUserChangeStatusInput{}
	gconv.Scan(req, &in)
	err = service.SysUser().ChangeStatus(ctx, *in)
	return
}
