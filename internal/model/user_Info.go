package model

//
//var (
//	ErrIvdPtr        = errors.New("空指针错误")
//	ErrEmptyUserList = errors.New("用户列表为空")
//)

type UserInfo struct {
	//相当于之前的id
	//UserId        int64  `gorm:"Column:user_id" json:"user_id"`
	UserName      string `gorm:"Column:username" json:"username"`
	Password      string `gorm:"Column:password" json:"password"`
	FollowCount   int64  `gorm:"Column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"Column:follower_count" json:"follower_count"`
	BasePo
}

//
//type UserInfoDAO struct {
//}
//
//var (
//	userInfoDAO  *UserInfoDAO
//	userInfoOnce sync.Once
//)
//
//func NewUserInfoDAO() *UserInfoDAO {
//	userInfoOnce.Do(func() {
//		userInfoDAO = new(UserInfoDAO)
//	})
//	return userInfoDAO
//}
//
//func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
//	if userinfo == nil {
//		return ErrIvdPtr
//	}
//	//DB.Where("id=?",userId).First(userinfo)
//	db.DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
//	//id为零值，说明sql执行失败
//	if userinfo.ID == 0 {
//		return errors.New("该用户不存在")
//	}
//	return nil
//}
//
//func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
//	if userinfo == nil {
//		return ErrIvdPtr
//	}
//	return db.DB.Create(userinfo).Error
//}
//
//func (u *UserInfoDAO) IsUserExistById(id int64) bool {
//	var userinfo UserInfo
//	if err := db.DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
//		log.Println(err)
//	}
//	if userinfo.ID == 0 {
//		return false
//	}
//	return true
//}
