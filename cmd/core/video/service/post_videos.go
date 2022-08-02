package service

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/video/dal/db"
	"github.com/evpeople/douyin/kitex_gen/publish"
)

type PostVideosService struct {
	ctx context.Context
}

//  NewPostVideosService new PostVideosService
func NewPostVideosService(ctx context.Context) *PostVideosService {
	return &PostVideosService{ctx: ctx}
}

// MPostVideos multiple post video info
func (s *PostVideosService) PostVideos(req *publish.UploadFileRequest) error {
	v := db.Video{
		Title:         req.Title,
		Author:        req.AuthorId,
		PlayUrl:       req.VideoUrl,
		CoverUrl:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	return db.CreateVideo(context.Background(), []*db.Video{&v})
}
