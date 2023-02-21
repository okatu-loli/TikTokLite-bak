package response

import "github.com/okatu-loli/TikTokLite/internal/model"

// FavoriteList 用户喜欢的视频列表
type FavoriteList struct {
	CommonResponse
	Videos []model.Video `json:"video_list"`
}

// FavoriteActionResponse 点赞的响应
type FavoriteActionResponse struct {
	CommonResponse
}
