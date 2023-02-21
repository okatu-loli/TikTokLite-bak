package db

import (
	"fmt"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"gorm.io/gorm"
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

// FindUserById 通过用户ID返回用户
func FindUserById(userId string) (*model.User, error) {
	res := new(model.User)

	if err := DB.Where("id = ?", userId).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil

}

func IsFollow(myId string, id string) bool {
	type Follow struct {
		gorm.Model
		UserId   string
		ToUserId string
		Status   int
	}
	follow := new(Follow)
	tx := DB.Where("user_id = ? and to_user_id = ?", myId, id).Find(&follow)
	rowsAff := tx.RowsAffected
	if rowsAff == 0 {
		return false
	}
	return true
}
