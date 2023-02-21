package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/okatu-loli/TikTokLite/internal/request"
	"github.com/okatu-loli/TikTokLite/internal/service/comment"
	"log"
	"net/http"
)

type ICommentController interface {
	CommentPost(ctx context.Context, c *app.RequestContext) // 发表评论
	ListComment(ctx context.Context, c *app.RequestContext) // 获取评论列表
}

type CommentController struct {
	CommentService comment.ICommentService
}

func (c2 CommentController) ListComment(ctx context.Context, c *app.RequestContext) {
	var listRequest request.CommentListParam
	//BindQuery会自动返回400状态码并将Content-Type 被设置为 text/plain; charset=utf-8，ShouldBindQuery则不会。
	//BindQuery改成了BindAndValidate
	err := c.BindAndValidate(&listRequest)
	if err != nil {
		log.Printf("PostComment|参数错误|%v", listRequest)
		return
	}

	commentListResponse := c2.CommentService.ListComment(listRequest)
	c.JSON(http.StatusOK, commentListResponse)
}

func (c2 CommentController) CommentPost(ctx context.Context, c *app.RequestContext) {
	var postRequest request.CommentActionParam
	err := c.BindAndValidate(&postRequest)
	if err != nil {
		log.Printf("PostComment|参数错误|%v", postRequest)
		return
	}
	value, _ := c.Get("userId")
	commentPostResponse := c2.CommentService.CommentPost(value.(int64), postRequest)
	c.JSON(http.StatusOK, commentPostResponse)
}

func NewCommentController() ICommentController {
	return CommentController{CommentService: comment.NewCommentService()}
}
