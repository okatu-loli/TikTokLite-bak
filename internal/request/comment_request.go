package request

// ==================================================评论相关参数=========================================================

// CommentActionParam 添加或删除评论相关参数
type CommentActionParam struct {
	UserId      int64  `form:"user_id"`
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int64  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

// CommentListParam 评论列表相关参数
type CommentListParam struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
	//Comments    *[]*model.Comment  `form:"comments"`
}
