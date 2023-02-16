package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
	"github.com/okatu-loli/TikTokLite/internal/middleware"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service"
	"github.com/okatu-loli/TikTokLite/internal/service/util"
)

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	statusCode := 0   // 状态码
	statusMsg := ""   // 返回状态
	tokenString := "" // 鉴权token
	userId := -1

	var user *model.User

	var err error

	if err = util.UserInfoCheck(username, password); err != nil { // 检查传入的用户名和密码是否合法
		statusCode = -1
		statusMsg = err.Error()
	} else {
		// MD5编码密码
		password = util.Md5Encode(password)
		// 注册用户
		if user, err = service.Register(username, password); err != nil {
			statusCode = -1
			statusMsg = err.Error()
			fmt.Println(statusMsg)
		} else { // 注册成功
			statusMsg = "Register Success!"
			userId = int(user.ID)
			// 签发token
			token := jwt.New(jwt.GetSigningMethod(middleware.JwtMiddleware.SigningAlgorithm))
			claims := token.Claims.(jwt.MapClaims)
			claims[middleware.JwtMiddleware.IdentityKey] = user.ID
			expire := middleware.JwtMiddleware.TimeFunc().Add(middleware.JwtMiddleware.Timeout)
			claims["exp"] = expire.Unix()
			claims["orig_iat"] = middleware.JwtMiddleware.TimeFunc().Unix()
			tokenString, _ = token.SignedString(middleware.JwtMiddleware.Key)
		}
	}

	c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"user_id":     userId,
		"token":       tokenString,
	})
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	id := c.Query("user_id")
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态
	name := ""
	follow_count := 0
	follower_count := 0
	is_follow := false // 暂时还没写
	user, err := service.FindUserById(id)
	if err != nil {
		statusCode = -1
		statusMsg = err.Error()
	} else {
		name = user.UserName
		follow_count = user.FollowCount
		follower_count = user.FollowerCount

	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"user": utils.H{
			"id":             id,
			"name":           name,
			"follow_count":   follow_count,
			"follower_count": follower_count,
			"is_follow":      is_follow,
		},
	})
}
