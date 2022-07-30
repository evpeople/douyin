package dal

import "github.com/evpeople/douyin/cmd/core/video/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
