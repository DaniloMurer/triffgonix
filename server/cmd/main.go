package main

import (
	websocket "github.com/DaniloMurer/triffgonix/server/internal/api"
	"github.com/DaniloMurer/triffgonix/server/internal/api/game"
	"github.com/DaniloMurer/triffgonix/server/internal/api/player"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	fiberwebsocket "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
)

//	@title			Triffgonix API
//	@version		1.0
//	@description	Triffgonix api documentation

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// sets up the router
func setupFiberApp() *fiber.App {
	database.AutoMigrate()
	app := fiber.New()

	api := app.Group("/api")

	api.Post("/player", apiplayer.CreatePlayer)
	api.Get("/player", apiplayer.GetPlayers)
	api.Post("/game", apigame.CreateGame)
	api.Get("/game", apigame.GetGames)

	ws := app.Group("/ws")
	ws.Get("/dart", func(c *fiber.Ctx) error {
		if fiberwebsocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	ws.Get("/dart/general", fiberwebsocket.New(websocket.HandleGeneralWebsocket))
	ws.Get("/dart/:gameId", fiberwebsocket.New(websocket.HandleDartWebSocket))
	return app
}

func main() {
	app := setupFiberApp()
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
