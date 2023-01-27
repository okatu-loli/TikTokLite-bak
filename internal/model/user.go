package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string // 昵称
	Password      string // 密码, MD5 加密后不可逆
	FollowCount   int    // 关注人数
	FollowerCount int    // 粉丝人数
}
