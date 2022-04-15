package service

import (
	"context"
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

// 添加操作日志
func (s *sOperLog) Create(ctx context.Context) (err error) {
	// var operInfo *entity.SysOperLog

	r := g.RequestFromCtx(ctx)
	url := r.Request.URL //请求地址
	perms := gstr.TrimLeft(url.Path, "/api")
	menuName, err := dao.SysMenu.Ctx(ctx).Fields("menu_name").Where("perms=?", perms).Value()
	if err != nil {
		return
	}
	rawQuery := url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	ipaddr := r.GetClientIp()
	location := utils.GetCityByIp(ctx, r.GetClientIp())
	operInfo := entity.SysOperLog{
		Title:         menuName.String(),
		BusinessType:  0, // TODO 未完成
		Method:        url.Path,
		RequestMethod: r.Method,
		OperatorType:  0,
		OperName:      Context().Get(ctx).UserInfo.User.UserName,
		DeptName:      Context().Get(ctx).UserInfo.Dept.DeptName,
		OperUrl:       url.Path + rawQuery,
		OperIp:        ipaddr,
		OperLocation:  location,
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

	g.Log().Info(ctx, operInfo)

	// dao.SysOperLog.Ctx(ctx).Insert(entity.SysOperLog{
	// 	Title: ,
	// })

	return
}
