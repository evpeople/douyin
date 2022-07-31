package main

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/video/service"
	"github.com/evpeople/douyin/kitex_gen/publish"
	"github.com/evpeople/douyin/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// GetPublishVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishVideos(ctx context.Context, req *publish.DouyinPublishRequest) (resp *publish.DouyinPublishResponse, err error) {
	// TODO: Your code here...
	return
}

// PostVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PostVideos(ctx context.Context, req *publish.UploadFileRequest) (resp *publish.BaseResponse, err error) {

	err = service.NewPostVideosService(ctx).PostVideos(req)
	nerr := errno.ConvertErr(err)
	return &publish.BaseResponse{StatusMsg: &nerr.ErrMsg, StatusCode: int32(nerr.ErrCode)}, err
}

// GetVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetVideos(ctx context.Context, req *publish.DouyinPublishRequest) (resp *publish.DouyinPublishResponse, err error) {
	// TODO: Your code here...
	return
}
