// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package pack

import (
	"github.com/evpeople/douyin/cmd/core/feed/dal/db"
	"github.com/evpeople/douyin/kitex_gen/feed"
)

// User pack user info
func Video(u *db.Video) *feed.Video {
	if u == nil {
		return nil
	}

	return &feed.Video{
		Id: 0,
		Author: &feed.User{
			Id:            0,
			Name:          "",
			FollowCount:   new(int64),
			FollowerCount: new(int64),
			IsFollow:      false,
		},
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "",
	}
}

// Users pack list of user info
func Videos(us []*db.Video) []*feed.Video {
	videos := make([]*feed.Video, 0)
	for _, u := range us {
		if user2 := Video(u); user2 != nil {
			videos = append(videos, user2)
		}
	}
	return videos
}
