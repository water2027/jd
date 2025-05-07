package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"os"

	"jd/controller"
	"jd/service"

	"jd/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	origin := os.Getenv("FRONTEND_URL")
	if origin == "" {
		origin = "http://localhost:5173" // 默认值
	}

	// 设置跨域请求
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{origin},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	r.Use(middleware.AuthMiddleware())

	user := controller.NewUserController(service.NewUserService())
	media := controller.NewMediaController(service.NewMediaService())
	post := controller.NewPostController(service.NewPostService())
	// 设置路由组
	api := r.Group("/api")
	api.POST("/info", user.GetUserInfo)

	mediaRoute := api.Group("/media")
	mediaRoute.POST("/upload-image", media.UploadImage)

	postRoute := api.Group("/post")
	postRoute.POST("/create", post.CreatePost)

	publicRoute := api.Group("/public")

	publicRoute.POST("/register", user.Register)
	publicRoute.POST("/login", user.Login)
	publicRoute.POST("/reset", user.ResetPassword)
	publicRoute.POST("/send-code", user.SendVCode)

	publicRoute.POST("/get-post", post.GetPost)

	return r
}