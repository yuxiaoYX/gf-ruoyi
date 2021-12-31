// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import(
"github.com/gogf/gf/v2/os/gtime"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline struct {
    Id        uint64      `json:"id"        `//            
    Token     string      `json:"token"     `// 用户token  
    UserName  string      `json:"userName"  `// 用户名     
    Ip        string      `json:"ip"        `// 登录ip     
    Explorer  string      `json:"explorer"  `// 浏览器     
    Os        string      `json:"os"        `// 操作系统   
    CreatedAt *gtime.Time `json:"createdAt" `// 登录时间   
}