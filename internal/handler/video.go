package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/videoservice"
)

// UploadVideo 上传视频绑定接口
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	video, _ := c.FormFile("file")
	title := c.PostForm("title")
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	user, _ := c.Get("id")
	err := videoservice.UploadVideoService(video, title, user.(*model.User).ID)
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
