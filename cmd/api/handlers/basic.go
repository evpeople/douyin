package handlers

import (
	"net/http"

	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func SendBaseResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	resp := struct {
		Code    int64  `json:"status_code"`
		Message string `json:"status_msg"`
	}{Err.ErrCode, Err.ErrMsg}
	c.JSON(http.StatusOK,
		resp)
}
