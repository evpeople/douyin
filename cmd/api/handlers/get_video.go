package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/feed"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// userID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		LastestTime string `json:"lastest_time" form:"lastest_time"`
		Token       string `json:"token" form:"token"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		SendBaseResponse(c, errno.ConvertErr(err))
	}
	lst, err := strconv.ParseInt(queryVar.LastestTime, 10, 0)
	if err != nil && queryVar.LastestTime != "" {
		logger.Debug(err)
		SendBaseResponse(c, errno.ConvertErr(err))
		return
	} else {
		queryVar.LastestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	req := &feed.DouyinFeedRequest{
		LastestTime: &lst,
		Token:       &queryVar.Token,
	}

	videos, err := rpc.GetVideos(context.Background(), req)
	logger.Debug("getVideo", videos)
	// notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendBaseResponse(c, errno.ConvertErr(err))
		return
	}
	// SendResponse(c, errno.Success, map[string]interface{}{constants.Total: total, constants.Notes: notes})
	sendFeedResponse(c, err, videos)
}

// SendRegisterResponse pack response
func sendFeedResponse(c *gin.Context, err error, data []*feed.Video) {
	Err := errno.ConvertErr(err)
	if data == nil {
		c.JSON(http.StatusOK, RegisterResponse{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, feed.Video{
		Id: 0,
		Author: &feed.User{
			Id:            0,
			Name:          "",
			FollowCount:   new(int64),
			FollowerCount: new(int64),
			IsFollow:      false,
		},
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "test video 1",
	})
}
