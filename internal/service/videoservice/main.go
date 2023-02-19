package main

import (
	"TikTokLiteV2/internal/service/videoservice/dal"
	"TikTokLiteV2/internal/service/videoservice/handler"
	tiktok "TikTokLiteV2/kitex_gen/tiktok/videoservice"
	"log"
)

func main() {
	svr := tiktok.NewServer(new(handler.VideoServiceImpl))
	dal.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
