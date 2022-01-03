// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import(
"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
    UserId    uint        `json:"userId"    `// 用户ID                   
    UserName  string      `json:"userName"  `// 用户账号                 
    NickName  string      `json:"nickName"  `// 用户昵称                 
    Password  string      `json:"password"  `// 登录密码                 
    Mobile    string      `json:"mobile"    `// 手机号                   
    Avatar    string      `json:"avatar"    `// 用户头像地址             
    Status    string      `json:"status"    `// 用户状态；0:禁用,1:正常  
    RoleIds   string      `json:"roleIds"   `// 角色id字符串             
    ShopIds   string      `json:"shopIds"   `// 店铺ID字符串             
    DeptId    string      `json:"deptId"    `// 部门ID                   
    Remark    string      `json:"remark"    `// 备注                     
    CreatedAt *gtime.Time `json:"createdAt" `// 创建时间                 
    UpdatedAt *gtime.Time `json:"updatedAt" `// 更新时间                 
    DeletedAt *gtime.Time `json:"deletedAt" `// 删除时间                 
}