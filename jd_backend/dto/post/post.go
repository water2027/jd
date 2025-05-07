package post

import (
	"time"
)

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID           uint   `json:"id"`
	Author       User   `json:"author"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CoverImage   string `json:"cover"`
	ViewCount    int    `json:"views"`
	LikeCount    int    `json:"likes"`
	CommentCount int    `json:"comments"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
