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

	api := router.Group("/api")
	{
		api.POST("/user", handlers.CreatePlayer)
		api.GET("/user", handlers.GetPlayers)
		api.POST("/game", handlers.CreateGame)
		api.GET("/game", handlers.GetGames)
	}
	webSocket := router.Group("/ws")
	{
		webSocket.GET("/dart/:gameId", handlers.HandleDartWebSocket)
		webSocket.GET("/dart", handlers.HandleGeneralWebsocket)
	}

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("Error while starting server")
	}
}
