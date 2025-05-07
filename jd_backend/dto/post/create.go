package post

import (
	"fmt"
)

type CreatePostRequest struct {
	AuthorID uint
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func (cpr *CreatePostRequest) Examine() error {
	if cpr.Title == "" {
		return fmt.Errorf("title is required")
	}
	if cpr.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}