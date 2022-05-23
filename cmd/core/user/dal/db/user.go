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

type User struct {
	gorm.Model
	UserName       string `json:"user_name" gorm:"unique"`
	Password       string `json:"password"`
	UserId         int64  `json:"User_id"`
	Follow_count   int64  `json:"follow_count`
	Follower_count int64  `json:"follower_count`
	IsFollow       bool   `json:"is_follow`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUser(ctx context.Context, userIDs int64) (*User, error) {
	// res := make([]*User, 0)
	res := new(User)
	if err := DB.WithContext(ctx).Where("id = ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, username string) (*User, error) {
	// res := make([]*User, 0)
	res := new(User)
	// ans:=DB.First(res, "user_name = ?", username)
	if err := DB.First(res, "user_name = ?", username).Error; err != nil {
		//没有找到数据，可能返回的是 RecordNotExist
		return res, err
	}
	return res, nil
}
