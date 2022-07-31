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

	"github.com/evpeople/douyin/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title         string ` json:"title,omitempty"`
	Author        int64  ` json:"author,omitempty"`
	PlayUrl       string ` json:"play_url,omitempty"`
	CoverUrl      string ` json:"cover_url,omitempty"`
	FavoriteCount int64  ` json:"favorite_count,omitempty"`
	CommentCount  int64  ` json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func (u *Video) TableName() string {
	return constants.VideoTableName
}

//TODO: 添加外键约束和关联查询。
func MGetPublishVideo(ctx context.Context, userIDs int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author= ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//TODO:考虑在查询中加上 time参数
func MGetVideo(ctx context.Context, time int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Limit(3).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateVideo(ctx context.Context, videos []*Video) error {
	return DB.WithContext(ctx).Create(videos).Error
}

// // QueryUser query list of user info
// func QueryUser(ctx context.Context, username string) (*User, error) {
// 	// res := make([]*User, 0)
// 	res := new(User)
// 	// ans:=DB.First(res, "user_name = ?", username)
// 	if err := DB.First(res, "user_name = ?", username).Error; err != nil {
// 		//没有找到数据，可能返回的是 RecordNotExist
// 		return res, err
// 	}
// 	return res, nil
// }
