package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
	"gorm.io/gorm"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.DouyinUserRequest) (uint, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Debug("test create")
		h := md5.New()
		if _, err = io.WriteString(h, req.Password); err != nil {
			return 0, err
		}
		passWord := fmt.Sprintf("%x", h.Sum(nil))
		ur := []*db.User{{
			UserName: req.Username,
			Password: passWord,
		}}
		err := db.CreateUser(s.ctx, ur)
		return ur[0].ID, err
	} else {
		return users.ID, errno.UserAlreadyExistErr
	}

}
