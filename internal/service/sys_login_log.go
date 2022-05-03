package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type sLoginLog struct{}

// 登录日志管理服务
func SysLoginLog() *sLoginLog {
	return &sLoginLog{}
}

// 获取登录日志列表
func (s *sLoginLog) GetList(ctx context.Context, in model.SysLoginLogListInput) (out model.SysLoginLogListOutput, err error) {
	m := dao.SysLoginLog.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": in.UserName,
		"status":    in.Status,
	})
	if in.Ipaddr != "" {
		m = m.Where("ipaddr LIKE ?", "%"+in.Ipaddr+"%")
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

// 删除登录日志
func (s *sLoginLog) Delete(ctx context.Context, in model.SysLoginLogDeleteInput) (err error) {
	for _, v := range in.InfoIdList {
		if _, err = dao.SysLoginLog.Ctx(ctx).Delete("info_id=?", v); err != nil {
			return
		}
	}
	return
}

// 清空登录日志
func (s *sLoginLog) Clean(ctx context.Context) (err error) {
	_, err = dao.SysLoginLog.Ctx(ctx).Where("info_id>0").Delete()
	return
}

// 新增登录日志
// TODO 如果使用，用户昵称登录，保存的userName是nickName
func (s *sLoginLog) Create(ctx context.Context, in model.SysLoginLogCreateInput) (err error) {
	if in.Err != nil {
		in.Status = "1"
		in.Msg = in.Err.Error()
	} else {
		in.Status = "0"
		in.Msg = "登录成功"
	}
	_, err = dao.SysLoginLog.Ctx(ctx).Save(in)
	return
}
