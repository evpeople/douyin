package constants

const (
	EtcdAddress        = "127.0.0.1:2379"
	MySQLDefaultDSN    = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	UserTableName      = "user"
	VideoTableName     = "video"
	UserServiceName    = "demouser"
	VideoServiceName   = "demovideo"
	PublishServiceName = "demopublish"
	ApiServiceName     = "demoapi"
	SecretKey          = "secret key"
	IdentityKey        = "id"
	StatusCode         = "status_code"
	StatusMsg          = "status_msg"
	User               = "user"
	ID                 = "id"
	Name               = "name"
	FollowCount        = "follow_count"
	FollowerCount      = "follower_count"
	IsFollow           = "is_follow"
)
