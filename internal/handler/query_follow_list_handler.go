package handler

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/user_info"
	"net/http"
)

type FollowListResponse struct {
	model.CommonResponse
	FollowList *user_info.FollowList
}

func QueryFollowListHandler(ctx context.Context, c *app.RequestContext) {
	NewProxyQueryFollowList(c).Do()
}

type ProxyQueryFollowList struct {
	Ctx *app.RequestContext

	userId int64

	FollowList *user_info.FollowList
}

func NewProxyQueryFollowList(c *app.RequestContext) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Ctx: c}
}

func (p *ProxyQueryFollowList) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendError(err.Error())
		return
	}
	p.SendOk("请求成功")
}

// 解析id
func (p *ProxyQueryFollowList) parseNum() error {
	rawUserId, _ := p.Ctx.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowList) prepareData() error {
	list, err := user_info.QueryFollowList(p.userId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}

func (p *ProxyQueryFollowList) SendError(msg string) {
	p.Ctx.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: model.CommonResponse{StatusCode: 1, StatusMsg: msg},
	})
}

func (p *ProxyQueryFollowList) SendOk(msg string) {
	p.Ctx.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: model.CommonResponse{StatusCode: 0, StatusMsg: msg},
		FollowList:     p.FollowList,
	})
}
