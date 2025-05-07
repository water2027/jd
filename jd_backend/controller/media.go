package controller

import (
	"jd/service"
)

type MediaController struct {
	mediaService *service.MediaService
}

func NewMediaController() *MediaController {
	return &MediaController{
		mediaService: service.NewMediaService(),
	}
}
