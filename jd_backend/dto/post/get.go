package post

import (
	"time"
)

type GetPostRequest struct {
	Limit  int `json:"limit"`
	LastId int `json:"id"`
}

type SimplePost struct {
	ID           uint   `json:"id"`
	Author       User   `json:"author"`
	Title        string `json:"title"`
	CoverImage   string `json:"cover"`
	ViewCount    int    `json:"views"`
	LikeCount    int    `json:"likes"`
	CommentCount int    `json:"comments"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GetPostResponse []SimplePost

func (gpr *GetPostRequest) Examine() error {
	if gpr.Limit <= 0 {
		gpr.Limit = 10
	}
	return nil
}
