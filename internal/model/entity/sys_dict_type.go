// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package entity

import(
"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
    DictId    int64       `json:"dictId"    `// 字典主键             
    DictName  string      `json:"dictName"  `// 字典名称             
    DictType  string      `json:"dictType"  `// 字典类型             
    Status    string      `json:"status"    `// 状态（0正常 1停用）  
    Remark    string      `json:"remark"    `// 备注                 
    CreatedAt *gtime.Time `json:"createdAt" `// 创建日期             
    UpdatedAt *gtime.Time `json:"updatedAt" `// 修改日期             
    DeletedAt *gtime.Time `json:"deletedAt" `// 删除日期             
}