package handlers

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/constants"
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
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	User    User   `json:"user"`
}
type User struct {
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	FollowCount   *int64 `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3,oneof" json:"follow_count"`
	FollowerCount *int64 `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3,oneof" json:"follower_count"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow"`
}

func SendUserResponse(c *gin.Context, err error, data *user.User) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		User: User{
			Id:            data.Id,
			Name:          data.Name,
			FollowCount:   data.FollowCount,
			FollowerCount: data.FollowerCount,
			IsFollow:      data.IsFollow,
		},
	})
}

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

			loginVar.UserName = c.Query("username")
			loginVar.PassWord = c.Query("password")
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
