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

func FeedList(ctx context.Context, c *app.RequestContext) {
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	list, err := videoservice.GetFeed()
	if err != nil {
		statusCode = -1
		statusMsg = err.Error()
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
