package handlers

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/constants"
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
	// logger.Debug("tetststs", resp, *resp)
	if err != nil {
		logger.Debug(err)
		SendRegisterResponse(c, errno.ConvertErr(err), nil)
		return
	}
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {

				return jwt.MapClaims{
					// constants.IdentityKey: v,
					"ID": v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVar UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				logger.Debug("noBind")
				logger.Debug(err)
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				logger.Debug(loginVar)
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &user.DouyinUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.JSON(http.StatusOK, RegisterResponse{
				Token: message,
			})
			// logger.Debug(c.Get("JWT_PAYLOAD"))
		},
	})
	token, _, _ := authMiddleware.TokenGenerator(resp.BaseMessage.UserId)
	SendRegisterResponse(c, errno.Success, &UserResp{resp.BaseMessage.UserId, token})
	// SendRegisterResponse(c, errno.Success, &UserResp{2123, "tete"})
}
