package comment

import (
	"errors"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

const (
	PLUS  = 1
	MINUS = 2
)

func PostFavorState(userId, videoId, actionType int64) error {
	return NewPostFavorStateFlow(userId, videoId, actionType).Do()
}

type PostFavorStateFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

func NewPostFavorStateFlow(userId, videoId, action int64) *PostFavorStateFlow {
	return &PostFavorStateFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: action,
	}
}

func (p *PostFavorStateFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}

	switch p.actionType {
	case PLUS:
		err = p.PlusOperation()
	case MINUS:
		err = p.MinusOperation()
	default:
		return errors.New("未定义的操作")
	}
	return err
}

// PlusOperation 点赞操作
func (p *PostFavorStateFlow) PlusOperation() error {
	//视频点赞数目+1
	err := model.NewVideoDAO().PlusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("不要重复点赞")
	}
	//后期可以通过redis缓存查看 对应的用户是否点赞的映射状态更新

	return nil
}

// MinusOperation 取消点赞
func (p *PostFavorStateFlow) MinusOperation() error {
	//视频点赞数目-1
	err := model.NewVideoDAO().MinusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("点赞数目已经为0")
	}
	//后期可以通过redis缓存查看 对应的用户是否点赞的映射状态更新

	return nil
}

func (p *PostFavorStateFlow) checkNum() error {

	if p.actionType != PLUS && p.actionType != MINUS {
		return errors.New("未定义的行为")
	}
	return nil
}
