package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"jd/dto"
	"jd/dto/post"
	"jd/service"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req post.CreatePostRequest
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	id, exist := c.Get("userId")
	if !exist {
		dto.ErrorResponse(c, dto.WithMessage("userId not found"))
		return
	}

	req.AuthorID = id.(uint)

	err = pc.postService.CreatePost(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c)
}

func (pc *PostController) GetPost(c *gin.Context) {
	var req post.GetPostRequest
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	posts, err := pc.postService.GetPosts(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	fmt.Println(posts)
	dto.SuccessResponse(c, dto.WithData(posts))
}
