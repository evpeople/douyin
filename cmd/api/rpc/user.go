package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/kitex_gen/user/userservice"
	"github.com/evpeople/douyin/pkg/constants"
	"github.com/evpeople/douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	// resp, err := userClient.CreateUser(ctx, req)
	resp, err := userClient.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResponse.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResponse.StatusCode), *resp.BaseResponse.StatusMsg)
	}
	return resp, nil
}

// CheckUser check user info
//TODO:似乎有些错误
func CheckUser(ctx context.Context, req *user.DouyinUserRequest) (int64, error) {
	// resp, err := userClient.CheckUser(ctx, req)
	resp, err := userClient.LoginUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResponse.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.BaseResponse.StatusCode), *resp.BaseResponse.StatusMsg)
		// return 0, errno.NewErrNo(resp.BaseResponse.StatusCode, resp.BaseResponse.)
	}
	return resp.BaseMessage.UserId, nil
}
func GetUser(ctx context.Context, req *user.DouyinUserMessageRequest) (*user.User, error) {
	resp, err := userClient.GetUser(ctx, req)

	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), *resp.BaseResp.StatusMsg)
	}
	return resp.User, nil
}
