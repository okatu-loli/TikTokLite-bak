package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/okatu-loli/TikTokLite/internal/request"
	"github.com/okatu-loli/TikTokLite/internal/service/comment"
	"log"
	"net/http"
)

type ICommentController interface {
	CommentPost(c *gin.Context) // 发表评论
	ListComment(c *gin.Context) // 获取评论列表
}

type CommentController struct {
	CommentService comment.ICommentService
}

func (c2 CommentController) ListComment(c *gin.Context) {
	var listRequest request.CommentListParam
	//BindQuery会自动返回400状态码并将Content-Type 被设置为 text/plain; charset=utf-8，ShouldBindQuery则不会。
	err := c.ShouldBindQuery(&listRequest)
	if err != nil {
		log.Printf("PostComment|参数错误|%v", listRequest)
		return
	}

	commentListResponse := c2.CommentService.ListComment(listRequest)
	c.JSON(http.StatusOK, commentListResponse)
}

func (c2 CommentController) CommentPost(c *gin.Context) {
	var postRequest request.CommentActionParam
	err := c.ShouldBindQuery(&postRequest)
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
