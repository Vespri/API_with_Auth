package routers

import (
	"final_project/config"
	"final_project/controller"
	"final_project/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	db := config.StartDB()
	router := gin.Default()
	user := controller.NewUserController(db)
	photo := controller.NewPhotoController(db)
	social := controller.NewSocialController(db)
	comment := controller.NewCommentController(db)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", user.UserLogin)
		userGroup.POST("/register", user.CreateUser)
		userGroup.PUT("/:userId", middleware.Auth(), user.UpdateUser)
		userGroup.DELETE("/:userId", middleware.Auth(), user.DeleteUser)
	}

	photoGroup := router.Group("/photos")
	{
		photoGroup.GET("/", middleware.Auth(), photo.FindAllPhoto)
		photoGroup.POST("/", middleware.Auth(), photo.CreatePhoto)
		photoGroup.PUT("/:photoId", middleware.Auth(), photo.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middleware.Auth(), photo.DeletePhoto)
	}

	socialGroup := router.Group("/socials")
	{
		socialGroup.GET("/", middleware.Auth(), social.FindAllSocial)
		socialGroup.POST("/", middleware.Auth(), social.CreateSocial)
		socialGroup.PUT("/:socialMediaId", middleware.Auth(), social.UpdateSocial)
		socialGroup.DELETE("/:socialMediaId", middleware.Auth(), social.DeleteSocial)
	}

	commentGroup := router.Group("/comments")
	{
		commentGroup.GET("/", middleware.Auth(), comment.FindAllComment)
		commentGroup.POST("/", middleware.Auth(), comment.CreateComment)
		commentGroup.PUT("/:commentId", middleware.Auth(), comment.UpdateComment)
		commentGroup.DELETE("/:commentId", middleware.Auth(), comment.DeleteComment)
	}

	return router
}
