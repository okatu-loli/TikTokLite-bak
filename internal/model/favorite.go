package model

type Favorite struct {
	BasePo
	UserId     int64 `gorm:"size:64" json:"user_id"`
	VideoId    int64 `gorm:"size:64" json:"video_id"`
	UserInfoId int64 `json:"-"`
}

type FavoriteDAO struct {
}

var (
	favoriteDAO FavoriteDAO
)

func NewFavoriteDAO() *FavoriteDAO {
	return &favoriteDAO
}
