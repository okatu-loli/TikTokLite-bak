package chatservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/okatu-loli/TikTokLite/cmd/dal/rdb"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/repository"
	"gorm.io/gorm"
	"time"
)

type IChatServiceImpl interface {
	SendMessage(userId uint, toUserId uint, content string) error
}

type ChatService struct {
}

func NewChatService() IChatServiceImpl {
	return ChatService{}
}

func (c ChatService) SendMessage(userId uint, toUserId uint, content string) error {
	now := time.Now()
	id, err2 := repository.SendMsg(userId, toUserId, content, now)
	if err2 != nil {
		logger.Error("发送消息失败")
		return err2
	}
	redisKey := fmt.Sprintf("chat_msg%dto%d", userId, toUserId)
	logger.Debug(id)
	chatMsg := model.Message{
		Model: gorm.Model{
			ID: id,
		},
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
	}
	data, err3 := json.Marshal(chatMsg)
	if err3 != nil {
		logger.Error("发送消息失败")
		return err3
	}
	err := rdb.RDB.RPush(context.Background(), redisKey, data).Err()
	if err != nil {
		logger.Error("发送消息失败")
		return err
	}
	return nil
}
