package service

import (
	"context"
	"log"

	userdb "github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/cmd/core/video/dal/db"
	"github.com/evpeople/douyin/kitex_gen/publish"
)

type GetPublishVideosService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetPublishVideosService(ctx context.Context) *GetPublishVideosService {
	return &GetPublishVideosService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *GetPublishVideosService) MGetVideos(req *publish.DouyinPublishRequest) ([]*publish.Video, error) {
	// modelUser, err := db.MGetUser(s.ctx, req.BaseResp.UserId)
	// if err != nil {
	videos, err := db.MGetPublishVideo(context.TODO(), *req.UserId)
	if err != nil {
		log.Println(err)
	}
	user, err := userdb.MGetUser(context.TODO(), *req.UserId)
	if err != nil {
		log.Println(err)
	}
	ans := make([]*publish.Video, len(videos))
	for i := 0; i < len(videos); i++ {
		ans[i] = &publish.Video{
			Id: int64(videos[i].ID),
			Author: &publish.User{
				Id:            user.UserId,
				Name:          user.UserName,
				FollowCount:   &user.Follow_count,
				FollowerCount: &user.Follower_count,
				IsFollow:      user.IsFollow,
			},
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavorite,
			Title:         videos[i].Title,
		}
	}
	return ans, nil
	// }
	// return pack.User(modelUser), nil
}
