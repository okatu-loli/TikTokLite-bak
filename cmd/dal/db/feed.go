package db

import "github.com/okatu-loli/TikTokLite/internal/model"

func GetFeed() ([]model.Video, error) {
	var res []model.Video
	// field := []string{"videos.id", "videos.play_url", "videos.cover_url", "videos.favorite_count", "videos.comment_count", "videos.title", "users.id", "users.user_name", "users.follow_count", "users.follower_count"}
	// err := DB.Model(&model.Video{}).Select(field).Preload("Author").Where("author_id = ?", user).Joins("left join users on users.id = videos.author_id").Find(&res).Error
	err := DB.Model(&model.Video{}).Order("created_at DESC").Limit(20).Preload("User").Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
