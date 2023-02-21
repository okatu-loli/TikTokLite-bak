package repository

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"gorm.io/gorm"
	"time"
)

func SendMsg(userId uint, toUserId uint, content string, now time.Time) (uint, error) {
	msg := &model.Message{
		Model:    gorm.Model{CreatedAt: now},
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
	}
	result := db.DB.Create(msg)
	logger.Info(result.RowsAffected)
	logger.Info(result.Error)
	if result.Error != nil {
		return 0, result.Error
	}

	return msg.ID, nil
}
