package handlers

import (
	"context"
	"log"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		log.Println(registerVar)
		SendRegisterResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := rpc.CreateUser(context.Background(), &user.DouyinUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	// logger.Debug("tetststs", resp, *resp)
	if err != nil {
		logger.Debug(err)
		SendRegisterResponse(c, errno.ConvertErr(err), nil)
		return
	}
	token, _, _ := AuthMiddleware.TokenGenerator(resp.BaseMessage.UserId)
	SendRegisterResponse(c, errno.Success, &UserResp{resp.BaseMessage.UserId, token})
	// SendRegisterResponse(c, errno.Success, &UserResp{2123, "tete"})
}
