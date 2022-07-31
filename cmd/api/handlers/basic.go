package handlers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
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
func GetIdFromRequest(c *gin.Context) int64 {

	return int64(jwt.ExtractClaims(c)["id"].(float64))
}
