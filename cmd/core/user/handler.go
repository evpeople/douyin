package main

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/user/pack"
	"github.com/evpeople/douyin/cmd/core/user/service"
	"github.com/evpeople/douyin/kitex_gen/user"
	"github.com/evpeople/douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResponse = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResponse = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResponse = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResponse = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResponse = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseMessage.UserId = uid
	resp.BaseResponse = pack.BuildBaseResp(errno.Success)
	return resp, nil
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserMessageRequest) (resp *user.DouyinUesrMessageResponse, err error) {
	resp = new(user.DouyinUesrMessageResponse)

	// if len(req.UserIds) == 0 {
	// 	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	// 	return resp, nil
	// }
	users, err := service.NewMGetUserService(ctx).MGetUser(req)

	// users, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = users
	return resp, nil
}
