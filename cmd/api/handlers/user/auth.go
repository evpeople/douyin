package user

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/constants"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func init() {

	AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar UserParam

			if err := c.BindQuery(&loginVar); err != nil {
				handlers.SendBaseResponse(c, err)
			}
			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			id, err := rpc.CheckUser(context.Background(), &user.DouyinUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
			c.Set("userID", id)
			return id, err
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup:   "cookie: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			id := c.GetInt64("userID")
			c.JSON(http.StatusOK, RegisterResponse{
				Code:    int64(code),
				Message: errno.Success.ErrMsg,
				UserID:  id,
				Token:   message,
			})
		},
	})
}
