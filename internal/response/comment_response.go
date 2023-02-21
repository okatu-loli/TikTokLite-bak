package response

type Comment struct {
	CommentId  int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

// CommentAction 用户评论
type CommentAction struct {
	CommonResponse
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	CommonResponse
	CommentList []Comment `json:"comment_list"`
}
