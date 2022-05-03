package service

import (
	"context"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type sJob struct{}

// 定时任务管理服务
func SysJob() *sJob {
	return &sJob{}
}

// 获取定时任务列表
func (s *sJob) GetList(ctx context.Context, in model.SysJobListInput) (out model.SysJobListOutput, err error) {
	m := dao.SysJob.Ctx(ctx).OmitEmpty().Where(g.Map{
		"job_group": in.JobGroup,
		"status":    in.Status,
	})
	if in.JobName != "" {
		m = m.Where("job_name LIKE ?", "%"+in.JobName+"%")
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

// 获取定时任务详细信息
func (s *sJob) GetOne(ctx context.Context, in model.SysJobOneInput) (out *model.SysJobOneOutput, err error) {
	err = dao.SysJob.Ctx(ctx).Where("job_id", in.JobId).Scan(&out)
	return
}

// 新增定时任务
func (s *sJob) Create(ctx context.Context, in model.SysJobCreateInput) (err error) {
	_, err = dao.SysJob.Ctx(ctx).Insert(in)
	return
}

// 更新定时任务
func (s *sJob) Update(ctx context.Context, in model.SysJobUpdateInput) (err error) {
	_, err = dao.SysJob.Ctx(ctx).OmitEmpty().Data(in).Where("job_id", in.JobId).Update()
	return
}

// 删除定时任务
func (s *sJob) Delete(ctx context.Context, in model.SysJobDeleteInput) (err error) {
	_, err = dao.SysJob.Ctx(ctx).Delete("job_id IN(?)", in.JobIdList)
	return
}
