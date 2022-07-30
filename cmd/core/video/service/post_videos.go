package service

import (
	"context"

	"github.com/evpeople/douyin/kitex_gen/publish"
	"github.com/evpeople/douyin/pkg/errno"
)

type PostVideosService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewPostVideosService(ctx context.Context) *PostVideosService {
	return &PostVideosService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *PostVideosService) PostVideos(req *publish.DouyinPublishRequest) (*publish.Video, error) {
	return nil, errno.LoginErr
}
