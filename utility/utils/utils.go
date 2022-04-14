package utils

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

//获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

//获取ip所属城市
func GetCityByIp(ctx context.Context, ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" || gstr.Contains(ip, "172.") {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(ctx, url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if gconv.Int(json.Get("code")) == 0 {
		city := fmt.Sprintf("%s %s", gconv.String(json.Get("pro")), gconv.String(json.Get("city")))
		return city
	} else {
		return ""
	}
}
