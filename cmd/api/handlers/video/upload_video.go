package video

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/publish"
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
	title := c.PostForm("title")
	vurl, curl, err := cos.UploadVideo(title, src)
	if err != nil {
		log.Println(err)
	}
	_, err = rpc.UploadVideo(context.Background(), &publish.UploadFileRequest{
		Title:    title,
		VideoUrl: vurl,
		CoverUrl: curl,
		AuthorId: handlers.GetIdFromRequest(c),
	})

	handlers.SendBaseResponse(c, err)
}
