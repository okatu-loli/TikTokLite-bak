package handler

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"github.com/okatu-loli/TikTokLite/internal/service/user_info"
	"net/http"
)

type FollowerListResponse struct {
	model.CommonResponse
	FollowerList *user_info.FollowerList
}

func QueryFollowerHandler(ctx context.Context, c *app.RequestContext) {
	NewProxyQueryFollowerHandler(c).Do()
}

type ProxyQueryFollowerHandler struct {
	Ctx *app.RequestContext

	userId int64

	FollowerList *user_info.FollowerList
}

func NewProxyQueryFollowerHandler(c *app.RequestContext) *ProxyQueryFollowerHandler {
	return &ProxyQueryFollowerHandler{Ctx: c}
}

func (p *ProxyQueryFollowerHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, user_info.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("准备数据出错")
		}
		return
	}
	p.SendOk("成功")
}

// 解析id
func (p *ProxyQueryFollowerHandler) parseNum() error {
	rawUserId, _ := p.Ctx.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

// 查询数据
func (p *ProxyQueryFollowerHandler) prepareData() error {
	list, err := user_info.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

func (p *ProxyQueryFollowerHandler) SendError(msg string) {
	p.Ctx.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryFollowerHandler) SendOk(msg string) {
	p.Ctx.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
		FollowerList: p.FollowerList,
	})
}
