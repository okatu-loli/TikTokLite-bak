// Code generated by hertz generator.

package router

import (
	handler "github.com/TikTokLite/internal/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func CustomizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
