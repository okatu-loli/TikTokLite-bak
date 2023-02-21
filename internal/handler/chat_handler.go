package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/chatservice"
	"strconv"
)

type IChatHandler interface {
	SendChat(ctx context.Context, c *app.RequestContext)
}

type ChatHandler struct {
	chatService chatservice.IChatServiceImpl
}

func NewChatHandler() IChatHandler {
	return ChatHandler{chatService: chatservice.NewChatService()}
}

func (ch ChatHandler) SendChat(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get("id")
	userId := user.(*model.User).ID
	tuId, _ := strconv.Atoi(c.Query("to_user_id"))
	toUserId := uint(tuId)
	content := c.Query("content")
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	err := ch.chatService.SendMessage(userId, toUserId, content)

	if err != nil {
		statusCode = -1 //暂定
		statusMsg = err.Error()
		return
	}
	statusCode = 0
	statusMsg = "发送成功"
	defer c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
	})
}
