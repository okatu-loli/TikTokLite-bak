package request

// ====================================================点赞相关参数=======================================================

// FavoriteActionParam 点赞或取消点赞相关参数
type FavoriteActionParam struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int64  `form:"action_type"`
}

// FavoriteListParam 喜欢列表相关参数
type FavoriteListParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}
