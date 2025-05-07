package service

import (
	"fmt"
	"jd/database"
	"jd/dto/post"
	postModel "jd/model/post"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (ps *PostService) CreatePost(req post.CreatePostRequest) error {
	post := &postModel.Post{
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Content:  req.Content,
	}

	db := database.GetMysqlDb()
	if err := db.Create(post).Error; err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	return nil
}
