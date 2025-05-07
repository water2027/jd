package service

import (
	"fmt"
	"jd/database"
	postDto "jd/dto/post"
	postModel "jd/model/post"
)

type PostService struct {
}

func NewPostService() *PostService {
	return &PostService{}
}

func (ps *PostService) CreatePost(req postDto.CreatePostRequest) error {
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

func (ps *PostService) GetPosts(req postDto.GetPostRequest) (postDto.GetPostResponse, error) {
	var resp postDto.GetPostResponse

	type QueryResult struct {
		postModel.Post
		UserID   uint   `gorm:"column:user_id"`
		UserName string `gorm:"column:user_name"`
	}

	var results []QueryResult

	db := database.GetMysqlDb()

	query := db.Table("post").
		Select("post.*, user.id AS user_id, user.name AS user_name").
		Joins("LEFT JOIN user ON post.author_id = user.id")

	if req.LastId > 0 {
		query = query.Where("post.id <= ?", req.LastId)
	}

	if err := query.Limit(req.Limit).Order("post.id desc").Scan(&results).Error; err != nil {
		return nil, err
	}

	for _, result := range results {
		post := postDto.SimplePost{
			ID: result.ID,
			Author: postDto.User{
				Id:   result.UserID,
				Name: result.UserName,
			},
			Title:        result.Post.Title,
			CoverImage:   result.Post.CoverImage,
			LikeCount:    result.Post.LikeCount,
			CommentCount: result.Post.CommentCount,
			ViewCount:    result.Post.ViewCount,
			CreatedAt:    result.Post.CreatedAt,
			UpdatedAt:    result.Post.UpdatedAt,
		}

		resp = append(resp, post)
	}

	return resp, nil
}

func (ps *PostService) GetMaxId() (postDto.GetMaxIdResponse, error) {
	var maxId int64
	db := database.GetMysqlDb()
	if err := db.Table("post").Select("MAX(id)").Scan(&maxId).Error; err != nil {
		return 0, err
	}
	return postDto.GetMaxIdResponse(maxId), nil
}