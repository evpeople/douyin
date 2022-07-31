package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/evpeople/douyin/kitex_gen/publish"
	"github.com/evpeople/douyin/kitex_gen/publish/publishservice"
	"github.com/evpeople/douyin/pkg/constants"
	"github.com/evpeople/douyin/pkg/errno"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var publishClient publishservice.Client

func initPublishRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
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
	publishClient = c
}

func UploadVideo(ctx context.Context, req *publish.UploadFileRequest) (*publish.BaseResponse, error) {
	// resp, err := userClient.CreateUser(ctx, req)
	// resp, err := userClient.RegisterUser(ctx, req)
	resp, err := publishClient.PostVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
func GetVideos(ctx context.Context, auth bool, req *publish.DouyinPublishRequest) (resp *publish.DouyinPublishResponse, err error) {
	if auth {
		resp, err = publishClient.GetPublishVideos(ctx, req)
	} else {
		resp, err = publishClient.GetVideos(ctx, req)
	}
	return
}
