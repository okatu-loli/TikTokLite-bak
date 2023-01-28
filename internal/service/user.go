package service

import (
	"github.com/henrylee2cn/goutil/errors"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
)

func Register(username string, password string) error {
	// 注册用户,这里的用户名肯定不会重复 因为数据库已经设置好
	if err := db.CreateUser(username, password); err != nil {
		return errors.New("CreateUser fail : " + err.Error())
	}

	return nil
}
