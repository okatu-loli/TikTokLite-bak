package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"github.com/okatu-loli/TikTokLite/internal/service/chatservice"
	"strconv"
)

type IChatHandler interface {
	SendChat(ctx context.Context, c *app.RequestContext)
	ChatMeg(ctx context.Context, c *app.RequestContext)
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

func (ch ChatHandler) ChatMeg(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get("id")
	userId := user.(*model.User).ID
	tuId, _ := strconv.Atoi(c.Query("to_user_id"))
	toUserId := uint(tuId)
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态

	messages, err := ch.chatService.GetMsg(userId, toUserId)
	if err != nil {
		statusCode = -1 //暂定
		statusMsg = err.Error()
		return
	}
	statusCode = 0
	statusMsg = "发送成功"
	var list []response.MessageRes
	for _, message := range messages {
		list = append(list, response.MessageRes{
			Id:         message.ID,
			ToUserId:   message.ToUserId,
			FromUserId: message.UserId,
			Content:    message.Content,
			CreateTime: message.CreatedAt,
		})
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code":  statusCode,
		"status_msg":   statusMsg,
		"message_list": list,
	})
}
