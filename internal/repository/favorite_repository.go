package repository

import (
	"errors"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	PlusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error          // 增加点赞记录
	MinusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error         // 取消点赞记录
	QueryFavorVideoListByUserId(userId int64, videoList *[]*model.Video) error //查询喜欢列表
	CheckIsLike(userid int64, id2 int64) (bool, error)
}

// FavoriteRepository UserRepository 定义一个结构体
type FavoriteRepository struct {
}

// CheckIsLike 检查视频是否被点赞
func (f FavoriteRepository) CheckIsLike(userId int64, videoId int64) (bool, error) {
	var count int64 = 0
	err := db.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error
	if count == 0 {
		return false, err
	}
	return true, err
}

// PlusOneFavorByUserIdAndVideoId 增加点赞记录
func (f FavoriteRepository) PlusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE video SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `favorite` (`user_id`,`video_id`) VALUES (?,?)", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

// MinusOneFavorByUserIdAndVideoId 减少一个赞
func (f FavoriteRepository) MinusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		//执行-1之前需要先判断是否合法（不能被减少为负数
		if err := tx.Exec("UPDATE video SET favorite_count=favorite_count-1 WHERE id = ? AND favorite_count>0", videoId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `favorite`  WHERE `user_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

func (f FavoriteRepository) QueryFavorVideoListByUserId(userId int64, videoList *[]*model.Video) error {
	if videoList == nil {
		return errors.New("QueryFavorVideoListByUserId videoList 空指针")
	}
	//多表查询，左连接得到结果，再映射到数据
	if err := db.DB.Raw("SELECT v.* FROM favorite u , video v WHERE u.user_id = ? AND u.video_id = v.id", userId).Scan(videoList).Error; err != nil {
		return err
	}
	//如果id为0，则说明没有查到数据
	if len(*videoList) == 0 || (*videoList)[0].ID == 0 {
		return errors.New("点赞列表为空")
	}
	return nil
}

// NewFavoriteRepository UserRepository构造函数
func NewFavoriteRepository() IFavoriteRepository {
	return FavoriteRepository{}
}
