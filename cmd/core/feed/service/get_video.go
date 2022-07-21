package service

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/feed/dal/db"
	"github.com/evpeople/douyin/cmd/core/feed/pack"
	"github.com/evpeople/douyin/kitex_gen/feed"
)

type GetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *GetUserService) MGetVideo(req *feed.DouyinFeedRequest) ([]*feed.Video, error) {
	modelUser, err := db.MGetVideo(s.ctx, *req.LastestTime)
	if err != nil {
		return nil, err
	}
	return pack.Videos(modelUser), nil
}
