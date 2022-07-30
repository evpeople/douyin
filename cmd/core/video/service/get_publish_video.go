package service

import (
	"context"

	"github.com/evpeople/douyin/kitex_gen/publish"
	"github.com/evpeople/douyin/pkg/errno"
)

type GetPublishVideosService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetPublishVideosService(ctx context.Context) *GetPublishVideosService {
	return &GetPublishVideosService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *GetPublishVideosService) MGetVideos(req *publish.DouyinPublishRequest) (*publish.Video, error) {
	// modelUser, err := db.MGetUser(s.ctx, req.BaseResp.UserId)
	// if err != nil {
	return nil, errno.LoginErr
	// }
	// return pack.User(modelUser), nil
}
