package controller

import (
	"fmt"
	"os"

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

	req.AuthorID = 0
	if os.Getenv("GIN_MODE") == "release" {
		id, exist := c.Get("userId")
		if !exist {
			dto.ErrorResponse(c, dto.WithMessage("userId not found"))
			return
		}
		req.AuthorID = id.(uint)
	}

	err = pc.postService.CreatePost(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c)
}

func (pc *PostController) GetPostsList(c *gin.Context) {
	var req post.GetPostsListRequest
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	posts, err := pc.postService.GetPostsList(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	fmt.Println(posts)
	dto.SuccessResponse(c, dto.WithData(posts))
}

func (pc *PostController) GetMaxId(c *gin.Context) {
	maxId, err := pc.postService.GetMaxId()
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(maxId))
}

func (pc *PostController) GetPostById(c *gin.Context) {
	var req post.GetPostRequest
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	post, err := pc.postService.GetPost(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(post))
}