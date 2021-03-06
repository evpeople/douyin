// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/evpeople/douyin/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RegisterUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
	LoginUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
	GetUser(ctx context.Context, Req *user.DouyinUserMessageRequest, callOptions ...callopt.Option) (r *user.DouyinUesrMessageResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) RegisterUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterUser(ctx, Req)
}

func (p *kUserServiceClient) LoginUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginUser(ctx, Req)
}

func (p *kUserServiceClient) GetUser(ctx context.Context, Req *user.DouyinUserMessageRequest, callOptions ...callopt.Option) (r *user.DouyinUesrMessageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}
