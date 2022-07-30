package cos

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// var DefaultClient *cos.Client

var DefaultClient = &cos.Client{}

func init() {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(os.Getenv("BACKETURL"))
	b := &cos.BaseURL{BucketURL: u}
	DefaultClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("SECRETID"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})

}
func UploadVideo(key string, r io.Reader) (string, error) {
	opt := &cos.ObjectPutOptions{}
	_, err := DefaultClient.Object.Put(context.Background(), key, r, opt)
	resp := DefaultClient.Object.GetObjectURL(key)
	return resp.String(), err
}
