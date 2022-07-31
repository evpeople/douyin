package video

import (
	"context"
	"net/http"
	"time"

	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/publish"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type VideoResp struct {
	Code    int64            `json:"status_code"`
	Message string           `json:"status_msg"`
	Video   []*publish.Video `json:"video_list"`
}

func GetVideos(c *gin.Context) {
	now := time.Now().Unix()
	resp, err := rpc.GetVideos(context.TODO(), false, &publish.DouyinPublishRequest{
		UserId:      new(int64),
		LastestTime: &now,
	})
	sendResponse(c, err, resp)
}

func GetPublishVideos(c *gin.Context) {
	userID := handlers.GetIdFromRequest(c)
	resp, err := rpc.GetVideos(context.TODO(), true, &publish.DouyinPublishRequest{
		UserId:      &userID,
		LastestTime: new(int64),
	})
	sendResponse(c, err, resp)
}

func sendResponse(c *gin.Context, err error, data *publish.DouyinPublishResponse) {
	Errno := errno.ConvertErr(err)
	if data == nil || err != nil {
		handlers.SendBaseResponse(c, err)
	} else {
		c.JSON(http.StatusOK, VideoResp{
			Errno.ErrCode,
			Errno.ErrMsg,
			data.VideoList,
		})
	}
}
