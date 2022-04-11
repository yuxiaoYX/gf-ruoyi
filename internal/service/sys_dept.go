package service

import (
	"context"
	"errors"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/text/gstr"
)

type sDept struct{}

// 部门管理服务
func SysDept() *sDept {
	return &sDept{}
}

// 获取部门列表
// TODO OmitEmpty支持特殊字符空值过滤，相关函数可以改写了
func (s *sDept) GetList(ctx context.Context, in model.SysDeptListInput) (out []*model.SysDeptOneOutput, err error) {
	m := dao.SysDept.Ctx(ctx).OmitEmpty().Where("status", in.Status)
	if in.DeptName != "" {
		m = m.Where("dept_name like ?", "%"+in.DeptName+"%")
	}
	if in.DeptId != 0 {
		m = m.Where("dept_id != ?", in.DeptId)
	}
	err = m.Scan(&out)
	return
}

// 获取部门详细信息
func (s *sDept) GetOne(ctx context.Context, in model.SysDeptOneInput) (out *model.SysDeptOneOutput, err error) {
	err = dao.SysDept.Ctx(ctx).Where("dept_id", in.DeptId).Scan(&out)
	return
}

// 新增部门
func (s *sDept) Create(ctx context.Context, in model.SysDeptCreateInput) (err error) {
	deptCount, err := dao.SysDept.Ctx(ctx).Where("dept_name=?", in.DeptName).Count()
	if err != nil {
		return err
	}
	if deptCount > 0 {
		return errors.New("部门名称已存在！")
	}
	_, err = dao.SysDept.Ctx(ctx).Insert(in)
	return
}

// 更新部门
func (s *sDept) Update(ctx context.Context, in model.SysDeptUpdateInput) (err error) {
	_, err = dao.SysDept.Ctx(ctx).OmitEmpty().Data(in).Where("dept_id=? or parent_id=?", in.DeptId, in.DeptId).Update()
	return
}

// 删除部门
func (s *sDept) Delete(ctx context.Context, in model.SysDeptDeleteInput) (err error) {
	deptIdList := gstr.Split(in.DeptIdStr, ",")
	_, err = dao.SysDept.Ctx(ctx).Delete("dept_id IN(?)", deptIdList)
	return
}

// 查询部门下拉树结构
func (s *sDept) Treeselect(ctx context.Context) (treeList []map[string]interface{}, err error) {
	var deptEntitys []*model.SysDeptOneOutput
	if err = dao.SysDept.Ctx(ctx).Where("status=0").OrderAsc("order_num").Scan(&deptEntitys); err != nil {
		return
	}
	treeList, _ = s.formDeptTree(ctx, 0, deptEntitys)
	return
}

// 构造树形部门列表
func (s *sDept) formDeptTree(ctx context.Context, parentId int64, entities []*model.SysDeptOneOutput) (treeRoute []map[string]interface{}, err error) {
	for _, entity := range entities {
		if entity.ParentId == parentId {
			subTree, err := s.formDeptTree(ctx, entity.DeptId, entities)
			if err != nil {
				return nil, err
			}
			children := make(map[string]interface{})
			children["id"] = entity.DeptId
			children["label"] = entity.DeptName
			if subTree != nil {
				children["children"] = subTree
			}
			treeRoute = append(treeRoute, children)
		}
	}
	return treeRoute, err
}
