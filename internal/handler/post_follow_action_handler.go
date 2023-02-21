package handler

import (
	"context"
	"errors"
	"github.com/okatu-loli/TikTokLite/internal/model"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/okatu-loli/TikTokLite/internal/service/user_info"
)

type ProxyPostFollowAction struct {
	Context *app.RequestContext

	userId     int64
	followId   int64
	actionType int
}

func PostFollowActionHandler(ctx context.Context, c *app.RequestContext) {
	NewProxyPostFollowAction(c).Do()
}

func NewProxyPostFollowAction(c *app.RequestContext) *ProxyPostFollowAction {
	return &ProxyPostFollowAction{Context: c}
}

func (p *ProxyPostFollowAction) Do() {
	var err error
	if err = p.prepareNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.startAction(); err != nil {
		//当错误为model层发生的，那么就是重复键值的插入了
		if errors.Is(err, user_info.ErrIvdAct) || errors.Is(err, user_info.ErrIvdFolUsr) {
			p.SendError(err.Error())
		} else {
			p.SendError("请勿重复关注")
		}
		return
	}
	p.SendOk("操作成功")
}

// 封装数据
func (p *ProxyPostFollowAction) prepareNum() error {
	ctx := p.Context
	rawUserId, _ := ctx.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	//解析需要关注的id
	followId := ctx.Query("to_user_id")
	parseInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		return err
	}
	p.followId = parseInt

	//解析action_type
	actionType := ctx.Query("action_type")
	parseInt, err = strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		return err
	}
	p.actionType = int(parseInt)
	return nil
}

func (p *ProxyPostFollowAction) startAction() error {
	err := user_info.PostFollowAction(p.userId, p.followId, p.actionType) //查询数据
	if err != nil {
		return err
	}
	return nil
}

// 操作失败
func (p *ProxyPostFollowAction) SendError(msg string) {
	p.Context.JSON(http.StatusOK, model.CommonResponse{StatusCode: 1, StatusMsg: msg})
}

// 操作成功
func (p *ProxyPostFollowAction) SendOk(msg string) {
	p.Context.JSON(http.StatusOK, model.CommonResponse{StatusCode: 1, StatusMsg: msg})
}
