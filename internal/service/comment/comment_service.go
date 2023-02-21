package comment

import (
	"github.com/jinzhu/copier"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/repository"
	"github.com/okatu-loli/TikTokLite/internal/request"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"log"
	"time"
)

type ICommentService interface {
	CommentPost(userId int64, commentPostRequest request.CommentActionParam) response.CommentAction
	ListComment(commentListRequest request.CommentListParam) response.CommentListResponse
	DeleteComment(commentPostRequest request.CommentActionParam) error
}

type CommentsService struct {
	CommentRepository repository.ICommentRepository
	UserRepository    repository.IUserRepository
}

// ListComment 获取评论列表
func (c CommentsService) ListComment(commentListRequest request.CommentListParam) response.CommentListResponse {
	videoId := commentListRequest.VideoId
	//[]commentListRequest
	//commentList := commentListRequest.Comments
	commentList, err := c.CommentRepository.QueryCommentListByVideoId(videoId)
	if err != nil {
		log.Printf("ListComment|数据库获取错误|%v", err)
		return response.CommentListResponse{}
	}

	var responseComment []response.Comment
	err = copier.Copy(&responseComment, &commentList)
	if err != nil {
		log.Printf("ListComment|获取用户失败|%v", err)
		return response.CommentListResponse{}
	}

	for index, comment := range commentList {
		userId := comment.UserId

		user, err := c.UserRepository.GetUserById(userId)
		if err != nil {
			log.Printf("ListComment|获取用户失败|%v", err)
			return response.CommentListResponse{}
		}
		var responseUser response.User
		err = copier.Copy(&responseUser, &user)
		if err != nil {
			log.Printf("ListComment|对象复制错误|%v", err)
			return response.CommentListResponse{}
		}

		responseComment[index].User = responseUser
	}
	return response.CommentListResponse{
		CommonResponse: response.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: responseComment,
	}
}

// CommentPost 发表评论
func (c CommentsService) CommentPost(userId int64, commentPostRequest request.CommentActionParam) response.CommentAction {
	var userResponse response.User
	user, err := c.UserRepository.GetUserById(userId)
	if err != nil {
		log.Printf("CommentPost|类型转换错误|%v", err)
		return response.CommentAction{}
	}
	//就是将user对象中的字段赋值到userResponse的同名字段中。如果目标对象中没有同名的字段，则该字段被忽略。
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		log.Printf("CommentPost|对象转化错误|%v", err)
		return response.CommentAction{}
	}

	comment := model.Comment{UserId: userId, VideoId: commentPostRequest.VideoId, Content: commentPostRequest.CommentText}
	err = c.CommentRepository.InsertComment(&comment)
	if err != nil {
		log.Printf("CommentPost|数据库插入错误|%v", err)
		return response.CommentAction{}
	}

	return response.CommentAction{
		CommonResponse: response.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		Comment: response.Comment{
			CommentId:  comment.ID,
			Content:    comment.Content,
			CreateDate: time.Now().Format("01-02"),
			User:       userResponse,
		},
	}
}

// DeleteComment 删除评论
func (c CommentsService) DeleteComment(commentPostRequest request.CommentActionParam) error {
	videoId := commentPostRequest.VideoId
	commentId := commentPostRequest.CommentId
	userId := commentPostRequest.UserId
	//获取comment
	comment := model.Comment{UserId: userId, VideoId: videoId, Content: commentPostRequest.CommentText}
	err := c.CommentRepository.QueryCommentById(commentId, &comment)
	if err != nil {
		return err
	}
	//删除comment
	err = c.CommentRepository.DeleteCommentAndUpdateCountById(commentId, videoId)
	if err != nil {
		return err
	}
	return nil
}

func NewCommentService() ICommentService {
	return CommentsService{
		CommentRepository: repository.NewCommentRepository(),
		UserRepository:    repository.NewUserRepository(),
	}
}
