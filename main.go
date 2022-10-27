package main

import (
	"final_project/auth"
	"final_project/config"
	"final_project/handler"
	"final_project/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.StartDB()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.PUT("/:id", auth.Authentication(userService), userHandler.UpdateUser)
	userRouter.DELETE("/:id", auth.Authentication(userService), userHandler.DeleteUser)

	router.Run()
}
