package handlers

import (
	"context"

	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendRegisterResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := rpc.CreateUser(context.Background(), &user.DouyinUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendRegisterResponse(c, errno.Success, &UserResp{resp.BaseMessage.UserId, resp.BaseMessage.Token})
}
