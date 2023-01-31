package oss

import (
	"github.com/okatu-loli/TikTokLite/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	UpToken string
	//分片上传的对象，暂定分片上传
	ResumeUploader *storage.ResumeUploaderV2
)

func OSSInit() {
	//配置oss
	putPolicy := storage.PutPolicy{
		Scope: config.OSSScope,
	}
	mac := qbox.NewMac(config.OSSAccessKey, config.OSSSecretKey)
	UpToken = putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true

	ResumeUploader = storage.NewResumeUploaderV2(&cfg)

}
