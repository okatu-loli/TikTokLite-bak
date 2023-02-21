package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/okatu-loli/TikTokLite/internal/request"
	"github.com/okatu-loli/TikTokLite/internal/response"
	"github.com/okatu-loli/TikTokLite/internal/service/favorite"
	"log"
	"net/http"
)

type IFavoriteController interface {
	FavoriteAction(ctx context.Context, c *app.RequestContext)    // 点赞操作
	FavoriteVideoList(ctx context.Context, c *app.RequestContext) // 获取点赞列表
}

type FavoriteController struct {
	FavoriteService favorite.IFavoriteService
}

func NewFavoriteController() IFavoriteController {
	favoriteController := FavoriteController{FavoriteService: favorite.NewFavoriteService()}
	return favoriteController
}

// FavoriteAction 点赞操作
func (f FavoriteController) FavoriteAction(ctx context.Context, c *app.RequestContext) {

	var favoriteRequest request.FavoriteActionParam
	err := c.BindAndValidate(&favoriteRequest)
	if err != nil {
		log.Printf("FavoriteAction|参数错误|%v", err)
		return
	}

	value, _ := c.Get("userId")
	f.FavoriteService.FavoriteAction(value.(int64), favoriteRequest)

	c.JSON(http.StatusOK, response.FavoriteActionResponse{
		CommonResponse: response.CommonResponse{StatusCode: 0, StatusMsg: "success"},
	})
}

// FavoriteVideoList 获取用户点赞的视频
func (f FavoriteController) FavoriteVideoList(ctx context.Context, c *app.RequestContext) {
	var videoRequest request.FavoriteListParam
	err := c.BindAndValidate(&videoRequest)
	if err != nil {
		log.Printf("ListVideo|请求参数错误|%v", err)
		c.JSON(http.StatusBadRequest, response.ErrorMessage{
			Response: response.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			},
		})
		return
	}
	video := f.FavoriteService.FavoriteVideoList(&videoRequest)
	c.JSON(http.StatusOK, video)
}
