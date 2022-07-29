package handlers

import (
	"net/http"

	"github.com/evpeople/douyin/cmd/core/user/pack"
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

// SendRegisterResponse pack response
func SendRegisterResponse(c *gin.Context, err error, data *UserResp) {
	Err := errno.ConvertErr(err)
	if data == nil {
		c.JSON(http.StatusOK, RegisterResponse{
			Code:    Err.ErrCode,
			Message: Err.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		UserID:  data.UserID,
		Token:   data.Token,
	})
}

func SendBaseResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, RegisterResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
	return
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserResp struct {
	UserID int64
	Token  string
}

type UserResponse struct {
	// Code    int64  `json:"status_code"`
	// Message string `json:"status_msg"`
	// User  `json:"user_id"`
	user.DouyinUesrMessageResponse
}

func SendUserResponse(c *gin.Context, err error, data *user.User) {
	// Err := errno.ConvertErr(err)
	var resp user.DouyinUesrMessageResponse
	resp.BaseResponse = pack.BuildBaseResp(err)
	resp.User = data
	c.JSON(http.StatusOK,
		resp)
}
