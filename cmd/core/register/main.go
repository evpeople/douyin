package main

import (
	rigister "github.com/evpeople/douyin/kitex_gen/rigister/registerservice"
	"log"
)

func main() {
	svr := rigister.NewServer(new(RegisterServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
