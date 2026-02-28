package oss

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// 请用下方你给我的密钥进行初始化
const (
	r2AccountId = "0fecf98846d30faea6b4cac6bb09bfe7"
	r2AccessKey = "86d9fa7b402873fdde8359361aa03cc3"
	r2SecretKey = "6048b8103d4bcfaf20ccc33d5c4f6929062da1721767a5d43fd39b2ce95f5408"
	r2Bucket    = "cex"
	r2PublicUrl = "https://pub-89269418e8bf4556904fd21774ca64e1.r2.dev"
)

var s3Client *s3.Client

func init() {
	// 初始化 Cloudflare R2 的连接配置
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2AccountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2AccessKey, r2SecretKey, "")),
		config.WithRegion("auto"), // R2 推荐配置
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to load R2 SDK config, err: %v", err))
	}

	s3Client = s3.NewFromConfig(cfg)
}

// UploadFile 流式上传文件到 Cloudflare R2 对象存储，并返回公共访问链接
func UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 生成无碰撞的文件名: 2026/02/28/uuid.jpg
	ext := getFileExtension(file.Filename)
	datePath := time.Now().Format("2006/01/02")
	objectKey := fmt.Sprintf("uploads/%s/%s%s", datePath, uuid.New().String(), ext)

	// 提取文件扩展名猜测 Content-Type
	contentType := "application/octet-stream"
	if ext == ".png" {
		contentType = "image/png"
	} else if ext == ".jpg" || ext == ".jpeg" {
		contentType = "image/jpeg"
	}

	// 注入到 Bucket
	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2Bucket),
		Key:         aws.String(objectKey),
		Body:        src,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	// 拼接公网 CDN 提取链接
	return fmt.Sprintf("%s/%s", r2PublicUrl, objectKey), nil
}

func getFileExtension(filename string) string {
	idx := strings.LastIndex(filename, ".")
	if idx == -1 {
		return ""
	}
	return strings.ToLower(filename[idx:])
}
