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
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic("error while upgrading to websocket protocol")
	}
	mockCreateGame()
	gameId := c.Param("gameId")
	// get message from socket
	var message dto.Message
	err = conn.ReadJSON(&message)
	if err != nil {
		log.Printf("error: %+v", err)
	}
	// if a new handshake is made, register client in the correct hub. if no hub exists, create one
	switch *message.Type {
	case dto.Handshake:
		hub, exists := hubs[gameId]
		if exists {
			hub.RegisterNewClient(conn)
		} else {
			hub = Hub{Id: gameId, Clients: map[*Client]bool{}, Game: mockCreateGame()}
			hub.RegisterNewClient(conn)
			hubs[gameId] = hub
		}
	}
	hub := hubs[gameId]
	go hub.HandleConnection(conn)
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
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1}, Turns: []engine.Turn{}}
	// create new game
	game := engine.Game{
		Name:    "test",
		Players: &engine.Players{},
		Engine:  x01.New(301),
	}
	game.Players.Add(&player)
	game.Players.Add(&player2)
	return game
}
