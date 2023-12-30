package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/go-crowfunding/auth"
	"github.com/handarudwiki/go-crowfunding/connection"
	"github.com/handarudwiki/go-crowfunding/handler"
	"github.com/handarudwiki/go-crowfunding/middleware"
	"github.com/handarudwiki/go-crowfunding/module/campaign"
	"github.com/handarudwiki/go-crowfunding/module/user"
)

func main() {
	db := connection.GetConnection()

	userRepository := user.NewRepositry(db)
	campaignRepository := campaign.NewRepositry(db)

	userService := user.NewService(userRepository)
	campignService := campaign.NewService(campaignRepository)
	authService := auth.NewJWTService()

	userHandler := handler.NewUserHandler(userService)
	campaignHandler := handler.NewCampaignHandler(campignService)

	r := gin.Default()
	r.Static("/images", "./images")
	api := r.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailability)
	api.POST("/upload-avatar", middleware.Auth(authService, userService), userHandler.UploadAvatar)

	api.GET("campaigns", campaignHandler.GetCampaigns)
	api.GET("campaigns/:id", campaignHandler.GetCampaignByID)
	api.POST("campaigns", middleware.Auth(authService, userService), campaignHandler.Create)
	api.PUT("campaigns/:id", middleware.Auth(authService, userService), campaignHandler.Update)

	r.Run()
}
