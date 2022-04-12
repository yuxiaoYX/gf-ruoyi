package service

import (
	"context"
	"gf-ruoyi/internal/consts"
	"gf-ruoyi/internal/model"
	"gf-ruoyi/internal/model/entity"
	"gf-ruoyi/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

type sFile struct{}

// File 文件管理服务
func SysFile() *sFile {
	return &sFile{}
}

// 同一上传文件
func (s *sFile) Upload(ctx context.Context, in model.SysFileUploadInput) (*model.SysFileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.SysFile.Ctx(ctx).
		Where(dao.SysFile.Columns().UserId, Context().Get(ctx).User.UserId).
		WhereGTE(dao.SysFile.Columns().CreatedAt, gtime.Now().Add(-1*time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	// 记录到数据表

	data := entity.SysFile{
		Name: fileName,
		Src:  gfile.Join(uploadPath, dateDirName, fileName),
		// TODO Url填写网络地址
		// Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId:   Context().Get(ctx).User.UserId,
		FileType: in.FileType,
	}
	result, err := dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &model.SysFileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Path: data.Src,
		Url:  data.Url,
	}, nil
}
