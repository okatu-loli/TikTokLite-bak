package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"github.com/okatu-loli/TikTokLite/internal/service/videoservice"
)

type IVideoHandler interface {
	UploadVideo(ctx context.Context, c *app.RequestContext)
	PublishList(ctx context.Context, c *app.RequestContext)
}

type VideoHandle struct {
	videoservice videoservice.IVideoService
}

func NewVideoHandler() IVideoHandler {
	return VideoHandle{videoservice: videoservice.NewVideoService()}
}

// UploadVideo 上传视频绑定接口
func (v VideoHandle) UploadVideo(ctx context.Context, c *app.RequestContext) {
	video, _ := c.FormFile("file")
	title := c.PostForm("title")
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	user, _ := c.Get("id")
	err := v.videoservice.UploadVideoService(video, title, user.(*model.User).ID)
	if err != nil {
		statusCode = -1 //暂定
		statusMsg = err.Error()
		return
	}
	statusCode = 0
	statusMsg = "上传成功！"
	defer c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
	})
}

func (v VideoHandle) PublishList(ctx context.Context, c *app.RequestContext) {
	user, err := strconv.Atoi(c.Query("user_id"))
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}

	list, err2 := v.videoservice.GetList(uint(user))
	if err2 != nil {
		statusCode = -1
		statusMsg = err2.Error()
		c.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取视频列表失败",
			"error":       err.Error(),
		})
		return
	}
	statusCode = 0
	statusMsg = "成功"
	videoList := []response.VideoRes{}
	for _, video := range list {
		videoList = append(videoList, response.VideoRes{
			ID:    video.ID,
			Title: video.Title,
			Author: response.UserInfoRes{
				ID:            video.User.ID,
				UserName:      video.User.UserName,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: 0,
			CommentCount:  0,
		})
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"video_list":  videoList,
	})

}
