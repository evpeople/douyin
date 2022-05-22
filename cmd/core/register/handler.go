package main

import (
	"context"
	"github.com/evpeople/douyin/kitex_gen/rigister"
)

// RegisterServiceImpl implements the last service interface defined in the IDL.
type RegisterServiceImpl struct{}

// RegisterUser implements the RegisterServiceImpl interface.
func (s *RegisterServiceImpl) RegisterUser(ctx context.Context, req *rigister.DouyinUserRegisterRequest) (resp *rigister.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}
