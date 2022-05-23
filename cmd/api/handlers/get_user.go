package handlers

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// userID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		UserID int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), nil)
	}

	req := &user.DouyinUserMessageRequest{BaseResp: &user.BaseMessage{UserId: queryVar.UserID, Token: queryVar.Token}}
	// req := &notedemo.QueryNoteRequest{UserId: userID, Offset: queryVar.Offset, Limit: queryVar.Limit}
	// if len(queryVar.SearchKeyword) != 0 {
	// 	req.SearchKey = &queryVar.SearchKeyword
	// }
	user, err := rpc.GetUser(context.Background(), req)
	logger.Debug("getUser", user)
	// notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendUserResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// SendResponse(c, errno.Success, map[string]interface{}{constants.Total: total, constants.Notes: notes})
	SendUserResponse(c, errno.Success, user)
}
