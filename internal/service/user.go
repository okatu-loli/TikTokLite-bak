package service

import (
	"github.com/henrylee2cn/goutil/errors"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

func Register(username string, password string) (*model.User, error) {
	// 注册用户,这里的用户名肯定不会重复 因为数据库已经设置好
	var user *model.User
	var err error
	if user, err = db.CreateUser(username, password); err != nil {
		return nil, errors.New("CreateUser fail : " + err.Error())
	}

	return user, nil
}

func FindUserById(id string) (*model.User, error) {
	res, err := db.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func IsFollow(myId string, id string) bool {
	return db.IsFollow(myId, id)
}
