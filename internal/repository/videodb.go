package repository

import (
	"fmt"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"

	"github.com/okatu-loli/TikTokLite/internal/model"
)

// CreateVideo 保存视频信息
func CreateVideo(title string, playUrl string, coverUrl string, auId uint, preUrl string) error {
	result := db.DB.Create(&model.Video{
		Title:    title,
		AuthorId: auId,
		PlayUrl:  preUrl + playUrl,
		CoverUrl: preUrl + coverUrl,
	})
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetVideoList(user uint) ([]model.Video, error) {
	var res []model.Video
	// field := []string{"videos.id", "videos.play_url", "videos.cover_url", "videos.favorite_count", "videos.comment_count", "videos.title", "users.id", "users.user_name", "users.follow_count", "users.follower_count"}
	// err := DB.Model(&model.Video{}).Select(field).Preload("Author").Where("author_id = ?", user).Joins("left join users on users.id = videos.author_id").Find(&res).Error
	err := db.DB.Model(&model.Video{}).Where("author_id = ?", user).Order("created_at DESC").Preload("User").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
