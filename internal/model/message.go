package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserId   uint   `json:"form_user_id"`
	ToUserId uint   `json:"to_user_id"`
	Content  string `json:"content"`
}
