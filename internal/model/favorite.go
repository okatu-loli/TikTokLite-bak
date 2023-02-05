package model

type Favorite struct {
	//FavoriteId相当于之前的id
	//FavoriteId  int64 `gorm:"Column:favorite_id;AUTO_INCREMENT" json:"favorite_id"`
	BasePo
	UserId  int64 `gorm:"size:64" json:"user_id"`
	VideoId int64 `gorm:"size:64" json:"video_id"`
}

//
//type FavoriteDAO struct {
//}
//
//var (
//	favoriteDAO FavoriteDAO
//)
//
//func NewFavoriteDAO() *FavoriteDAO {
//	return &favoriteDAO
//}
