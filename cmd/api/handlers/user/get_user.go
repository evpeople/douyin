package user

import (
	"context"
	"net/http"

	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

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

func GetUser(c *gin.Context) {
	var queryVar struct {
		UserID int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}

	if err := c.BindQuery(&queryVar); err != nil {
		handlers.SendBaseResponse(c, err)
	}
	req := &user.DouyinUserMessageRequest{BaseResp: &user.BaseMessage{UserId: queryVar.UserID, Token: queryVar.Token}}
	user, err := rpc.GetUser(context.Background(), req)
	if err != nil {
		handlers.SendBaseResponse(c, err)
		return
	}
	sendUserResponse(c, errno.Success, user)
}

func sendUserResponse(c *gin.Context, err error, data *user.User) {
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
