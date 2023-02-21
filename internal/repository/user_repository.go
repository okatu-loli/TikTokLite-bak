package repository

import (
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"log"
)

type IUserRepository interface {
	GetUserById(id int64) (model.UserInfo, error) // 获取单个用户

	CheckUser(userName string, password string) (model.UserInfo, error) // 检查是否存在用户
}

// UserRepository 定义一个结构体
type UserRepository struct {
}

// NewUserRepository UserRepository构造函数
func NewUserRepository() IUserRepository {
	return UserRepository{}
}

// CheckUser 检查用户是否存在
func (ur UserRepository) CheckUser(userName string, password string) (model.UserInfo, error) {
	var user model.UserInfo
	err := db.DB.Where("username = ? and password = ?", userName, password).Scan(&user).Error
	return user, err
}

// GetUserById 获取单个用户
func (ur UserRepository) GetUserById(id int64) (model.UserInfo, error) {
	log.Printf("GetUserById|获取用户信息|%v", id)
	var user model.UserInfo
	err := db.DB.Where("user_id = ?", id).First(&user).Error

	return user, err
}
