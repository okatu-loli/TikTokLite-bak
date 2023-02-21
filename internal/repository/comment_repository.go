package repository

import (
	"errors"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	InsertComment(comment *model.Comment) error                       //添加评论
	QueryCommentListByVideoId(videoId int64) ([]model.Comment, error) // 查看评论列表
	DeleteCommentAndUpdateCountById(commentId, videoId int64) error   //删除评论
	QueryCommentById(id int64, comment *model.Comment) error
}

type CommentRepository struct {
}

// QueryCommentListByVideoId 获取评论列表
func (c CommentRepository) QueryCommentListByVideoId(videoId int64) ([]model.Comment, error) {
	var commentList []model.Comment
	if commentList == nil {
		return nil, errors.New("QueryCommentListByVideoId comments空指针")
	}
	err := db.DB.Model(&model.Comment{}).Where("video_id=?", videoId).Find(commentList).Error
	return commentList, err

	//return nil ,err
}

// InsertComment 添加评论
func (c CommentRepository) InsertComment(comment *model.Comment) error {
	//添加评论并且更新数量
	if comment == nil {
		return errors.New("InsertComment comment空指针")
	}
	//执行事务
	return db.DB.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//增加count
		if err := tx.Exec("UPDATE video v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.ID).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

}

// DeleteCommentAndUpdateCountById 删除评论   这个ID 指的是video_id
func (c CommentRepository) DeleteCommentAndUpdateCountById(commentId, ID int64) error {
	//执行事务
	return db.DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM comment WHERE id = ?", commentId).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE video v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", ID).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

}

func (c CommentRepository) QueryCommentById(id int64, comment *model.Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	return db.DB.Where("id=?", id).First(comment).Error
}

func NewCommentRepository() ICommentRepository {
	return CommentRepository{}
}
