package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *user.DouyinUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	user, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errno.UserNotExistErr
	}
	// u := users[0]
	if user.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(user.ID), nil
}
