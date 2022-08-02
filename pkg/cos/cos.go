package cos

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	VideoClient *cos.Client
	PicClient   *cos.Client
)

func init() {
	VideoClient = createClient("BACKETVIDEOURL")
	PicClient = createClient("BACKETPICURL")
}

//UploadVideo 用于上传实拍，同时会生成cover_url
func UploadVideo(key string, r io.Reader) (videoUrl string, picUrl string, err error) {
	optPut := &cos.ObjectPutOptions{}
	_, err = VideoClient.Object.Put(context.Background(), key, r, optPut)
	if err != nil {
		log.Panicln(err)
	}
	respVideo := VideoClient.Object.GetObjectURL(key)

	optSnap := &cos.GetSnapshotOptions{
		Time: 1,
	}
	PicResp, err := VideoClient.CI.GetSnapshot(context.Background(), key, optSnap)
	if err != nil {
		log.Panicln(err)
	}
	_, err = PicClient.Object.Put(context.Background(), key, PicResp.Body, optPut)
	if err != nil {
		log.Panicln(err)
	}
	respPic := PicClient.Object.GetObjectURL(key)
	return respVideo.String(), respPic.String(), err
}
func createClient(backetUrl string) *cos.Client {

	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, err := url.Parse(os.Getenv(backetUrl))
	if err != nil {
		panic(err)
	}
	return cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("SECRETID"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})
}
