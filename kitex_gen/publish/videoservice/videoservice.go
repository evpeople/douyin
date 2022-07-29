// Code generated by Kitex v0.3.4. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/evpeople/douyin/kitex_gen/publish"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*publish.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPublishVideos": kitex.NewMethodInfo(getPublishVideosHandler, newGetPublishVideosArgs, newGetPublishVideosResult, false),
		"PostVideos":       kitex.NewMethodInfo(postVideosHandler, newPostVideosArgs, newPostVideosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "publish",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.4",
		Extra:           extra,
	}
	return svcInfo
}

func getPublishVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.DouyinPublishRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.VideoService).GetPublishVideos(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPublishVideosArgs:
		success, err := handler.(publish.VideoService).GetPublishVideos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPublishVideosResult)
		realResult.Success = success
	}
	return nil
}
func newGetPublishVideosArgs() interface{} {
	return &GetPublishVideosArgs{}
}

func newGetPublishVideosResult() interface{} {
	return &GetPublishVideosResult{}
}

type GetPublishVideosArgs struct {
	Req *publish.DouyinPublishRequest
}

func (p *GetPublishVideosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPublishVideosArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPublishVideosArgs) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPublishVideosArgs_Req_DEFAULT *publish.DouyinPublishRequest

func (p *GetPublishVideosArgs) GetReq() *publish.DouyinPublishRequest {
	if !p.IsSetReq() {
		return GetPublishVideosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPublishVideosArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPublishVideosResult struct {
	Success *publish.DouyinPublishResponse
}

var GetPublishVideosResult_Success_DEFAULT *publish.DouyinPublishResponse

func (p *GetPublishVideosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPublishVideosResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPublishVideosResult) Unmarshal(in []byte) error {
	msg := new(publish.DouyinPublishResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPublishVideosResult) GetSuccess() *publish.DouyinPublishResponse {
	if !p.IsSetSuccess() {
		return GetPublishVideosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPublishVideosResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.DouyinPublishResponse)
}

func (p *GetPublishVideosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func postVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(publish.UploadFileRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(publish.VideoService).PostVideos(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PostVideosArgs:
		success, err := handler.(publish.VideoService).PostVideos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PostVideosResult)
		realResult.Success = success
	}
	return nil
}
func newPostVideosArgs() interface{} {
	return &PostVideosArgs{}
}

func newPostVideosResult() interface{} {
	return &PostVideosResult{}
}

type PostVideosArgs struct {
	Req *publish.UploadFileRequest
}

func (p *PostVideosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PostVideosArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PostVideosArgs) Unmarshal(in []byte) error {
	msg := new(publish.UploadFileRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PostVideosArgs_Req_DEFAULT *publish.UploadFileRequest

func (p *PostVideosArgs) GetReq() *publish.UploadFileRequest {
	if !p.IsSetReq() {
		return PostVideosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PostVideosArgs) IsSetReq() bool {
	return p.Req != nil
}

type PostVideosResult struct {
	Success *publish.BaseResponse
}

var PostVideosResult_Success_DEFAULT *publish.BaseResponse

func (p *PostVideosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PostVideosResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PostVideosResult) Unmarshal(in []byte) error {
	msg := new(publish.BaseResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PostVideosResult) GetSuccess() *publish.BaseResponse {
	if !p.IsSetSuccess() {
		return PostVideosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PostVideosResult) SetSuccess(x interface{}) {
	p.Success = x.(*publish.BaseResponse)
}

func (p *PostVideosResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetPublishVideos(ctx context.Context, Req *publish.DouyinPublishRequest) (r *publish.DouyinPublishResponse, err error) {
	var _args GetPublishVideosArgs
	_args.Req = Req
	var _result GetPublishVideosResult
	if err = p.c.Call(ctx, "GetPublishVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PostVideos(ctx context.Context, Req *publish.UploadFileRequest) (r *publish.BaseResponse, err error) {
	var _args PostVideosArgs
	_args.Req = Req
	var _result PostVideosResult
	if err = p.c.Call(ctx, "PostVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}