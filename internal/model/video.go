package model

import (
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
  "gorm.io/gorm"
	"log"
	"sync"
)

type Video struct {
	BasePo
	//videoId相当于之前的id
	//VideoId       int64     `gorm:"Column:video_id" json:"video_id"`
	//这里面的AuthorId相当于user_id
	AuthorId      int64  `gorm:"Column:author_id" json:"author_id"`
	Title         string `gorm:"Column:title" json:"title"`
	PlayUrl       string `gorm:"Column:play_url" json:"play_url"`
	CoverUrl      string `gorm:"Column:cover_url" json:"cover_url"`
	FavoriteCount int64  `gorm:"Column:favorite_count" json:"favorite_count"`
	CommentCount  int64  `gorm:"Column:comment_count" json:"comment_count"`
}
type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}

// PlusOneFavorByUserIdAndVideoId 增加一个赞
//func (v *VideoDAO) PlusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
//	return db.DB.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Exec("UPDATE video SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
//			return err
//		}
//		if err := tx.Exec("INSERT INTO `favorite` (`user_info_id`,`video_id`) VALUES (?,?)", userId, videoId).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//}
//
//// MinusOneFavorByUserIdAndVideoId 减少一个赞
//func (v *VideoDAO) MinusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
//	return db.DB.Transaction(func(tx *gorm.DB) error {
//		//执行-1之前需要先判断是否合法（不能被减少为负数
//		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count-1 WHERE id = ? AND favorite_count>0", videoId).Error; err != nil {
//			return err
//		}
//		if err := tx.Exec("DELETE FROM `favorite`  WHERE `user_info_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//}
//
//func (v *VideoDAO) QueryFavorVideoListByUserId(userId int64, videoList *[]*Video) error {
//	if videoList == nil {
//		return errors.New("QueryFavorVideoListByUserId videoList 空指针")
//	}
//	//多表查询，左连接得到结果，再映射到数据
//	if err := db.DB.Raw("SELECT v.* FROM favorite u , videos v WHERE u.user_info_id = ? AND u.video_id = v.id", userId).Scan(videoList).Error; err != nil {
//		return err
//	}
//	//如果id为0，则说明没有查到数据
//	if len(*videoList) == 0 || (*videoList)[0].ID == 0 {
//		return errors.New("点赞列表为空")
//	}
//	return nil
//}

func (v *VideoDAO) IsVideoExistById(id int64) bool {
	var video Video
	if err := db.DB.Where("id=?", id).Select("id").First(&video).Error; err != nil {
		log.Println(err)
	}
	if video.ID == 0 {
		return false
	}
	return true
}
