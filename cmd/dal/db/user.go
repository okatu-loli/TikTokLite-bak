package db

import (
	"fmt"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

func CreateUser(username string, password string) error {

	result := DB.Create(&model.User{UserName: username, Password: password})
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
