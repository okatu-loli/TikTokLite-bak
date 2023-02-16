package main

import (
	tiktok "TikTokLiteV2/kitex_gen/tiktok"
	"TikTokLiteV2/kitex_gen/tiktok/testservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

func main() {
	client, err := testservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &tiktok.TestRequest{
			Name: "这是客户端的测试信息",
		}
		resp, err := client.Test(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)

	}
}
