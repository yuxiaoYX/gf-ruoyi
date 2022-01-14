package handler

import (
	"gf-RuoYi/apiv1"
	"gf-RuoYi/internal/model"
	"gf-RuoYi/internal/service"
	"gf-RuoYi/utility/response"
	"gf-RuoYi/utility/utils"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/mssola/user_agent"
)

// GToken登录验证管理
var (
	Auth       = hAuth{}
	multiLogin = true // 后台是否允许多端同时在线
	GfToken    = &gtoken.GfToken{
		CacheMode:        1,
		CacheKey:         "GToken:",
		Timeout:          10800000, // 3个小时
		MaxRefresh:       5400000,
		TokenDelimiter:   "_",
		EncryptKey:       []byte("5awuMBAXOucydRZvb1vNRTCPaq5ME7JP"),
		AuthFailMsg:      "登录超时，请重新登录",
		MultiLogin:       multiLogin,
		LoginPath:        "/login",
		LoginBeforeFunc:  Auth.login,
		LoginAfterFunc:   Auth.loginAfter,
		LogoutPath:       "/logout",
		AuthBeforeFunc:   Auth.authBeforeFunc,
		AuthAfterFunc:    Auth.authAfterFunc,
		LogoutBeforeFunc: Auth.loginOut,
	}
)

type hAuth struct{}

//后台用户登陆验证
func (c *hAuth) login(r *ghttp.Request) (string, interface{}) {
	var ctx = r.GetCtx()
	// 数据效验
	var apiReq *apiv1.LoginDoReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	user_entity, err := service.SysUser.Login(ctx, model.SysUserLoginInput{
		UserName: apiReq.UserName,
		Password: apiReq.Password,
	})
	// 登录失败日志记录

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return user_entity.UserName, nil
}

//登录成功返回
func (c *hAuth) loginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		//保存用户在线状态token到数据库
		UserName := respData.GetString("userKey")
		token := respData.GetString("token")
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		explorer, _ := ua.Browser()
		service.SysUserOnline.Create(r.GetCtx(), model.SysUserOnlineCreateInput{
			Token: token,
			// UserId:   userId,
			UserName: UserName,
			Ip:       utils.GetClientIp(r),
			Os:       ua.OS(),
			Explorer: explorer,
		})
		response.JsonExit(r, 0, "用户登录成功", g.Map{"token": token})
	}
}

// gToken认证验证方法
func (c *hAuth) authBeforeFunc(r *ghttp.Request) bool {
	// authHeader := r.Header.Get("Authorization")
	// if authHeader != "" {
	// 	parts := strings.SplitN(authHeader, " ", 2)
	// 	if len(parts) == 2 && parts[0] == "Bearer" && parts[1] != "" {
	// 		fmt.Println(parts[1])
	// 	} else {
	// 		return false
	// 	}
	// }
	return true
}

//gToken验证后返回
func (c *hAuth) authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	glog.Debug(r.GetCtx(), respData)
	if r.Method == "OPTIONS" || respData.Success() {
		// 将用户信息写入上下文

		r.Middleware.Next()
	} else {
		response.JsonExit(r, respData.Code, "用户信息验证失败")
	}
}

//后台退出登陆
func (c *hAuth) loginOut(r *ghttp.Request) bool {
	return true
}
