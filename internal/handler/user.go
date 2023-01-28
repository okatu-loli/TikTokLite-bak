package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/okatu-loli/TikTokLite/internal/service"
	"github.com/okatu-loli/TikTokLite/internal/service/util"
)

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	statusCode := 0 // 状态码
	statusMsg := "" // 返回状态
	token := ""     // 鉴权token
	var err error

	if err = util.UserInfoCheck(username, password); err != nil { // 检查传入的用户名和密码是否合法
		statusCode = -1
		statusMsg = err.Error()
		return // 如果不合法则返回
	}

	// MD5编码密码
	password = util.Md5Encode(password)

	// 注册用户
	if err = service.Register(username, password); err != nil {
		statusCode = -1
		statusMsg = err.Error()

	} else { // 注册成功
		statusMsg = "Register Success!"
	}

	// 签发token

	defer c.JSON(consts.StatusOK, utils.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"user_id":     username,
		"token":       token,
	})
}
