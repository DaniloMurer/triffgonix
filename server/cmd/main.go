package main

import (
	websocket "github.com/DaniloMurer/triffgonix/server/internal/api"
	"github.com/DaniloMurer/triffgonix/server/internal/api/game"
	"github.com/DaniloMurer/triffgonix/server/internal/api/player"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//	@title			Triffgonix API
//	@version		1.0
//	@description	Triffgonix api documentation

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		http://localhost:8080

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
		group.POST("/user", apiplayer.CreatePlayer)
		group.GET("/user", apiplayer.GetPlayers)
		group.POST("/game", apigame.CreateGame)
		group.GET("/game", apigame.GetGames)
	}
	socketGroup := router.Group("/ws")
	{
		socketGroup.GET("/dart/:gameId", websocket.HandleDartWebSocket)
		socketGroup.GET("/dart", websocket.HandleGeneralWebsocket)
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
