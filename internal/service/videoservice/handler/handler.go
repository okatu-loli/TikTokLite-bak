package handler

import (
	tiktok "TikTokLiteV2/kitex_gen/tiktok"
	"context"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, request *tiktok.FeedRequest) (resp *tiktok.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, request *tiktok.PublishActionRequest) (resp *tiktok.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, request *tiktok.PublishListRequest) (resp *tiktok.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
