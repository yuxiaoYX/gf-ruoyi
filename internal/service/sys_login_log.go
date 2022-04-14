package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/model/entity"
	"gf-ruoyi/internal/service/internal/dao"
	"gf-ruoyi/utility/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mssola/user_agent"
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
		m = m.Where("login_time>? and login_time<?", in.BeginTime, in.EndTime)
	}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.Rows); err != nil {
		return
	}
	out.Total = len(out.Rows)
	return
}

// 删除登录日志
func (s *sLoginLog) Delete(ctx context.Context, in model.SysLoginLogDeleteInput) (err error) {
	InfoIdList := gstr.Split(in.InfoIdStr, ",")
	for _, v := range InfoIdList {
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
func (s *sLoginLog) Create(ctx context.Context, userName string, err1 error) (err2 error) {
	loginData := &entity.SysLoginLog{}
	r := g.RequestFromCtx(ctx)
	userAgent := r.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	loginData.UserName = userName
	loginData.Ipaddr = r.GetClientIp()
	loginData.LoginLocation = utils.GetCityByIp(ctx, r.GetClientIp())
	loginData.Browser, _ = ua.Browser()
	loginData.Os = ua.OS()
	if err1 != nil {
		loginData.Status = "1"
		loginData.Msg = err1.Error()
	} else {
		loginData.Status = "0"
		loginData.Msg = "登录成功"
	}
	loginData.LoginTime = gtime.Now()
	_, err2 = dao.SysLoginLog.Ctx(ctx).Save(loginData)
	return
}
