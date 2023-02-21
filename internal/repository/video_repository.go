package repository

import (
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"log"
)

type IVideoRepository interface {
	GetVideoListByUserId(userId int64) ([]model.Video, error)       // 获取视频列表
	GetVideoListByVideoIds(videoIds []int64) ([]model.Video, error) // 获取视频列表

}

// VideoRepository 定义一个结构体
type VideoRepository struct {
}

// GetVideoListByVideoIds 根据id查找视频
func (v VideoRepository) GetVideoListByVideoIds(videoIds []int64) ([]model.Video, error) {
	videos := make([]model.Video, 10)
	err := db.DB.Where("video_id in (?)", videoIds).Find(&videos).Error
	return videos, err
}

// NewVideoRepository VideoRepository构造函数
func NewVideoRepository() IVideoRepository {
	return VideoRepository{}
}

// GetVideoListByUserId 获取某个用户的视频
func (v VideoRepository) GetVideoListByUserId(userId int64) ([]model.Video, error) {
	log.Printf("GetVideoListByUserId|获取用户的视频|%v", userId)
	var video []model.Video
	err := db.DB.Where("user_id = ?", userId).Find(&video).Error

	return video, err
}
