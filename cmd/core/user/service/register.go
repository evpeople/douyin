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

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.DouyinUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if users != nil {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		Password: passWord,
	}})
}
