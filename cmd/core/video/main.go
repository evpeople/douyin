package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	userdb "github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/cmd/core/video/dal"
	publish "github.com/evpeople/douyin/kitex_gen/publish/publishservice"
	"github.com/evpeople/douyin/pkg/constants"
	tracer2 "github.com/evpeople/douyin/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.VideoServiceName)
	dal.Init()
	userdb.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}), // server name
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()), // tracer
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
