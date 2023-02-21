package user_info

import (
	"github.com/okatu-loli/TikTokLite/cache"
	"github.com/okatu-loli/TikTokLite/internal/model"
)

type FollowerList struct {
	UserList []*model.UserInfo `json:"user_list"`
}

func QueryFollowerList(userId int64) (*FollowerList, error) {
	return NewQueryFollowerListFlow(userId).Do()
}

type QueryFollowerListFlow struct {
	userId int64

	userList []*model.UserInfo

	*FollowerList
}

func NewQueryFollowerListFlow(userId int64) *QueryFollowerListFlow {
	return &QueryFollowerListFlow{userId: userId}
}

func (q *QueryFollowerListFlow) Do() (*FollowerList, error) {
	var err error
	if err = q.checkNum(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}
	return q.FollowerList, nil
}

// 检查id是否存在
func (q *QueryFollowerListFlow) checkNum() error {
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

// 查询数据
func (q *QueryFollowerListFlow) prepareData() error {

	err := model.NewUserInfoDAO().GetFollowerListByUserId(q.userId, &q.userList)
	if err != nil {
		return err
	}
	// 填充is_follow字段
	for _, v := range q.userList {
		v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(q.userId, v.Id)
	}
	return nil
}

// 封装数据
func (q *QueryFollowerListFlow) packData() error {
	q.FollowerList = &FollowerList{UserList: q.userList}

	return nil
}
