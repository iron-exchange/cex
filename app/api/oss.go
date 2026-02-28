package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadOssReq OSS 文件上传接口
type UploadOssReq struct {
	g.Meta `path:"/oss/upload" tags:"OSS" method:"post" summary:"全局文件、图片OSS上传中转" mime:"multipart/form-data"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择需要上传的文件" dc:"二进制流"`
}

// UploadOssRes OSS 文件上传响应接口
type UploadOssRes struct {
	Url string `json:"url" dc:"上传成功后的云存储公共读取链接"`
}
