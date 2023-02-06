package response

type User struct {
	UserId        int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type InfoResponse struct {
	CommonResponse
	User User `json:"user"`
}

type LoginResponse struct {
	CommonResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterResponse struct {
	CommonResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
