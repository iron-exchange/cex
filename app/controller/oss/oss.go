package oss

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "GoCEX/app/api"
	"GoCEX/internal/oss"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// Upload 独立大文件的云存储上传网关
func (c *Controller) Upload(ctx context.Context, req *v1.UploadOssReq) (res *v1.UploadOssRes, err error) {
	// 调用自定义的 OSS 模块，将二进制流直传给 Cloudflare R2
	url, err := oss.UploadFile(ctx, req.File.FileHeader)
	if err != nil {
		g.Log().Error(ctx, "OSS上传失败:", err)
		return nil, err
	}

	return &v1.UploadOssRes{Url: url}, nil
}
