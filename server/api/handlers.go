package handlers

import (
	"net/http"
	"server/api/dto"
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
	cleanupHubs()
	// FIXME: only temporary
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while upgrading request to websocket protocol: %v", err)
		return
	}
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
	var player dto.Player
	err := c.BindJSON(&player)
	if err != nil {
		logger.Error("error while parsing player json")
		c.Status(http.StatusInternalServerError)
		return
	}
	err, newPlayer := database.CreatePlayer(player.ToEntity())
	if err != nil {
		logger.Error("error while saving player to database: %+v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, &newPlayer)
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

func broadcastNewGame(newGame *database.Game) {
	game := dto.Game{}
	game.FromEntity(newGame)
	for _, hub := range hubs {
		hub.BroadcastToClients(game)
	}
}

// cleanupHubs removes hubs with zero clients connected to it
func cleanupHubs() {
	for gameId, hub := range hubs {
		if len(hub.Clients) == 0 {
			delete(hubs, gameId)
		}
	}
}
