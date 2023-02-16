package main

import (
	"TikTokLiteV2/handler"
	tiktok "TikTokLiteV2/kitex_gen/tiktok/testservice"
	"log"
)

func main() {
	svr := tiktok.NewServer(new(handler.TestServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
