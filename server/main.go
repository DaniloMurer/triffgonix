package main

import (
	handlers "server/api"
	"server/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.AutoMigrate()

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Authorization", "Origin"}

	router.Use(cors.New(corsConfig))

	api := router.Group("/api")
	{
		api.POST("/user", handlers.CreatePlayer)
		api.GET("/user", handlers.GetUsers)
	}

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("Error while starting server")
	}
}
