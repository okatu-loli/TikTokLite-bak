package response

import "time"

type MessageRes struct {
	Id         uint      `json:"id"`
	ToUserId   uint      `json:"to_user_id"`
	FromUserId uint      `json:"from_user_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
