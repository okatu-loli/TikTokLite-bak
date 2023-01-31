package videoservice

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/config/oss"
	"github.com/okatu-loli/TikTokLite/internal/service/util"
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

	//上传视频
	f, err := file.Open()
	if err != nil {
		logger.Error("UploadVideoService 视频文件无法打开！")
		return errors.New("文件无法打开！")
	}
	//使用oss
	ref := storage.PutRet{}
	rve := storage.RputV2Extra{}
	oss.ResumeUploader.Put(context.Background(), &ref, oss.UpToken, filename, f, file.Size, &rve)
	err2 := db.CreateVideo(title, ref.Key, "text", id)
	if err2 != nil {
		logger.Error("UploadVideoService 保存视频数据失败")
		return errors.New("保存视频数据失败")
	}
	return nil
}
