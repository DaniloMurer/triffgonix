package api

import (
	"github.com/DaniloMurer/triffgonix/server/internal/api/dto"
	"github.com/DaniloMurer/triffgonix/server/internal/api/socket"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine"
	"github.com/DaniloMurer/triffgonix/server/internal/dart/engine/x01"
	"github.com/DaniloMurer/triffgonix/server/internal/database"
	"github.com/DaniloMurer/triffgonix/server/internal/models"
	"github.com/DaniloMurer/triffgonix/server/pkg/logging"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader           = websocket.Upgrader{}
	hubs               = map[string]socket.Hub{}
	generalConnections []*websocket.Conn
)

var logger = logging.NewLogger()

// HandleDartWebSocket godoc
// @Summary Handle game-specific WebSocket connections
// @Description Handles WebSocket connections for specific dart games
// @Tags websocket
// @Accept json
// @Param gameId path string true "Game ID"
// @Success 101 "Switching Protocols to WebSocket"
// @Failure 400 "Bad Request"
// @Router /ws/dart/{gameId} [get]
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
	var message socket.IncomingMessage
	err = conn.ReadJSON(&message)
	if err != nil {
		logger.Error("error while reading from socket connection: %v", err)
		return
	}
	switch *message.Type {
	case socket.Handshake:
		hub, exists := hubs[gameId]
		if exists {
			hub.RegisterNewClient(conn)
		}
	}
	hub, exists := hubs[gameId]
	if exists {
		go hub.HandleConnection(conn)
	} else {
		err := conn.Close()
		if err != nil {
			logger.Error("error while closing websocket connection: %+v", err)
		}
	}
}

// HandleGeneralWebsocket godoc
// @Summary Handle general WebSocket connections
// @Description Handles WebSocket connections for general game updates
// @Tags websocket
// @Accept json
// @Success 101 "Switching Protocols to WebSocket"
// @Failure 400 "Bad Request"
// @Router /ws/dart [get]
func HandleGeneralWebsocket(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while upgrading request to websocket protocol: %v", err)
		return
	}
	var games []engine.GameDto
	for _, hub := range hubs {
		games = append(games, hub.GetGame())
	}
	generalConnections = append(generalConnections, conn)
	err = conn.WriteJSON(socket.OutgoingMessage{Type: socket.Games, Content: games})
	if err != nil {
		logger.Error("error while writing games message: %+v", err)
		return
	}
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Creates a new player in the system
// @Tags players
// @Accept json
// @Produce json
// @Param player body dto.Player true "Player information"
// @Success 201 {object} models.Player "Created player"
// @Failure 500 "Internal Server Error"
// @Router /api/user [post]
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

// GetPlayers godoc
// @Summary Get all players
// @Description Retrieves all players from the system
// @Tags players
// @Produce json
// @Success 200 {array} dto.Player "List of players"
// @Router /api/user [get]
func GetPlayers(c *gin.Context) {
	users := database.FindAllUsers()
	var userDtos []dto.Player
	for _, user := range users {
		userDto := dto.Player{}
		userDtos = append(userDtos, userDto.FromEntity(&user))
	}
	c.JSON(http.StatusOK, &userDtos)
}

// CreateGame godoc
// @Summary Create a new game
// @Description Creates a new dart game and sets up the corresponding hub
// @Tags games
// @Accept json
// @Produce json
// @Param game body dto.Game true "Game information"
// @Success 201 {object} models.Game "Created game"
// @Failure 500 "Internal Server Error"
// @Router /api/game [post]
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
	newHub := socket.Hub{Id: savedGame.Id, Clients: map[*socket.Client]bool{}, Game: game}
	hubs[strconv.FormatUint(uint64(savedGame.Id), 10)] = newHub
	broadcastNewGame(savedGame)
	c.JSON(http.StatusCreated, &savedGame)
}

// GetGames godoc
// @Summary Get all games
// @Description Retrieves all games from the system
// @Tags games
// @Produce json
// @Success 200 {array} models.Game "List of games"
// @Router /api/game [get]
func GetGames(c *gin.Context) {
	games := database.FindAllGames()
	c.JSON(http.StatusOK, &games)
}

func broadcastNewGame(newGame *models.Game) {
	game := dto.Game{}
	game.FromEntity(newGame)
	for _, conn := range generalConnections {
		var games []engine.GameDto
		for _, hub := range hubs {
			games = append(games, hub.GetGame())
		}
		generalConnections = append(generalConnections, conn)
		err := conn.WriteJSON(socket.OutgoingMessage{Type: socket.Games, Content: games})
		if err != nil {
			logger.Error("error while writing games message: %+v", err)
			return
		}
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
