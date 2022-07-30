package main

import (
	"context"
	"github.com/evpeople/douyin/idl/kitex_gen/publish"
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
	// TODO: Your code here...
	return
}

// GetVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetVideos(ctx context.Context, req *publish.DouyinPublishRequest) (resp *publish.DouyinPublishResponse, err error) {
	// TODO: Your code here...
	return
}
