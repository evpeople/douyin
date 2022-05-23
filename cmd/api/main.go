package main

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/constants"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/evpeople/douyin/pkg/tracer"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
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
			var loginVar handlers.UserParam
			if err := c.ShouldBind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			id, err := rpc.CheckUser(context.Background(), &user.DouyinUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
			c.Set("userID", id)
			return id, err
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			// token := jwt.ExtractClaimsFromToken(jwt.GetToken(c))
			id := c.GetInt64("userID")
			c.JSON(http.StatusOK, handlers.RegisterResponse{
				Code:    int64(code),
				Message: errno.Success.ErrMsg,
				UserID:  id,
				Token:   message,
			})
			// logger.Debug(c.Get("JWT_PAYLOAD"))
		},
	})

	v1 := r.Group("/douyin")
	user1 := v1.Group("/user")
	user1.GET("", handlers.GetUser)
	// authMiddl
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/register", handlers.Register)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
