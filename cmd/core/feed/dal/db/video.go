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

package db

import (
	"context"

	"github.com/evpeople/douyin/cmd/core/user/dal/db"
	"github.com/evpeople/douyin/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Author        db.User `json:"author"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int     `json:"favorite_count"`
	CommentCount  int     `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
	Title         string  `json:"title"`
}

func (u *Video) TableName() string {
	return constants.VideoTableName
}

// MGetVideos multiple get list of user info
func MGetVideo(ctx context.Context, userIDs int64) (*Video, error) {
	// res := make([]*Video, 0)
	res := new(Video)
	if err := DB.WithContext(ctx).Where("id = ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
