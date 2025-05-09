package post

import (
	"time"
)

type GetPostsListRequest struct {
	Limit  int `json:"limit"`
	LastId int `json:"id"`
}

type SimplePost struct {
	ID           uint      `json:"id"`
	Author       User      `json:"author"`
	Title        string    `json:"title"`
	CoverImage   string    `json:"cover"`
	ViewCount    int       `json:"views"`
	LikeCount    int       `json:"likes"`
	CommentCount int       `json:"comments"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetPostsListResponse []SimplePost

func (gpr *GetPostsListRequest) Examine() error {
	if gpr.Limit <= 0 {
		gpr.Limit = 10
	}
	return nil
}
