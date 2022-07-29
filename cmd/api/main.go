package main

import (
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/evpeople/douyin/cmd/api/handlers"
	"github.com/evpeople/douyin/cmd/api/rpc"
	"github.com/evpeople/douyin/pkg/constants"
	"github.com/evpeople/douyin/pkg/tracer"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()

	v1 := r.Group("/douyin")
	user1 := v1.Group("/user")
	user1.GET("", handlers.GetUser)
	// authMiddl
	user1.POST("/login", handlers.AuthMiddleware.LoginHandler)
	user1.POST("/register", handlers.Register)
	feed := v1.Group("/feed")
	feed.Use(handlers.AuthMiddleware.MiddlewareFunc())
	feed.GET("", handlers.GetVideos)
	publish := v1.Group("/publish")
	//TODO:添加中间件把Token放到header中
	publish.Use(handlers.AuthMiddleware.MiddlewareFunc())
	// publish.GET("/list",handler.)
	publish.POST("/action", handlers.UploadVideo)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
