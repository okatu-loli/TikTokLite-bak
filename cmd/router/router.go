// Code generated by hertz generator.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/okatu-loli/TikTokLite/internal/handler"
)

// CustomizedRegister registers customize routers.
func CustomizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	douyin := r.Group("/douyin")
	{
		//使用中间件
		//douyin.Use(basic_auth.BasicAuth(map[string]string{"test": "test"}))
		user := douyin.Group("/user")
		{
			user.POST("/register", handler.Register)
			//user.POST("/login", handler.Login)
			//user.GET("/", handler.GetUserInfo)
		}
	}
}
