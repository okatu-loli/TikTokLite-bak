package response

type VideoListResponse struct {
	VideoList []VideoRes `json:"video_list"`
}

type UserInfoRes struct {
	ID            uint   `json:"id"`
	UserName      string `json:"username"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type VideoRes struct {
	ID            uint        `json:"id"`
	Title         string      `json:"title"`
	Author        UserInfoRes `json:"author"`
	PlayUrl       string      `json:"play_url"`
	CoverUrl      string      `json:"cover_url"`
	FavoriteCount uint        `json:"favorite_count"`
	CommentCount  uint        `json:"comment_count"`
}
