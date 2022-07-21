// Code generated by Kitex v0.3.4. DO NOT EDIT.
package videoservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/evpeople/douyin/kitex_gen/feed"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler feed.VideoService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
