package main

import (
	"context"
	"github.com/evpeople/douyin/kitex_gen/publish"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetPublishVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideos(ctx context.Context, req *publish.DouyinPublishRequest) (resp *publish.DouyinPublishResponse, err error) {
	// TODO: Your code here...
	return
}

// PostVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PostVideos(ctx context.Context, req *publish.UploadFileRequest) (resp *publish.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
