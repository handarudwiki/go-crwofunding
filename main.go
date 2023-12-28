package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/go-crowfunding/auth"
	"github.com/handarudwiki/go-crowfunding/connection"
	"github.com/handarudwiki/go-crowfunding/handler"
	"github.com/handarudwiki/go-crowfunding/middleware"
	"github.com/handarudwiki/go-crowfunding/user"
)

func main() {
	db := connection.GetConnection()

	userRepository := user.NewRepositry(db)

	userService := user.NewService(userRepository)
	authService := auth.NewJWTService()

	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	api := r.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailability)
	api.POST("/upload-avatar", middleware.Auth(authService, userService), userHandler.UploadAvatar)

	r.Run()
}
