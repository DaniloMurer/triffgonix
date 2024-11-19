package main

import (
	"server/internal/triffgonix/api"
	"server/internal/triffgonix/database"

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

	group := router.Group("/group")
	{
		group.POST("/user", api.CreatePlayer)
		group.GET("/user", api.GetPlayers)
		group.POST("/game", api.CreateGame)
		group.GET("/game", api.GetGames)
	}
	socketGroup := router.Group("/ws")
	{
		socketGroup.GET("/dart/:gameId", api.HandleDartWebSocket)
		socketGroup.GET("/dart", api.HandleGeneralWebsocket)
	}

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("Error while starting server")
	}
}
