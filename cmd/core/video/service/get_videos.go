package service

import (
	"context"
	"log"

	userdb "github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/cmd/core/video/dal/db"
	"github.com/evpeople/douyin/kitex_gen/publish"
)

type GetVideosService struct {
	ctx context.Context
}

// NewMGetVideoService new MGetVideoService
func NewMGetVideosService(ctx context.Context) *GetVideosService {
	return &GetVideosService{ctx: ctx}
}

// MGetVideos multiple get list of videos info
func (s *GetVideosService) MGetVideos(req *publish.DouyinPublishRequest) ([]*publish.Video, error) {
	videos, err := db.MGetVideo(context.TODO(), *req.UserId)
	if err != nil {
		log.Println(err)
	}
	log.Println("MGetVideos", videos)
	ans := make([]*publish.Video, len(videos))
	for i := 0; i < len(videos); i++ {
		user, err := userdb.MGetUser(context.TODO(), videos[i].Author)
		if err != nil {
			log.Println(err)
		}
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
}
