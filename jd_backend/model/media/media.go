package media

import (
	// "gorm.io/gorm"
	"time"
)

type MediaType int

const (
	MediaTypeImage MediaType = iota + 1
	MediaTypeAudio
	MediaTypeVideo
)

type Media struct {
	ID        uint      `gorm:"primaryKey"`
	AuthorID  uint      `gorm:"not null"`
	MediaType MediaType `gorm:"type:tinyint;not null"`
	MediaURL  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
