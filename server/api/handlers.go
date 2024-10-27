package handlers

import (
	"net/http"
	"server/api/dto"
	"server/core/domain"
	"server/dart/engine"
	"server/dart/engine/x01"
	"server/database"
	"server/logging"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	hubs     = map[string]Hub{}
)

var logger logging.Logger = logging.NewLogger()

func HandleDartWebSocket(c *gin.Context) {
	// FIXME: only temporary
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while upgrading request to websocket protocol: %v", err)
		return
	}
	mockCreateGame()
	gameId := c.Param("gameId")
	// get message from socket
	var message dto.Message
	err = conn.ReadJSON(&message)
	if err != nil {
		logger.Error("error while reading from socket connection: %v", err)
		return
	}
	// if a new handshake is made, register client in the correct hub. if no hub exists, create one
	switch *message.Type {
	case dto.Handshake:
		hub, exists := hubs[gameId]
		if exists {
			hub.RegisterNewClient(conn)
		}
	}
	hub, exists := hubs[gameId]
	if exists {
		go hub.HandleConnection(conn)
	} else {
		conn.Close()
	}
}

func CreatePlayer(c *gin.Context) {
	player := domain.Player{Id: 0, PlayerName: "test2"}
	database.CreatePlayer(player.ToPlayerEntity())
	c.JSON(http.StatusAccepted, gin.H{"text": "hello"})
}

func GetPlayers(c *gin.Context) {
	users := database.FindAllUsers()
	c.JSON(http.StatusFound, &users)
}

func CreateGame(c *gin.Context) {
	var newGame dto.Game
	err := c.BindJSON(&newGame)
	if err != nil {
		logger.Error("error while binding json: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	logger.Trace("new game to be saved: %+v", &newGame)
	err, savedGame := database.CreateGame(newGame.ToEntity())
	if err != nil {
		logger.Error("error while saving game to the databse: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// create new game and hub
	players := engine.Players{}
	for _, player := range newGame.Players {
		players.Add(&engine.Player{Value: player.ToDomain()})
	}
	game := engine.Game{
		Name:    newGame.Name,
		Players: &players,
		Engine:  x01.New(newGame.StartingScore),
	}
	newHub := Hub{Id: savedGame.Id, Clients: map[*Client]bool{}, Game: game}
	hubs[strconv.FormatUint(uint64(savedGame.Id), 10)] = newHub

	c.JSON(http.StatusCreated, &savedGame)
}

func GetGames(c *gin.Context) {
	games := database.FindAllGames()
	c.JSON(http.StatusFound, &games)
}

func mockCreateGame() engine.Game {
	player := engine.Player{Value: &domain.Player{PlayerName: "1", Id: 1, Score: 301}, Turns: []engine.Turn{}}
	player2 := engine.Player{Value: &domain.Player{PlayerName: "2", Id: 2, Score: 301}, Turns: []engine.Turn{}}
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
