package chatservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/go-redis/redis/v8"
	"github.com/okatu-loli/TikTokLite/cmd/dal/rdb"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/repository"
	"gorm.io/gorm"
	"time"
)

type IChatServiceImpl interface {
	SendMessage(userId uint, toUserId uint, content string) error
	GetMsg(userId uint, toUserId uint) ([]model.Message, error)
}

type ChatService struct {
}

func NewChatService() IChatServiceImpl {
	return ChatService{}
}

func (c ChatService) SendMessage(userId uint, toUserId uint, content string) error {
	//获取当前时间
	now := time.Now()
	//讲发送信息，存入mysql中
	id, err2 := repository.SendMsg(userId, toUserId, content, now)
	if err2 != nil {
		logger.Error("发送消息失败")
		return err2
	}
	//格式化出来一个key值
	redisKey := fmt.Sprintf("chat_msg%dto%d", userId, toUserId)
	//序列化这个结构体
	chatMsg := model.Message{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: now,
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
	//存入redis中，并设置过期时间
	err := rdb.RDB.RPush(context.Background(), redisKey, data).Err()
	_, err4 := rdb.RDB.Expire(context.Background(), redisKey, time.Hour*2).Result()
	if err != nil || err4 != nil {
		logger.Error("发送消息失败")
		return err
	}
	return nil
}

func (c ChatService) GetMsg(userId uint, toUserId uint) ([]model.Message, error) {
	//格式化key
	redisKey := fmt.Sprintf("chat_msg%dto%d", userId, toUserId)
	//先从redis查询
	result, err := rdb.RDB.LRange(context.Background(), redisKey, 0, -1).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	//如果查询到直接返回
	if err != redis.Nil {
		var r []model.Message
		for _, s := range result {
			var m model.Message
			err6 := json.Unmarshal([]byte(s), &m)
			if err6 != nil {
				return nil, err6
			}
			r = append(r, m)
		}
		if err != nil {
			return nil, err
		}
		logger.Debug("我走的是缓存")
		return r, nil
	}
	//否则进行mysql查询
	msg, err2 := repository.GetMsg(userId, toUserId)
	if err2 != nil {
		logger.Error("数据库查询聊天记录失败")
		return nil, err2
	}
	//开启协程来进行异步的缓存写入
	go func() {
		data, _ := json.Marshal(msg)
		_, err4 := rdb.RDB.Del(context.Background(), redisKey).Result()
		_, err3 := rdb.RDB.RPush(context.Background(), redisKey, data).Result()
		_, err5 := rdb.RDB.Expire(context.Background(), redisKey, time.Hour*2).Result()
		if err5 != nil || err4 != nil || err3 != nil {
			logger.Error("redis出现故障，消息模块！！！")
		}
	}()
	return msg, nil
}
