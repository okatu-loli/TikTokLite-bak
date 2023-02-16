package db

import (
	"fmt"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

func CreateUser(username string, password string) (*model.User, error) {

	user := model.User{UserName: username, Password: password}
	result := DB.Create(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CheckUser jwt从数据库检查用户
func CheckUser(username string, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)

	if err := DB.Where("user_name = ?", username).
		Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
