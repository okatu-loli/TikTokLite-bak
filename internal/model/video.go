package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title         string `json:"title"`
	AuthorId      uint
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
}
