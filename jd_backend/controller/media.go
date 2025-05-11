package controller

import (
	"fmt"
	"jd/database"
	mediaModel "jd/model/media"
	"jd/service"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageResponse struct {
	URL    string `json:"url"`
	Alt    string `json:"alt"`
	Poster string `json:"poster"`
}

const fileFolder = "uploads"

type MediaController struct {
	basePath     string
	mediaService *service.MediaService
}

func NewMediaController(mediaService *service.MediaService) *MediaController {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "http://localhost:8080/media" // 默认值
	}
	return &MediaController{
		mediaService: mediaService,
		basePath:     basePath,
	}
}

func (mc *MediaController) UploadImage(c *gin.Context) {
	authorId := uint(0)
	if os.Getenv("GIN_MODE") == "release" {
		userId, exist := c.Get("userId")
		if !exist {
			c.JSON(401, gin.H{
				"errno": 1,
				"msg":   "User not authenticated",
			})
			return
		}
	
		authorId = userId.(uint)
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{
			"errno": 1,
			"msg":   "Invalid form data",
		})
		return
	}

	files := form.File

	storePath := filepath.Join(".", fileFolder)
	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		if err := os.MkdirAll(storePath, os.ModePerm); err != nil {
			c.JSON(500, gin.H{
				"errno": 1,
				"msg":   "Failed to create directory",
			})
			return
		}
		return
	}
	imgLinks := []ImageResponse{}
	db := database.GetMysqlDb()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(500, gin.H{
				"errno": 1,
				"msg":   "Internal server error",
			})
		}
	}()
	for _, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			rawFileName := fileHeader.Filename
			fileName := generateRandomFileName(rawFileName)
			fullFilePath := filepath.Join(storePath, fileName)
			if err := c.SaveUploadedFile(fileHeader, fullFilePath); err != nil {
				fmt.Printf("Failed to save file: %v", err)
				continue
			}

			url := fmt.Sprintf("%s/%s", mc.basePath, fileName)
			imgLink := ImageResponse{
				URL:  url,
				Alt:  rawFileName,
			}
			media := mediaModel.Media{
				AuthorID:  authorId,
				MediaType: 1,
				MediaURL:  url,
			}
			if err := db.Create(&media).Error; err != nil {
				fmt.Printf("Failed to save media record: %v", err)
				c.JSON(500, gin.H{
					"errno": 1,
					"msg":   "Failed to save media record",
				})
				tx.Rollback()
				return
			}
			imgLinks = append(imgLinks, imgLink)
		}
	}
	tx.Commit()

	c.JSON(200, gin.H{
		"errno": 0,
		"data":  imgLinks,
	})
}

func generateRandomFileName(originalName string) string {
	ext := path.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)
	return fmt.Sprintf("%s_%s%s", nameWithoutExt, uuid.New().String()[0:8], ext)
}
