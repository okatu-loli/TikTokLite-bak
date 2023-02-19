// Code generated by Kitex v0.4.4. DO NOT EDIT.
package testservice

import (
	server "github.com/cloudwego/kitex/server"
	tiktok "TikTokLiteV2/kitex_gen/tiktok"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler tiktok.TestService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}