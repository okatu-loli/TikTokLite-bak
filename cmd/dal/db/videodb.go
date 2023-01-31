package db

import (
	"fmt"

	"github.com/okatu-loli/TikTokLite/internal/model"
)

// CreateVideo 保存视频信息
func CreateVideo(title string, playUrl string, coverUrl string, auId uint) error {
	result := DB.Create(&model.Video{
		Title:    title,
		AuthorId: auId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	})
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
