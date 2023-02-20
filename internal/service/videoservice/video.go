package videoservice

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/config"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/util"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadVideoService(file *multipart.FileHeader, title string, id uint) error {

	//校验文件
	filename := file.Filename
	indexOfDot := strings.LastIndex(filename, ".") //获取文件后缀名前的.的位置
	if indexOfDot < 0 {
		logger.Error("UploadVideoService 获取文件后缀名失败！")
		return errors.New("没有获取到文件的后缀名")
	}
	suffix := filename[indexOfDot+1 : len(filename)] //获取后缀名
	suffix = strings.ToLower(suffix)                 //后缀名统一小写处理

	if !util.CheckVideo(suffix) {
		logger.Error("UploadVideoService 文件格式不支持！！")
		return errors.New("文件格式不支持")
	}
	//生成雪花id
	w, _ := util.NewWorker(5, 5).NextID()
	t := time.Now()
	//生成新的文件名
	filename = fmt.Sprintf("tiktok/%4d/%02d/%02d/", t.Year(), t.Month(), t.Day()) + strconv.FormatUint(w, 10) + "." + suffix
	//生成封面名
	coverName := fmt.Sprintf("tiktok/%4d/%02d/%02d/cover/", t.Year(), t.Month(), t.Day()) + strconv.FormatUint(w, 10) + ".pdf"
	entry := config.OSSScope + ":" + coverName
	encodedEntryURI := base64.StdEncoding.EncodeToString([]byte(entry))

	//配置oss
	putPolicy := &storage.PutPolicy{
		Scope:         config.OSSScope,
		PersistentOps: "vframe/jpg/offset/1|saveas/" + encodedEntryURI,
	}
	mac := qbox.NewMac(config.OSSAccessKey, config.OSSSecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true

	resumeUploader := storage.NewResumeUploaderV2(&cfg)

	go func() {
		//上传视频
		f, err := file.Open()
		if err != nil {
			logger.Error("UploadVideoService 视频文件无法打开！")
			return
			// return errors.New("文件无法打开！")
		}
		//使用oss
		ref := storage.PutRet{}
		rve := storage.RputV2Extra{}
		resumeUploader.Put(context.Background(), &ref, upToken, filename, f, file.Size, &rve)
		err2 := db.CreateVideo(title, ref.Key, coverName, id)
		if err2 != nil {
			logger.Error("UploadVideoService 保存视频数据失败")
			return
			// return errors.New("保存视频数据失败")
		}
	}()
	return nil
}

func GetList(uesrId uint) ([]model.Video, error) {
	vlr, err := db.GetVideoList(uesrId)
	if err != nil {
		logger.Error("GetList 获取视频失败")
		return nil, err
	}
	return vlr, nil
}

func GetFeed() ([]model.Video, error) {
	fe, err := db.GetFeed()
	if err != nil {
		logger.Error("GetFeed 获取视频失败")
		return nil, err
	}
	return fe, nil
}
