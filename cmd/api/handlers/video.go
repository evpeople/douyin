package handlers

import (
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/evpeople/douyin/pkg/cos"
	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		klog.Debug(err)
	}
	src, err := file.Open()
	if err != nil {
		log.Println(err)
	}
	defer src.Close()
	url, err := cos.UploadVideo(c.PostForm("title"), src)
	log.Println(url)
	SendBaseResponse(c, err)
}
