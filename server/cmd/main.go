package main

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// sets up the router
func setupRouter() *gin.Engine {
	database.AutoMigrate()

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Authorization", "Origin"}

	router.Use(cors.New(corsConfig))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome")
	})

	group := router.Group("/api")
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

	return router
}

func main() {
	router := setupRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("Error while starting server")
	}
}
