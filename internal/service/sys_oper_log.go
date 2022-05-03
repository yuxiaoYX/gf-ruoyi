package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/model/entity"
	"gf-ruoyi/internal/service/internal/dao"
	"gf-ruoyi/utility/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOperLog struct{}

// 操作日志管理
func SysOperLog() *sOperLog {
	return &sOperLog{}
}

// 获取操作日志列表
func (s *sOperLog) GetList(ctx context.Context, in model.SysOperLogListInput) (out model.SysOperLogListOutput, err error) {
	m := dao.SysOperLog.Ctx(ctx).OmitEmpty().Where(g.Map{
		"user_name": in.BusinessType,
		"status":    in.Status,
	})
	if in.Title != "" {
		m = m.Where("title LIKE ?", "%"+in.Title+"%")
	}
	if in.OperName != "" {
		m = m.Where("oper_name LIKE ?", "%"+in.OperName+"%")
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

// 删除操作日志
func (s *sOperLog) Delete(ctx context.Context, in model.SysOperLogDeleteInput) (err error) {
	for _, v := range in.OperIdList {
		if _, err = dao.SysOperLog.Ctx(ctx).Delete("oper_id=?", v); err != nil {
			return
		}
	}
	return
}

// 清空操作日志
func (s *sOperLog) Clean(ctx context.Context) (err error) {
	_, err = dao.SysOperLog.Ctx(ctx).Where("oper_id>0").Delete()
	return
}

// 添加操作日志
func (s *sOperLog) Create(ctx context.Context) (err error) {

	r := g.RequestFromCtx(ctx)
	url := r.Request.URL //请求地址
	// 用户登录，不记录操作日志
	if gstr.ContainsI(url.Path, "login") {
		return
	}

	var menuEntity *entity.SysMenu
	perms := gstr.TrimLeft(url.Path, "/api")
	err = dao.SysMenu.Ctx(ctx).Where("is_log=0 AND perms=?", perms).Scan(&menuEntity)
	if err != nil || menuEntity == nil {
		return
	}

	var operInfo entity.SysOperLog
	rawQuery := url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	operInfo.Title = menuEntity.MenuName
	operInfo.Method = url.Path
	operInfo.RequestMethod = r.Method
	operInfo.OperatorType = 0
	operInfo.OperName = Context().Get(ctx).UserInfo.User.UserName
	operInfo.DeptName = Context().Get(ctx).UserInfo.Dept.DeptName
	operInfo.OperUrl = url.Path + rawQuery
	operInfo.OperIp = r.GetClientIp()
	operInfo.OperLocation = utils.GetCityByIp(ctx, r.GetClientIp())

	methodList := gstr.SplitAndTrim(url.Path, "/")
	method := methodList[len(methodList)-1]
	if gstr.ContainsI(method, "create") {
		operInfo.BusinessType = 1
	} else if gstr.ContainsI(method, "delete") {
		operInfo.BusinessType = 2
	} else if gstr.ContainsI(method, "update") {
		operInfo.BusinessType = 3
	} else if gstr.ContainsI(method, "get") {
		operInfo.BusinessType = 4
	} else {
		operInfo.BusinessType = 0
	}
	params := r.GetMap()
	if params != nil {
		b, _ := gjson.Encode(params)
		if len(b) > 0 {
			operInfo.OperParam = string(b)
		}
	}
	operInfo.JsonResult = gconv.String(r.GetHandlerResponse())
	reqErr := r.GetError()
	if reqErr != nil {
		operInfo.Status = 1
		operInfo.ErrorMsg = reqErr.Error()
	} else {
		operInfo.Status = 0
	}

	dao.SysOperLog.Ctx(ctx).Insert(operInfo)
	return
}
