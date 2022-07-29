package main

import (
	"log"

	publish "github.com/evpeople/douyin/kitex_gen/publish/videoservice"
)

func main() {
	svr := publish.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
