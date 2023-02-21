package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"github.com/okatu-loli/TikTokLite/internal/service/videoservice"
)

type IFeedHandler interface {
	FeedList(ctx context.Context, c *app.RequestContext)
}

type FeedHandler struct {
	videoService videoservice.IVideoService
}

func NewFeedHandler() IFeedHandler {
	return FeedHandler{videoService: videoservice.NewVideoService()}
}

func (f FeedHandler) FeedList(ctx context.Context, c *app.RequestContext) {
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	list, err := f.videoService.GetFeed()
	if err != nil {
		statusCode = -1
		statusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, utils.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  statusMsg,
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
				FollowCount:   video.User.FollowCount,
				FollowerCount: video.User.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
		})
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"video_list":  videoList,
	})
}
