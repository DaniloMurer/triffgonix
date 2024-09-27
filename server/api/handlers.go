package handlers

import (
	"log"
	"net/http"
	"server/api/dto"
	"server/core/domain"
	"server/dart/engine"
	"server/dart/engine/x01"
	"server/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	hubs     = map[string]Hub{}
)

func HandleDartWebSocket(c *gin.Context) {
	// TODO: make all this concurrent using goroutines for blazingly fast performance
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic("error while upgrading to websocket protocol")
	}
	defer conn.Close()
	mockCreateGame()
	gameId := c.Param("gameId")
	for {
		var message dto.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error: %+v", err)
		}
		switch *message.Type {
		case dto.Handshake:
			hub, exists := hubs[gameId]
			if exists {
				hub.RegsiterNewClient(conn)
			} else {
				hub = Hub{Id: gameId, Clients: map[*Client]bool{}, Game: mockCreateGame()}
				hub.RegsiterNewClient(conn)
				hubs[gameId] = hub
			}
		case dto.Throw:
			hub := hubs[gameId]
			hub.BroadcastMessage(hub.Game)
		}
	}
}

func CreatePlayer(c *gin.Context) {
	player := domain.Player{Id: 0, PlayerName: "test"}
	database.CreatePlayer(player.ToPlayerEntity())
	c.JSON(http.StatusAccepted, gin.H{"text": "hello"})
}

func GetPlayers(c *gin.Context) {
	users := database.FindAllUsers()
	c.JSON(http.StatusFound, &users)
}

func CreateGame(c *gin.Context) {
	// TODO: implement game creation through post request
	/*game := engine.Game{
		Name:    "test",
		Players: &engine.Players{},
		Engine:  x01.New(301),
	}
	games["201"] = game*/
}

func mockCreateGame() engine.Game {
	// create new game
	game := engine.Game{
		Name:    "test",
		Players: &engine.Players{},
		Engine:  x01.New(301),
	}
	return game
}
