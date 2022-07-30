package user

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type RegisterResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	UserID  int64  `json:"user_id"`
	Token   string `json:"token"`
}

type UserParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.BindQuery(&registerVar); err != nil {
		handlers.SendBaseResponse(c, err)
	}
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		log.Println(registerVar)
		handlers.SendBaseResponse(c, errno.ParamErr)
		return
	}

	resp, err := rpc.CreateUser(context.Background(), &user.DouyinUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		logger.Debug(err)
		handlers.SendBaseResponse(c, err)
		return
	}
	token, _, _ := AuthMiddleware.TokenGenerator(resp.BaseMessage.UserId)
	sendRegisterResponse(c, errno.Success, &RegisterResponse{UserID: resp.BaseMessage.UserId, Token: token})
}

func sendRegisterResponse(c *gin.Context, err error, data *RegisterResponse) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RegisterResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserID:  data.UserID,
		Token:   data.Token,
	})
}
