// Code generated by hertz generator.

package main

import (
	"time"

	"github.com/hertz-contrib/cors"
	"github.com/okatu-loli/TikTokLite/cmd/dal"
	"github.com/okatu-loli/TikTokLite/cmd/router"
	"github.com/okatu-loli/TikTokLite/internal/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

type Text struct {
	Val string
}

func main() {
	h := server.Default()
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	middleware.InitJwt()
	router.CustomizedRegister(h)
	dal.Init()

	h.Spin()
}
