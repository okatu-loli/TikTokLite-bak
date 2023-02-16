package handler

import (
	tiktok "TikTokLiteV2/kitex_gen/tiktok"
	"context"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// Test implements the TestServiceImpl interface.
func (s *TestServiceImpl) Test(ctx context.Context, request *tiktok.TestRequest) (resp *tiktok.TestResponse, err error) {
	// TODO: Your code here...
	resp = &tiktok.TestResponse{
		Message: request.Name,
	}
	return
}
