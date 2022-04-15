package model

import (
	"gf-ruoyi/internal/model/entity"
)

// 获取登录日志列表输入
type SysLoginLogListInput struct {
	UserName  string // 登录账号
	Ipaddr    string // 登录IP地址
	Status    string // 登录状态（0成功 1失败）
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	PageNum   int    // 分页码
	PageSize  int    // 分页数量
}

// 获取登录日志列表输出
type SysLoginLogListOutput struct {
	Rows  []*entity.SysLoginLog // 列表
	Total int                   // 数据总数
}

// 删除登录日志输入
type SysLoginLogDeleteInput struct {
	InfoIdStr string // 需要删除的数据主键，例：[1,2,3]
}

// 新建登录日志输入
type SysLoginLogCreateInput struct {
	UserName      string // 登录账号
	Ipaddr        string // 登录IP地址
	LoginLocation string // 登录地点
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Status        string // 登录状态（0成功 1失败）
	Msg           string // 提示消息
	Err           error
}
