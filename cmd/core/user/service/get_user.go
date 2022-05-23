package service

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/cmd/core/user/pack"
	"github.com/evpeople/douyin/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *GetUserService) MGetUser(req *user.DouyinUserMessageRequest) (*user.User, error) {
	modelUser, err := db.MGetUser(s.ctx, req.BaseResp.UserId)
	if err != nil {
		return nil, err
	}
	return pack.User(modelUser), nil
}
