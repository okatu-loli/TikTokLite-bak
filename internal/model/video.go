package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title         string `json:"title"`
	AuthorId      uint   `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	User          User   `json:"author" gorm:"foreignKey:author_id"`
}
