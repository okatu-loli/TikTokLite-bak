package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"github.com/okatu-loli/TikTokLite/cmd/dal/db"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/util"
	"gorm.io/gorm"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	//用于标识身份的key
	IdentityKey = "id"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		//作用域
		Realm: "tiktok zone",
		//签名密钥
		Key: []byte("secret key"),
		//过期时间
		Timeout: time.Hour,
		//token的获取源
		TokenLookup: "query:token, header: Authorization",
		//最大token刷新时间
		MaxRefresh: time.Hour,
		//登录的相应函数，ps在内置的登陆函数“LoginHandler”中，这个是最后一个被调用的函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			id, _ := c.Get("id")
			c.JSON(http.StatusOK, utils.H{
				"code":       code,
				"token":      token,
				"user_id":    id,
				"status_msg": "登录成功",
			})
		},
		//登录验证鉴权
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				UserName string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := db.CheckUser(loginStruct.UserName, util.Md5Encode(loginStruct.Password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}
			c.Set("id", users[0].ID)
			return users[0], nil
		},
		IdentityKey: IdentityKey,
		//获取身份信息的函数
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			id := claims[IdentityKey].(float64)
			c.Set("user_id", id)
			return &model.User{
				Model: gorm.Model{ID: uint(id)},
			}
		},
		//给token自定义的添加负载信息，这里存了id
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				fmt.Println("存入成功", v.ID)
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		//流程失败的错误处理
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		//设置 jwt 授权失败后的响应函数
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
