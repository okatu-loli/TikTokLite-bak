package user_info

import (
	"errors"

	"github.com/okatu-loli/TikTokLite/cache"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

const (
	FOLLOW = 1
	CANCEL = 2
)

var (
	ErrIvdAct    = errors.New("未定义操作")
	ErrIvdFolUsr = errors.New("关注用户不存在")
)

func PostFollowAction(userId, userToId int64, actionType int) error {
	return NewPostFollowActionFlow(userId, userToId, actionType).Do()
}

type PostFollowActionFlow struct {
	userId     int64
	userToId   int64
	actionType int
}

func NewPostFollowActionFlow(userId int64, userToId int64, actionType int) *PostFollowActionFlow {
	return &PostFollowActionFlow{userId: userId, userToId: userToId, actionType: actionType}
}

func (p *PostFollowActionFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}
	if err = p.publish(); err != nil {
		return err
	}
	return nil
}

// 检查关注是否合法
func (p *PostFollowActionFlow) checkNum() error {
	//判断关注用户是否存在
	if !model.NewUserInfoDAO().IsUserExistById(p.userToId) {
		return ErrIvdFolUsr
	}
	//操作类型错误
	if p.actionType != FOLLOW && p.actionType != CANCEL {
		return ErrIvdAct
	}
	//不能关注自己
	if p.userId == p.userToId {
		return ErrIvdAct
	}
	return nil
}

func (p *PostFollowActionFlow) publish() error {
	userDAO := model.NewUserInfoDAO()
	var err error
	switch p.actionType {
	case FOLLOW:
		err = userDAO.AddUserFollow(p.userId, p.userToId)
		//更新redis的关注信息
		cache.NewProxyIndexMap().UpdateUserRelation(p.userId, p.userToId, true)
	case CANCEL:
		err = userDAO.CancelUserFollow(p.userId, p.userToId)
		cache.NewProxyIndexMap().UpdateUserRelation(p.userId, p.userToId, false)
	default:
		return ErrIvdAct
	}
	return err
}
