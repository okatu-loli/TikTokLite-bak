package response

import "github.com/okatu-loli/TikTokLite/internal/model"

// FavoriteList 用户点赞的视频列表
type FavoriteList struct {
	CommonResponse
	VideoList []model.Video `json:"video_list"`
}

// CommentAction 用户评论
type CommentAction struct {
	CommonResponse
	Comment model.Comment `json:"comment"`
}

type CommentListResponse struct {
	CommonResponse
	CommentList []model.Comment `json:"comment_list"`
}
