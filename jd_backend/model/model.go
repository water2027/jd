package model

import (
	"gorm.io/gorm"

	"jd/model/user"
	"jd/model/post"
	"jd/model/media"
)

func InitTable(db *gorm.DB) error {
	// 自动迁移数据库表结构
	err := db.AutoMigrate(
		&user.User{},
		&post.Post{},
		&media.Media{},
	)
	if err != nil {
		return err
	}
	return nil
}