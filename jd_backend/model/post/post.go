package post

import (
	// "gorm.io/gorm"
	"time"
)

type Post struct {
	ID           uint      `gorm:"primaryKey"`
	AuthorID     uint      `gorm:"not null"`
	Title        string    `gorm:"size:255;not null"`
	Content      string    `gorm:"type:text;not null"`
	CoverImage   string    `gorm:"size:255"`
	ViewCount    int       `gorm:"default:0"`
	LikeCount    int       `gorm:"default:0"`
	CommentCount int       `gorm:"default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
